package react

import (
	"html/template"
)

type Element struct {
	Type      int
	Tag       string
	Props     Props
	Childs    []*Element
	Text      string
	CSS       template.CSS
	JS        template.JS
	HTML      template.HTML
	render    ComponentFunc
	component ComponentStruct
}

func CreateElement(tag string, props Props, childs ...*Element) *Element {
	return &Element{
		Type:   TYPE_ELEMENT,
		Tag:    tag,
		Props:  props,
		Childs: childs,
	}
}

func (e *Element) isVoidElement() bool {
	if e.Type == TYPE_ELEMENT {
		return e.Tag == "area" ||
			e.Tag == "base" ||
			e.Tag == "br" ||
			e.Tag == "col" ||
			e.Tag == "embed" ||
			e.Tag == "hr" ||
			e.Tag == "img" ||
			e.Tag == "input" ||
			e.Tag == "link" ||
			e.Tag == "meta" ||
			e.Tag == "param" ||
			e.Tag == "source" ||
			e.Tag == "track" ||
			e.Tag == "wbr"
	}

	return false
}
