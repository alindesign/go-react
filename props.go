package react

import (
	"github.com/iancoleman/strcase"
	"github.com/spf13/cast"
	"html/template"
	"sort"
	"strings"
)

type Props struct {
	Accept          interface{}
	Acceptcharset   interface{}
	Accesskey       interface{}
	Action          interface{}
	Align           interface{}
	Allow           interface{}
	Alt             interface{}
	Async           interface{}
	Autocapitalize  interface{}
	Autocomplete    interface{}
	Autofocus       interface{}
	Autoplay        interface{}
	Background      interface{}
	Border          interface{}
	Buffered        interface{}
	Capture         interface{}
	Charset         interface{}
	Checked         interface{}
	ClassName       interface{}
	Cols            interface{}
	Colspan         interface{}
	Content         interface{}
	Contenteditable interface{}
	Contextmenu     interface{}
	Controls        interface{}
	Coords          interface{}
	Crossorigin     interface{}
	Csp             interface{}
	Datetime        interface{}
	Defer           interface{}
	Disabled        interface{}
	Download        interface{}
	Draggable       interface{}
	Dropzone        interface{}
	Enctype         interface{}
	For             interface{}
	Form            interface{}
	Formaction      interface{}
	Formenctype     interface{}
	Formmethod      interface{}
	Formnovalidate  interface{}
	Formtarget      interface{}
	Headers         interface{}
	Height          interface{}
	Hidden          interface{}
	Href            interface{}
	Hreflang        interface{}
	Id              interface{}
	Integrity       interface{}
	Itemprop        interface{}
	Label           interface{}
	Lang            interface{}
	Language        interface{}
	Loading         interface{}
	Max             interface{}
	Maxlength       interface{}
	Minlength       interface{}
	Media           interface{}
	Method          interface{}
	Min             interface{}
	Multiple        interface{}
	Muted           interface{}
	Name            interface{}
	Novalidate      interface{}
	Open            interface{}
	Pattern         interface{}
	Placeholder     interface{}
	Poster          interface{}
	Preload         interface{}
	Readonly        interface{}
	Referrerpolicy  interface{}
	Rel             interface{}
	Required        interface{}
	Reversed        interface{}
	Rows            interface{}
	Rowspan         interface{}
	Sandbox         interface{}
	Scope           interface{}
	Scoped          interface{}
	Selected        interface{}
	Shape           interface{}
	Size            interface{}
	Sizes           interface{}
	Slot            interface{}
	Span            interface{}
	Spellcheck      interface{}
	Src             interface{}
	Srcdoc          interface{}
	Srclang         interface{}
	Srcset          interface{}
	Start           interface{}
	Step            interface{}
	Style           *Style
	Summary         interface{}
	Tabindex        interface{}
	Target          interface{}
	Title           interface{}
	Type            interface{}
	Value           interface{}
	Width           interface{}
	Wrap            interface{}
	Data            interface{}

	CustomProps interface{}
	Property    interface{}
	HttpEquiv   interface{}
}

func NewProps() *Props {
	return &Props{}
}

func (p *Props) String() string {
	if p == nil {
		return ""
	}

	values := toMap(p)
	attrs := map[string]string{}

	p.attrs(values, func(key, value string) {
		attrs[key] = value
	})

	if p.CustomProps != nil {
		p.attrs(toMap(p.CustomProps), func(key, value string) {
			attrs[key] = value
		})
	}

	return strings.Join(p.toSlice(attrs), " ")
}

func (p *Props) attrs(values map[string]interface{}, fn func(key, val string)) {
	for key, value := range values {
		if p.omit(key, value) {
			continue
		}

		attr := p.getAttrName(key)
		strValue := attr

		switch val := value.(type) {
		case *Style:
			strValue = val.String()
			break
		case []string:
			strValue = strings.Join(val, " ")
			break
		default:
			strValue = template.HTMLEscapeString(cast.ToString(value))
			break
		}

		fn(attr, strValue)
	}
}

func (p *Props) getAttrName(key string) string {
	if key == "ClassName" {
		return "class"
	}

	return strcase.ToKebab(key)
}

func (p *Props) omit(key string, value interface{}) bool {
	if key == "CustomProps" {
		return true
	}

	return false
}

func (p *Props) toSlice(attrs map[string]string) []string {
	var slice []string

	for key, value := range attrs {
		slice = append(slice, key+`="`+value+`"`)
	}

	sort.Strings(slice)

	return slice
}
