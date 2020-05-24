package react

import (
	"errors"
	"html/template"
	"sync"
)

var voidElements = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}

func isVoidElement(element string) bool {
	for _, el := range voidElements {
		if el == element {
			return true
		}
	}

	return false
}

func genHtml(tag string, attrs *Props, childs []string) string {
	attrsStr := ""
	content := ""
	hasTag := tag != ""
	isVoid := isVoidElement(tag)

	if hasTag {
		attrsStr = attrs.String()
		if attrsStr != "" {
			attrsStr = " " + attrsStr
		}
	}

	if hasTag && isVoid {
		return "<" + tag + attrsStr + " />"
	}

	for _, child := range childs {
		content += child
	}

	if hasTag {
		content = "<" + tag + attrsStr + ">" + content + "</" + tag + ">"
	}

	return content
}

func render(component *Element, ctx *Context) (string, error) {
	if component == nil {
		return "", nil
	}

	if component.Tag == "" && !component.Fragment {
		return "", errors.New("component has an empty tag")
	}

	chunksLen := len(component.Childs)
	var chunks []string
	var err error

	for _, value := range component.Childs {
		err = CheckElement(value)
		if err != nil {
			return "", err
		}
	}

	if component.Childs != nil && chunksLen > 0 {
		chunks = make([]string, chunksLen)
		var wg sync.WaitGroup

		for i, child := range component.Childs {
			wg.Add(1)

			go func(j int, c interface{}) {
				var chunk string
				var errChunk error

				if c != nil {
					if IsComponent(c) {
						chunk, errChunk = renderElement(renderComponent(c, ctx), ctx)
					} else {
						chunk, errChunk = renderElement(c, ctx)
					}
				}

				if errChunk != nil {
					err = errChunk
					wg.Done()
					return
				}

				chunks[j] = chunk
				wg.Done()
			}(i, child)
		}

		// Wait until all childs are rendered
		wg.Wait()

		if err != nil {
			return "", err
		}
	}

	return genHtml(component.Tag, component.Props, chunks), nil
}

func renderElement(element interface{}, ctx *Context) (string, error) {
	switch el := element.(type) {
	case *Element:
		return render(el, ctx)
	case template.HTML:
		return string(el), nil
	case template.CSS:
		return string(el), nil
	case template.JS:
		return string(el), nil
	case string:
		return template.HTMLEscapeString(el), nil
	}

	return "", nil
}

func Render(component *Element, data ...map[string]interface{}) (string, error) {
	var err error
	var content string

	ctx := NewContext()
	if len(data) > 0 {
		ctx.SetData(data[0])
	}

	content, err = render(component, ctx)

	return content, err
}

func RenderBytes(component *Element, data ...map[string]interface{}) ([]byte, error) {
	content, err := Render(component, data...)

	if err != nil {
		return nil, err
	}

	return []byte(content), nil
}
