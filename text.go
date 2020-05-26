package react

import (
	"fmt"
	"html/template"
)

func Text(format string, args ...interface{}) *Element {
	return &Element{
		Type: TYPE_TEXT,
		Text: template.HTMLEscapeString(fmt.Sprintf(format, args...)),
	}
}

func CSS(css template.CSS) *Element {
	return &Element{
		Type: TYPE_CSS,
		CSS:  css,
	}
}

func CSSString(css string, args ...interface{}) *Element {
	return CSS(template.CSS(fmt.Sprintf(css, args...)))
}

func JS(js template.JS) *Element {
	return &Element{
		Type: TYPE_JS,
		JS:   js,
	}
}

func JSString(js string, args ...interface{}) *Element {
	return JS(template.JS(fmt.Sprintf(js, args...)))
}

func HTML(html template.HTML) *Element {
	return &Element{
		Type: TYPE_HTML,
		HTML: html,
	}
}

func HTMLString(html string, args ...interface{}) *Element {
	return HTML(template.HTML(fmt.Sprintf(html, args...)))
}
