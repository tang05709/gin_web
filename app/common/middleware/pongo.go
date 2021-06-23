package middleware

import (
	"net/http"
	"strings"

	"github.com/flosch/pongo2"
	. "github.com/gin-gonic/gin"
)

func Pongo2() HandlerFunc {
	return func(c *Context) {
		c.Next()

		name := stringFromContext(c, "template")

		if name == "" {
			return
		}
		if !strings.HasSuffix(name, ".html") {
			name = name + ".html"
		}
		name = "resource/views/" + name
		template := pongo2.Must(pongo2.FromFile(name))
		err := template.ExecuteWriter(InjectParams(c), c.Writer)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func stringFromContext(c *Context, input string) string {
	raw, ok := c.Get(input)
	if ok {
		strVal, ok := raw.(string)
		if ok {
			return strVal
		}
	}
	return ""
}

func convertContext(thing interface{}, menus interface{}, router interface{}, user interface{}) pongo2.Context {
	if thing != nil {
		context, isMap := thing.(map[string]interface{})
		if isMap {
			if menus != nil {
				context["menus"] = menus
			}
			if router != nil {
				context["router"] = router
			}
			if user != nil {
				context["user"] = router
			}
			return context
		}
	}
	return nil
}

func getContext(templateData interface{}, err error) pongo2.Context {
	if templateData == nil || err != nil {
		return nil
	}
	contextData, isMap := templateData.(map[string]interface{})
	if isMap {
		return contextData
	}
	return nil
}
