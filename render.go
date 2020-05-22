package react

import (
	"errors"
	"github.com/spf13/cast"
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
	attrsStr := attrs.String()
	content := "<" + tag

	if attrsStr != "" {
		content += " " + attrsStr
	}

	if isVoidElement(tag) {
		return content + " />"
	}

	content += ">"
	for _, child := range childs {
		content += child
	}

	return content + "</" + tag + ">"
}

func render(component *Element, ctx *Context) (string, error) {
	if component.Tag == "" {
		return "", errors.New("component has an empty tag")
	}

	chunksLen := len(component.Childs)
	var chunks []string
	var err error

	if component.Childs != nil && chunksLen > 0 {
		chunks = make([]string, chunksLen)
		var wg sync.WaitGroup

		for i, child := range component.Childs {
			wg.Add(1)

			go func(j int, c interface{}) {
				var chunk string
				var errChunk error

				switch c.(type) {
				case *Element:
					chunk, errChunk = render(c.(*Element), ctx)
					break

				case ComponentClass:
					result := c.(ComponentClass).Render(ctx)
					switch result.(type) {
					case *Element:
						chunk, errChunk = render(result.(*Element), ctx)
						break
					default:
						chunk = template.HTMLEscapeString(cast.ToString(result))
						break
					}

					break

				default:
					chunk = template.HTMLEscapeString(cast.ToString(c))
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

func Render(component *Element, data ...map[string]interface{}) (string, error) {
	ctx := NewContext()
	if len(data) > 0 {
		ctx.SetData(data[0])
	}

	return render(component, ctx)
}

func RenderBytes(component *Element, data ...map[string]interface{}) ([]byte, error) {
	content, err := Render(component, data...)

	if err != nil {
		return nil, err
	}

	return []byte(content), nil
}
