package toolmw

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/harnyk/commie/pkg/ui"
	"github.com/harnyk/gena"
)

type ConsentMmiddleware struct {
	messageTemplateString string
	messageTemplate       *template.Template
}

func NewConsentMiddleware(messageTemplate string) gena.ToolMiddleware {
	shieldEmoji := "🛡️"

	c := &ConsentMmiddleware{
		messageTemplateString: shieldEmoji + " " + messageTemplate,
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
		fmt.Println(ui.RenderMarkdown("**You approved this call**"))
		return gena.ToolMiddlewareResult{
			Params: params,
		}, nil
	}
	if answer == ui.ConsentResponseNo {
		fmt.Println(ui.RenderMarkdown("**You declined this call**"))
		return gena.ToolMiddlewareResult{
			Result: "User declined this call",
			Stop:   true,
		}, nil
	}
	if answer == ui.ConsentResponseFollowUp {
		fmt.Println(ui.RenderMarkdown("**You declined this call with a comment**: " + followup))
		return gena.ToolMiddlewareResult{
			Result: "User declined this call with a comment: " + followup,
			Stop:   true,
		}, nil
	}
	return gena.ToolMiddlewareResult{}, nil
}
