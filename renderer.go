package react

import (
	"strings"
	"sync"
)

const indentSize = "    "

type Renderer struct {
	component *Element
	context   *Context
	pretty    bool
}

func (r *Renderer) String() string {

	return r.render(r.component, 0)
}

func (r *Renderer) render(component *Element, level int) string {
	if component == nil {
		return ""
	}

	if component.Type == TYPE_TEXT {
		return component.Text
	} else if component.Type == TYPE_CSS {
		return string(component.CSS)
	} else if component.Type == TYPE_JS {
		return string(component.JS)
	} else if component.Type == TYPE_HTML {
		return string(component.HTML)
	} else if component.Type == TYPE_FRAGMENT {
		return r.renderChilds(component, level)
	} else if component.Type == TYPE_ELEMENT {
		return r.renderTag(component, level)
	}

	return ""
}

func (r *Renderer) renderChilds(component *Element, level int) string {
	if component == nil || component.Childs == nil {
		return ""
	}

	chunksLength := len(component.Childs)
	if chunksLength == 0 {
		return ""
	}

	chunks := make([]string, chunksLength)
	var wg sync.WaitGroup
	wg.Add(chunksLength)

	for i, child := range component.Childs {
		go func(i int, child *Element) {
			if child != nil {
				chunks[i] = r.render(child, level+1)
			} else {
				chunks[i] = ""
			}

			wg.Done()
		}(i, child)
	}

	wg.Wait()

	return r.joinChunks(chunks, level)
}

func (r *Renderer) renderTag(component *Element, level int) string {
	if component == nil {
		return ""
	}

	children := ""
	start := "<" + component.Tag
	end := "</" + component.Tag + ">"
	isVoid := component.isVoidElement()

	if isVoid {
		end = " />"
	} else {
		children = r.renderChilds(component, level+1)
	}

	if component.Tag == "" {
		return children
	}

	if component.Props != nil {
		attrs := component.Props.String()
		if attrs != "" {
			start += " " + attrs
		}
	}

	if !isVoid {
		start += ">"
	}

	return start + children + end
}

func (r *Renderer) joinChunks(chunks []string, level int) string {
	indent := ""
	if r.pretty {
		indent = "\n" + strings.Repeat(indentSize, level)
	}

	return indent + strings.Join(chunks, indent)
}

func NewRenderer(component *Element) *Renderer {
	return &Renderer{component: component}
}
