package main

import (
	"fmt"
	"github.com/alindesign/go-react"
)

func List(num int) *react.Element {
	var items []*react.Element

	for i := 0; i < num; i++ {
		items = append(items, react.CreateElement("li", react.Props{
			"className": fmt.Sprintf("list__item item-%d", i+1),
		}, react.Text("Item %d", i+1)))
	}

	return react.CreateElement("ul", react.Props{
		"className": "list__root",
	}, items...)
}

func App() *react.Element {
	return react.CreateElement("div", react.Props{
		"className": "root",
	}, List(98), react.Div(nil, List(82)))
}

type SimpleComponent struct{}

func (s *SimpleComponent) Render(ctx *react.Context) *react.Element {
	return react.Span(nil, react.Text("Hello!"))
}

func Greeter(ctx *react.Context) *react.Element {
	return react.H1(nil, react.Text("Hello %s!", ctx.GetString("name")))
}

func main() {
	html := react.Render(react.Component(Greeter), map[string]interface{}{
		"name": "Joe",
	})
	// html := react.RenderWithStats(custom)
	println(html)
}
