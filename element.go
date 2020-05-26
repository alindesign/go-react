package react

import (
	"html/template"
)

var voidElements = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}

type Element struct {
	Type   int
	Tag    string
	Props  *Props
	Childs []*Element
	Text   string
	CSS    template.CSS
	JS     template.JS
	HTML   template.HTML
}

func CreateElement(tag string, props *Props, childs ...*Element) *Element {
	return &Element{
		Type:   TYPE_ELEMENT,
		Tag:    tag,
		Props:  props,
		Childs: childs,
	}
}

func (e *Element) isVoidElement() bool {
	found := false

	if e.Type == TYPE_ELEMENT {
		for _, voidTag := range voidElements {
			if e.Tag == voidTag {
				found = true
				break
			}
		}
	}

	return found
}
