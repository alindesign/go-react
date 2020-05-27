package react

import (
	"github.com/iancoleman/strcase"
	"strings"
)

type Style map[string]string

func (s Style) String() string {
	var attrs strings.Builder

	for key, value := range s {
		attr := s.property(key)

		if value != "" {
			attrs.WriteString(attr)
			attrs.WriteString(":")
			attrs.WriteString(value)
			attrs.WriteString(";")
		}
	}

	return attrs.String()
}

func (s Style) property(key string) string {
	return strcase.ToKebab(key)
}
