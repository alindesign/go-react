package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestComponent struct {
	Name string
}

func (t *TestComponent) Render(ctx *Context) interface{} {
	return H1(nil, Text("Hello %s! %s", t.Name, ctx.GetString("welcomeMessage")))
}

func TestIsComponent(t *testing.T) {
	t.Run("it should check if struct it's a Component class like", func(t *testing.T) {
		comp := &TestComponent{}

		assert.True(t, IsComponent(comp))
	})
}

type LayoutMountPoints struct {
	Head    *Element
	Header  *Element
	Content *Element
}

type LayoutBaseComponent struct {
	mounts *LayoutMountPoints
}

func (l *LayoutBaseComponent) Render(ctx *Context) interface{} {
	if l.mounts == nil {
		l.mounts = &LayoutMountPoints{}
	}

	return Document(&DocumentProps{
		Head: Fragment(
			MetaCharset(),
			l.mounts.Head,
		),
		Body: Fragment(
			CreateElement("header", nil, l.mounts.Header),
			l.mounts.Content,
		),
	}, nil)
}

func LayoutBase(mounts *LayoutMountPoints) *LayoutBaseComponent {
	return &LayoutBaseComponent{mounts: mounts}
}

func TestRenderClassLike(t *testing.T) {
	// t.Run("it should correctly render component class like", func(t *testing.T) {
	// 	content := Render(Div(nil,
	// 		Span(nil, "span el"),
	// 		&TestComponent{Name: "David"},
	// 	), map[string]interface{}{
	// 		"welcomeMessage": "How are you?",
	// 	})
	//
	// 	assert.Equal(t, "<div><span>span el</span><h1>Hello David! How are you?</h1></div>", content)
	// })

	// t.Run("it should correctly render advanced layout structure", func(t *testing.T) {
	// 	comp := LayoutBase(&LayoutMountPoints{
	// 		Content: Fragment(
	// 			H1(nil, "Hello World!"),
	// 		),
	// 	})
	//
	// 	content, err := Render(Fragment(
	// 		comp,
	// 	))
	//
	// 	assert.True(t, IsComponent(comp))
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, `<!doctype html><html lang="en"><head><meta charset="UTF-8" /></head><body><header></header><h1>Hello World!</h1></body></html>`, content)
	// })
}

type TestStringComp struct{}

func (t TestStringComp) Render(ctx *Context) interface{} {
	return "test string"
}

func TestRenderComponent(t *testing.T) {
	r := renderComponent(&TestStringComp{}, nil)

	assert.Equal(t, "test string", r)
}
