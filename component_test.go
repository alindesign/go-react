package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func helloComponent(Name string) *Element {
	return Component(func(ctx *Context) *Element {
		return H1(nil, Textf("Hello %s! %s", Name, ctx.GetString("welcomeMessage")))
	})
}

type LayoutMountPoints struct {
	Head    *Element
	Header  *Element
	Content *Element
}

func LayoutBase(mounts *LayoutMountPoints) *Element {
	return Component(func(ctx *Context) *Element {
		if mounts == nil {
			mounts = &LayoutMountPoints{}
		}

		return Document(&DocumentProps{
			Head: Fragment(
				MetaCharset(),
				mounts.Head,
			),
			Body: Fragment(
				CreateElement("header", nil, mounts.Header),
				mounts.Content,
			),
		}, nil)
	})
}

func TestRenderClassLike(t *testing.T) {
	t.Run("it should correctly render Component class like", func(t *testing.T) {
		content := Render(Div(nil,
			Span(nil, Text("span el")),
			helloComponent("David"),
		), map[string]interface{}{
			"welcomeMessage": "How are you?",
		})

		assert.Equal(t, "<div><span>span el</span><h1>Hello David! How are you?</h1></div>", content)
	})

	t.Run("it should correctly render advanced layout structure", func(t *testing.T) {
		comp := LayoutBase(&LayoutMountPoints{
			Content: Fragment(
				H1(nil, Text("Hello World!")),
			),
		})

		content := Render(comp)

		assert.Equal(t, `<!doctype html><html lang="en"><head><meta charset="UTF-8"/></head><body><header></header><h1>Hello World!</h1></body></html>`, content)
	})
}

func renderComponentStr(ctx *Context) *Element {
	return Text("test string")
}

func TestRenderComponent(t *testing.T) {
	r := Render(Component(renderComponentStr), nil)

	assert.Equal(t, "test string", r)
}

type compStruct struct {
	Name string
}

func (c compStruct) Render(ctx *Context) *Element {
	return H1(nil, Textf("Hello %s!", c.Name))
}

type compPStruct struct {
	Name string
}

func (c *compPStruct) Render(ctx *Context) *Element {
	return H1(nil, Textf("Hello %s!", c.Name))
}

func TestComponentStruct(t *testing.T) {
	t.Run("it should render component structs correctly", func(t *testing.T) {
		r := Render(Component(compStruct{Name: "Arthur"}), nil)
		rp := Render(Component(&compStruct{Name: "Arthur"}), nil)

		assert.Equal(t, "<h1>Hello Arthur!</h1>", r)
		assert.Equal(t, "<h1>Hello Arthur!</h1>", rp)
	})

	t.Run("it should render component pointer render structs correctly", func(t *testing.T) {
		rp := Render(Component(&compPStruct{Name: "Arthur"}), nil)

		assert.Equal(t, "<h1>Hello Arthur!</h1>", rp)
	})
}
