package main

import (
	"fmt"
	"github.com/alindesign/go-react"
)

func main() {
	content, err := react.Render(react.CreateElement("span", nil, 1024))

	fmt.Printf("content: %s,\nerr: %v", content, err)
}
