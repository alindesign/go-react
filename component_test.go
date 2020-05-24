package react

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestComponent struct {
	Name string
}

func (t *TestComponent) Render(ctx *Context) interface{} {
	return H1(nil, fmt.Sprintf("Hello %s! %s", t.Name, ctx.GetString("welcomeMessage")))
}

func TestIsComponent(t *testing.T) {
	t.Run("it should check if struct it's a Component class like", func(t *testing.T) {
		comp := &TestComponent{}

		assert.True(t, IsComponent(comp))
	})
}

func TestRenderClassLike(t *testing.T) {
	t.Run("it should correctly render component class like", func(t *testing.T) {
		content, err := Render(Div(nil,
			Span(nil, "span el"),
			&TestComponent{Name: "David"},
		), map[string]interface{}{
			"welcomeMessage": "How are you?",
		})

		assert.NoError(t, err)
		assert.Equal(t, "<div><span>span el</span><h1>Hello David! How are you?</h1></div>", content)
	})
}
