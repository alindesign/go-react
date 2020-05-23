package react

import (
	"fmt"
	"html/template"
)

type Element struct {
	Tag      string
	Props    *Props
	Childs   []interface{}
	Fragment bool
}

func IsElement(value interface{}) bool {
	_, isElement := value.(*Element)
	_, isString := value.(string)
	_, isHtml := value.(template.HTML)
	_, isCss := value.(template.CSS)
	_, isJs := value.(template.JS)

	return IsFragment(value) || isElement || isString || isHtml || isCss || isJs
}

func CheckElement(value interface{}) error {
	if value != nil && !IsElement(value) {
		return fmt.Errorf(`Invalid element, child should be one of: 
	- *react.Element
	- template.HTML
	- template.CSS
	- template.JS
	- string
	- nil

	Received: '%T'
`, value)
	}

	return nil
}

func CreateElement(tag string, props *Props, childs ...interface{}) *Element {
	return &Element{Tag: tag, Props: props, Childs: childs}
}
