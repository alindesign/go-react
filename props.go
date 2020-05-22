package react

import (
	"github.com/fatih/structs"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cast"
	"html/template"
	"strings"
)

type Props struct {
	Accept          interface{} `structs:",omitempty"`
	Acceptcharset   interface{} `structs:",omitempty"`
	Accesskey       interface{} `structs:",omitempty"`
	Action          interface{} `structs:",omitempty"`
	Align           interface{} `structs:",omitempty"`
	Allow           interface{} `structs:",omitempty"`
	Alt             interface{} `structs:",omitempty"`
	Async           interface{} `structs:",omitempty"`
	Autocapitalize  interface{} `structs:",omitempty"`
	Autocomplete    interface{} `structs:",omitempty"`
	Autofocus       interface{} `structs:",omitempty"`
	Autoplay        interface{} `structs:",omitempty"`
	Background      interface{} `structs:",omitempty"`
	Border          interface{} `structs:",omitempty"`
	Buffered        interface{} `structs:",omitempty"`
	Capture         interface{} `structs:",omitempty"`
	Charset         interface{} `structs:",omitempty"`
	Checked         interface{} `structs:",omitempty"`
	ClassName       interface{} `structs:",omitempty"`
	Cols            interface{} `structs:",omitempty"`
	Colspan         interface{} `structs:",omitempty"`
	Content         interface{} `structs:",omitempty"`
	Contenteditable interface{} `structs:",omitempty"`
	Contextmenu     interface{} `structs:",omitempty"`
	Controls        interface{} `structs:",omitempty"`
	Coords          interface{} `structs:",omitempty"`
	Crossorigin     interface{} `structs:",omitempty"`
	Csp             interface{} `structs:",omitempty"`
	Datetime        interface{} `structs:",omitempty"`
	Defer           interface{} `structs:",omitempty"`
	Disabled        interface{} `structs:",omitempty"`
	Download        interface{} `structs:",omitempty"`
	Draggable       interface{} `structs:",omitempty"`
	Dropzone        interface{} `structs:",omitempty"`
	Enctype         interface{} `structs:",omitempty"`
	For             interface{} `structs:",omitempty"`
	Form            interface{} `structs:",omitempty"`
	Formaction      interface{} `structs:",omitempty"`
	Formenctype     interface{} `structs:",omitempty"`
	Formmethod      interface{} `structs:",omitempty"`
	Formnovalidate  interface{} `structs:",omitempty"`
	Formtarget      interface{} `structs:",omitempty"`
	Headers         interface{} `structs:",omitempty"`
	Height          interface{} `structs:",omitempty"`
	Hidden          interface{} `structs:",omitempty"`
	Href            interface{} `structs:",omitempty"`
	Hreflang        interface{} `structs:",omitempty"`
	Id              interface{} `structs:",omitempty"`
	Integrity       interface{} `structs:",omitempty"`
	Itemprop        interface{} `structs:",omitempty"`
	Label           interface{} `structs:",omitempty"`
	Lang            interface{} `structs:",omitempty"`
	Language        interface{} `structs:",omitempty"`
	Loading         interface{} `structs:",omitempty"`
	Max             interface{} `structs:",omitempty"`
	Maxlength       interface{} `structs:",omitempty"`
	Minlength       interface{} `structs:",omitempty"`
	Media           interface{} `structs:",omitempty"`
	Method          interface{} `structs:",omitempty"`
	Min             interface{} `structs:",omitempty"`
	Multiple        interface{} `structs:",omitempty"`
	Muted           interface{} `structs:",omitempty"`
	Name            interface{} `structs:",omitempty"`
	Novalidate      interface{} `structs:",omitempty"`
	Open            interface{} `structs:",omitempty"`
	Pattern         interface{} `structs:",omitempty"`
	Placeholder     interface{} `structs:",omitempty"`
	Poster          interface{} `structs:",omitempty"`
	Preload         interface{} `structs:",omitempty"`
	Readonly        interface{} `structs:",omitempty"`
	Referrerpolicy  interface{} `structs:",omitempty"`
	Rel             interface{} `structs:",omitempty"`
	Required        interface{} `structs:",omitempty"`
	Reversed        interface{} `structs:",omitempty"`
	Rows            interface{} `structs:",omitempty"`
	Rowspan         interface{} `structs:",omitempty"`
	Sandbox         interface{} `structs:",omitempty"`
	Scope           interface{} `structs:",omitempty"`
	Scoped          interface{} `structs:",omitempty"`
	Selected        interface{} `structs:",omitempty"`
	Shape           interface{} `structs:",omitempty"`
	Size            interface{} `structs:",omitempty"`
	Sizes           interface{} `structs:",omitempty"`
	Slot            interface{} `structs:",omitempty"`
	Span            interface{} `structs:",omitempty"`
	Spellcheck      interface{} `structs:",omitempty"`
	Src             interface{} `structs:",omitempty"`
	Srcdoc          interface{} `structs:",omitempty"`
	Srclang         interface{} `structs:",omitempty"`
	Srcset          interface{} `structs:",omitempty"`
	Start           interface{} `structs:",omitempty"`
	Step            interface{} `structs:",omitempty"`
	Style           *Style      `structs:",omitempty"`
	Summary         interface{} `structs:",omitempty"`
	Tabindex        interface{} `structs:",omitempty"`
	Target          interface{} `structs:",omitempty"`
	Title           interface{} `structs:",omitempty"`
	Type            interface{} `structs:",omitempty"`
	Value           interface{} `structs:",omitempty"`
	Width           interface{} `structs:",omitempty"`
	Wrap            interface{} `structs:",omitempty"`
	Data            interface{} `structs:",omitempty"`

	CustomProps interface{} `structs:",omitempty,omitnested"`
}

func (p *Props) String() string {
	if p == nil {
		return ""
	}

	values := structs.Map(p)
	attrs := map[string]string{}

	p.attrs(values, func(key, value string) {
		attrs[key] = value
	})

	if p.CustomProps != nil {
		p.attrs(structs.Map(p.CustomProps), func(key, value string) {
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

		switch value.(type) {
		case *Style:
			strValue = value.(*Style).String()
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

	return slice
}
