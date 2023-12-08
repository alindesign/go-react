package react

import (
	"math"
	"strings"
	"sync"
)

type Renderer struct {
	Context             *Context
	Component           *Element
	MaxConcurrentChunks int
}

func (r *Renderer) String() string {
	return r.render(r.Component)
}

func (r *Renderer) render(component *Element) string {
	r.Context.CountNode()

	if component == nil {
		return ""
	}

	switch component.Type {
	case TYPE_TEXT:
		return component.Text
	case TYPE_CSS:
		return string(component.CSS)
	case TYPE_JS:
		return string(component.JS)
	case TYPE_HTML:
		return string(component.HTML)
	case TYPE_FRAGMENT:
		return r.renderChildren(component)
	case TYPE_ELEMENT:
		return r.renderTag(component)
	case TYPE_COMPONENT:
		return r.render(component.render(r.getContext()))
	case TYPE_COMPONENT_STRUCT:
		return r.render(component.component.Render(r.getContext()))
	default:
		return ""
	}
}

func (r *Renderer) childChunks(component *Element) [][]*Element {
	childrenLen := len(component.Childs)
	possibleChunks := int(math.Ceil(float64(childrenLen) / float64(r.MaxConcurrentChunks)))
	chunks := make([][]*Element, possibleChunks)

	var chunk = 0
	var count = 0

	for _, el := range component.Childs {
		if count < r.MaxConcurrentChunks {
			chunks[chunk] = append(chunks[chunk], el)
			count += 1
		} else {
			count = 0
			chunk += 1
		}
	}

	return chunks
}

func (r *Renderer) renderChildren(component *Element) string {
	if component == nil || component.Childs == nil {
		return ""
	}

	chunksLength := len(component.Childs)
	if chunksLength == 0 {
		return ""
	}

	children := r.childChunks(component)
	chunks := make([]strings.Builder, len(children))

	var wg sync.WaitGroup
	wg.Add(len(children))

	for i, child := range children {
		go func(i int, child []*Element) {
			var group strings.Builder

			for _, element := range child {
				if element != nil {
					group.WriteString(r.render(element))
				}
			}

			chunks[i] = group
			wg.Done()
		}(i, child)
	}

	wg.Wait()

	return r.joinChunks(chunks)
}

func (r *Renderer) renderTag(component *Element) string {
	if component == nil {
		return ""
	}

	var isVoid = component.isVoidElement()
	var children strings.Builder
	var start strings.Builder
	var end strings.Builder

	start.WriteString("<")
	start.WriteString(component.Tag)

	if isVoid {
		end.WriteString("/>")
	} else {
		end.WriteString("</")
		end.WriteString(component.Tag)
		end.WriteString(">")
	}

	if !isVoid {
		children.WriteString(r.renderChildren(component))
	}

	if component.Tag == "" {
		return children.String()
	}

	if component.Props != nil {
		start.WriteString(component.Props.String())
	}

	if !isVoid {
		start.WriteString(">")
	}

	start.WriteString(children.String())
	start.WriteString(end.String())

	return start.String()
}

func (r *Renderer) joinChunks(chunks []strings.Builder) string {
	var content strings.Builder

	for _, chunk := range chunks {
		content.WriteString(chunk.String())
	}

	return content.String()
}

func (r *Renderer) getContext() *Context {
	if r.Context == nil {
		r.Context = NewContext()
	}

	return r.Context
}

func (r *Renderer) SetContext(ctx *Context) *Renderer {
	r.Context = ctx
	return r
}

func (r *Renderer) SetData(data ...map[string]any) *Renderer {
	if len(data) > 0 {
		r.Context.SetData(data[0])
	}
	return r
}

func (r *Renderer) Stats() {
	r.Context.Stats()
}

func (r *Renderer) Bytes() []byte {
	return []byte(r.String())
}

func NewRenderer(component *Element) *Renderer {
	return &Renderer{
		Context:             NewContext(),
		Component:           component,
		MaxConcurrentChunks: 32,
	}
}
