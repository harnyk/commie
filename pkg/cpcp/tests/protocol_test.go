package cpcp_test

import (
	"log/slog"
	"testing"

	"github.com/harnyk/commie/pkg/cpcp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProtocolSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Protocol Suite")
}

var _ = Describe("Protocol", func() {
	Context("when making request", func() {

		It("should return the correct sum", func() {
			proto := newProtocol()
			err := proto.Start()
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err := proto.Stop()
				Expect(err).ToNot(HaveOccurred())
			}()
			sum, err := sendAdd(proto, 10, 20)
			Expect(err).ToNot(HaveOccurred())
			Expect(sum).To(Equal(30))
		})

		It("should return an error", func() {
			proto := newProtocol()
			err := proto.Start()
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				err := proto.Stop()
				Expect(err).ToNot(HaveOccurred())
			}()
			err = sendMakeFail(proto)

			Expect(err).To(MatchError(ContainSubstring(`{"msg":"fail"}`)))
		})
	})
})

func sendAdd(proto *cpcp.ProtocolClient, a, b int) (int, error) {
	type requestBody = struct {
		Type string `json:"type"`
		A    int    `json:"a"`
		B    int    `json:"b"`
	}

	type responseBody = struct {
		C int `json:"c"`
	}

	req := &requestBody{
		Type: "add",
		A:    a,
		B:    b,
	}
	res := &responseBody{}

	if err := proto.Send(req, res); err != nil {
		return 0, err
	}

	return res.C, nil
}

func sendMakeFail(proto *cpcp.ProtocolClient) error {
	type MakeFailRequestPayload = struct {
		Type string `json:"type"`
	}

	req := &MakeFailRequestPayload{
		Type: "make_fail",
	}

	if err := proto.Send(req, nil); err != nil {
		return err
	}

	return nil
}

func newTransport() cpcp.DuplexClient {
	return cpcp.NewProcessClient(
		slog.Default(),
		"node",
		"./protocol_server.js",
	)
}

func newProtocol() *cpcp.ProtocolClient {
	return cpcp.NewProtocolClient(newTransport())
}
