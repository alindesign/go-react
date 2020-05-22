package main

import (
	"fmt"
	"github.com/alindesign/go-react"
)

func List(num int) *react.Element {
	var items []interface{}
	for i := 0; i < num; i++ {
		items = append(items, react.CreateElement("li", &react.Props{
			ClassName: fmt.Sprintf("list__item item-%d", i+1),
		}, fmt.Sprintf("Item %d", i+1)))
	}

	return react.CreateElement("ul", &react.Props{
		ClassName: "list__root",
	}, items...)
}

func App() *react.Element {
	return react.CreateElement("div", &react.Props{
		ClassName: "root",
	}, List(4))
}

func main() {
	html, err := react.Render(App())
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("Output: %s\n", html)
}
