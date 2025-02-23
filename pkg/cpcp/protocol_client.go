package cpcp

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ProtocolClient struct {
	rw                   DuplexClient
	timeout              time.Duration
	mapMut               sync.Mutex
	done                 chan struct{}
	responseChannelsByID map[string]chan Response
}

func NewProtocolClient(rw DuplexClient) *ProtocolClient {
	return &ProtocolClient{
		rw:                   rw,
		timeout:              time.Second * 1,
		done:                 make(chan struct{}),
		mapMut:               sync.Mutex{},
		responseChannelsByID: map[string]chan Response{},
	}
}

func (h *ProtocolClient) Start() error {
	err := h.rw.Start()
	if err != nil {
		return err
	}
	go h.receive()
	return nil
}

func (h *ProtocolClient) SetTimeout(timeout time.Duration) *ProtocolClient {
	h.timeout = timeout
	return h
}

func (h *ProtocolClient) Stop() error {
	h.done <- struct{}{}
	return h.rw.Stop()
}

func (h *ProtocolClient) Send(requestPayload any, responsePayload any) error {
	id := h.nextID()
	req, err := h.serializeRequest(requestPayload, id)
	if err != nil {
		return err
	}
	ch := make(chan Response, 1)
	h.putResponseChannel(id, ch)
	h.rw.Send(req)
	select {
	case res := <-ch:
		{
			switch res.Type {
			case ResError:
				return NewErrorResponse(res.Payload)
			case ResOK:
				return json.Unmarshal([]byte(res.Payload), responsePayload)
			default:
				return fmt.Errorf("unknown response type: %s", res.Type)
			}
		}
	case <-time.After(h.timeout):
		return errors.New("timeout")
	}
}

func (h *ProtocolClient) receive() {
	for {
		select {
		case <-h.done:
			return
		case line := <-h.rw.Receive():
			var res Response
			err := json.Unmarshal([]byte(line), &res)
			if err != nil {
				continue
			}
			ch, ok := h.getResponseChannel(res.ID)
			if !ok {
				continue
			}
			ch <- res
			h.deleteResponseChannel(res.ID)
		}
	}
}

func (h *ProtocolClient) serializeRequest(requestPayload any, id string) (string, error) {
	payloadJson, err := json.Marshal(requestPayload)
	if err != nil {
		return "", err
	}
	req := Request{
		ID:      id,
		Payload: string(payloadJson),
	}
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	return string(reqBytes), nil
}

func (h *ProtocolClient) putResponseChannel(id string, ch chan Response) {
	h.mapMut.Lock()
	defer h.mapMut.Unlock()
	h.responseChannelsByID[id] = ch
}

func (h *ProtocolClient) getResponseChannel(id string) (chan Response, bool) {
	h.mapMut.Lock()
	defer h.mapMut.Unlock()
	ch, ok := h.responseChannelsByID[id]
	return ch, ok
}

func (h *ProtocolClient) deleteResponseChannel(id string) {
	h.mapMut.Lock()
	defer h.mapMut.Unlock()
	delete(h.responseChannelsByID, id)
}

func (h *ProtocolClient) nextID() string {
	return uuid.New().String()
}
