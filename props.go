package react

import (
	"github.com/iancoleman/strcase"
	"github.com/spf13/cast"
	"html/template"
	"sort"
	"strings"
)

type Props map[string]any

func NewProps() Props {
	return Props{}
}

func (p Props) String() string {
	if p == nil {
		return ""
	}

	var attrs strings.Builder
	var attr string
	keys := p.keys()

	for _, key := range keys {
		attr = p.getAttrName(key)

		attrs.WriteString(" ")
		attrs.WriteString(attr)
		attrs.WriteString("=\"")

		attrs.WriteString(p.Get(key))
		attrs.WriteString("\"")
	}

	return attrs.String()
}

func (p Props) getAttrName(key string) string {
	if key == "className" {
		return "class"
	}

	return strcase.ToKebab(key)
}

func (p Props) keys() []string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

func (p Props) Get(key string) string {
	value, ok := p[key]

	if ok {
		switch v := value.(type) {
		case *Style:
		case Style:
			return v.String()
		case []string:
			return strings.Join(v, " ")
		default:
			return template.HTMLEscapeString(cast.ToString(v))
		}
	}

	return ""
}

func (p Props) Has(key string) bool {
	value, ok := p[key]
	return value != "" && ok
}
