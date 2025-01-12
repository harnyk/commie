package toolmw

import (
	"strings"
	"text/template"

	"github.com/harnyk/commie/pkg/ui"
	"github.com/harnyk/gena"
)

type ConsentMmiddleware struct {
	messageTemplateString string
	messageTemplate       *template.Template
}

func NewConsentMmiddleware(messageTemplate string) gena.ToolMiddleware {
	c := &ConsentMmiddleware{
		messageTemplateString: messageTemplate,
	}

	c.compileMessageTemplate()

	return c
}

func (c *ConsentMmiddleware) compileMessageTemplate() {
	c.messageTemplate = template.Must(template.New("consent").Parse(c.messageTemplateString))
}

func (c *ConsentMmiddleware) Execute(params gena.H, tool *gena.Tool) (gena.ToolMiddlewareResult, error) {
	builder := strings.Builder{}
	err := c.messageTemplate.Execute(&builder, params)
	if err != nil {
		return gena.ToolMiddlewareResult{}, err
	}
	answer, followup, err := ui.ShowConsent(builder.String())
	if err != nil {
		return gena.ToolMiddlewareResult{}, err
	}

	if answer == ui.ConsentResponseYes {
		return gena.ToolMiddlewareResult{
			Params: params,
		}, nil
	}
	if answer == ui.ConsentResponseNo {
		return gena.ToolMiddlewareResult{
			Result: "User declined this call",
			Stop:   true,
		}, nil
	}
	if answer == ui.ConsentResponseFollowUp {
		return gena.ToolMiddlewareResult{
			Result: "User declined this call with a comment: " + followup,
			Stop:   true,
		}, nil
	}
	return gena.ToolMiddlewareResult{}, nil
}
