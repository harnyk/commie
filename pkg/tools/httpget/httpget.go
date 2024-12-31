package httpget

import (
	"io/ioutil"
	"net/http"

	"github.com/harnyk/commie/pkg/agent"
)

type HTTPGetParams struct {
	URL     string            `mapstructure:"url"`
	Headers map[string]string `mapstructure:"headers,omitempty"`
}

var HTTPGetHandler agent.TypedHandler[HTTPGetParams, string] = func(params HTTPGetParams) (string, error) {
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

func HTTPGet() *agent.Tool {
	return agent.NewTool().
		WithName("http_get").
		WithDescription("Makes an HTTP GET request and returns the response body as a string").
		WithHandler(HTTPGetHandler.AcceptingMapOfAny()).
		WithSchema(
			agent.H{
				"type": "object",
				"properties": agent.H{
					"url": agent.H{"type": "string"},
					"headers": agent.H{
						"type":                 "object",
						"additionalProperties": agent.H{"type": "string"},
					},
				},
				"required": []string{"url"},
			},
		)
}
