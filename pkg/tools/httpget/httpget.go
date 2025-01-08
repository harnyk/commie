package httpget

import (
	"io/ioutil"
	"net/http"

	"github.com/harnyk/gena"
)

type HTTPGetParams struct {
	URL     string            `mapstructure:"url"`
	Headers map[string]string `mapstructure:"headers,omitempty"`
}

type HttpGetHandler struct {
}

func NewHttpGetHandler() gena.ToolHandler {
	return &HttpGetHandler{}
}

func (h *HttpGetHandler) Execute(params gena.H) (any, error) {
	return gena.ExecuteTyped(h.execute, params)
}

func (h *HttpGetHandler) execute(params HTTPGetParams) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", params.URL, nil)
	if err != nil {
		return "", err
	}
	for key, value := range params.Headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func New() *gena.Tool {
	return gena.NewTool().
		WithName("http_get").
		WithDescription("Makes an HTTP GET request and returns the response body as a string").
		WithHandler(NewHttpGetHandler()).
		WithSchema(
			gena.H{
				"type": "object",
				"properties": gena.H{
					"url": gena.H{"type": "string"},
					"headers": gena.H{
						"type":                 "object",
						"additionalProperties": gena.H{"type": "string"},
					},
				},
				"required": []string{"url"},
			},
		)
}
