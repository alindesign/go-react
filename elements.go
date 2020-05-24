package react

import (
	"fmt"
	"html/template"
	"strings"
)

func Comment(comment string, args ...interface{}) *Element {
	return CreateFragment(template.HTML("<!-- " + template.HTMLEscapeString(fmt.Sprintf(comment, args...)) + " -->"))
}

func Doctype(doctype ...interface{}) *Element {
	d := "html"

	if len(doctype) > 0 {
		if doctype[0] != nil {
			dStr, ok := doctype[0].(string)
			if ok && dStr != "" {
				d = dStr
			}

			if strings.HasPrefix(d, "<!doctype") || strings.HasPrefix(d, "<!Doctype") || strings.HasPrefix(d, "<!DOCTYPE") {
				return CreateFragment(template.HTML(d))
			}
		}
	}

	return CreateFragment(template.HTML("<!doctype " + d + ">"))
}

func A(props *Props, children ...interface{}) *Element {
	return CreateElement("a", props, children...)
}

func Abbr(props *Props, children ...interface{}) *Element {
	return CreateElement("abbr", props, children...)
}

func Address(props *Props, children ...interface{}) *Element {
	return CreateElement("address", props, children...)
}

func Area(props *Props, children ...interface{}) *Element {
	return CreateElement("area", props, children...)
}

func Article(props *Props, children ...interface{}) *Element {
	return CreateElement("article", props, children...)
}

func Aside(props *Props, children ...interface{}) *Element {
	return CreateElement("aside", props, children...)
}

func Audio(props *Props, children ...interface{}) *Element {
	return CreateElement("audio", props, children...)
}

func B(props *Props, children ...interface{}) *Element {
	return CreateElement("b", props, children...)
}

func Base(props *Props, children ...interface{}) *Element {
	return CreateElement("base", props, children...)
}

func Bdi(props *Props, children ...interface{}) *Element {
	return CreateElement("bdi", props, children...)
}

func Bdo(props *Props, children ...interface{}) *Element {
	return CreateElement("bdo", props, children...)
}

func Blockquote(props *Props, children ...interface{}) *Element {
	return CreateElement("blockquote", props, children...)
}

func Body(props *Props, children ...interface{}) *Element {
	return CreateElement("body", props, children...)
}

func Br() *Element {
	return CreateElement("br", nil)
}

func Button(props *Props, children ...interface{}) *Element {
	return CreateElement("button", props, children...)
}

func Canvas(props *Props, children ...interface{}) *Element {
	return CreateElement("canvas", props, children...)
}

func Caption(props *Props, children ...interface{}) *Element {
	return CreateElement("caption", props, children...)
}

func Cite(props *Props, children ...interface{}) *Element {
	return CreateElement("cite", props, children...)
}

func Code(props *Props, children ...interface{}) *Element {
	return CreateElement("code", props, children...)
}

func Col(props *Props, children ...interface{}) *Element {
	return CreateElement("col", props, children...)
}

func Colgroup(props *Props, children ...interface{}) *Element {
	return CreateElement("colgroup", props, children...)
}

func Data(props *Props, children ...interface{}) *Element {
	return CreateElement("data", props, children...)
}

func Datalist(props *Props, children ...interface{}) *Element {
	return CreateElement("datalist", props, children...)
}

func Dd(props *Props, children ...interface{}) *Element {
	return CreateElement("dd", props, children...)
}

func Del(props *Props, children ...interface{}) *Element {
	return CreateElement("del", props, children...)
}

func Details(props *Props, children ...interface{}) *Element {
	return CreateElement("details", props, children...)
}

func Dfn(props *Props, children ...interface{}) *Element {
	return CreateElement("dfn", props, children...)
}

func Dialog(props *Props, children ...interface{}) *Element {
	return CreateElement("dialog", props, children...)
}

func Div(props *Props, children ...interface{}) *Element {
	return CreateElement("div", props, children...)
}

func Dl(props *Props, children ...interface{}) *Element {
	return CreateElement("dl", props, children...)
}

func Dt(props *Props, children ...interface{}) *Element {
	return CreateElement("dt", props, children...)
}

func Em(props *Props, children ...interface{}) *Element {
	return CreateElement("em", props, children...)
}

func Embed(props *Props, children ...interface{}) *Element {
	return CreateElement("embed", props, children...)
}

func Fieldset(props *Props, children ...interface{}) *Element {
	return CreateElement("fieldset", props, children...)
}

func Figure(props *Props, children ...interface{}) *Element {
	return CreateElement("figure", props, children...)
}

func Footer(props *Props, children ...interface{}) *Element {
	return CreateElement("footer", props, children...)
}

func Form(props *Props, children ...interface{}) *Element {
	return CreateElement("form", props, children...)
}

func H1(props *Props, children ...interface{}) *Element {
	return CreateElement("h1", props, children...)
}

func H2(props *Props, children ...interface{}) *Element {
	return CreateElement("h2", props, children...)
}

func H3(props *Props, children ...interface{}) *Element {
	return CreateElement("h3", props, children...)
}

func H4(props *Props, children ...interface{}) *Element {
	return CreateElement("h4", props, children...)
}

func H5(props *Props, children ...interface{}) *Element {
	return CreateElement("h5", props, children...)
}

func H6(props *Props, children ...interface{}) *Element {
	return CreateElement("h6", props, children...)
}

func Head(props *Props, children ...interface{}) *Element {
	return CreateElement("head", props, children...)
}

func Header(props *Props, children ...interface{}) *Element {
	return CreateElement("header", props, children...)
}

func Hgroup(props *Props, children ...interface{}) *Element {
	return CreateElement("hgroup", props, children...)
}

func Hr(props ...*Props) *Element {
	var p *Props

	if len(props) > 0 {
		p = props[0]
	}

	return CreateElement("hr", p)
}

func Html(props *Props, children ...interface{}) *Element {
	return CreateElement("html", props, children...)
}

func I(props *Props, children ...interface{}) *Element {
	return CreateElement("i", props, children...)
}

func Iframe(props *Props, children ...interface{}) *Element {
	return CreateElement("iframe", props, children...)
}

func Img(src string, alt string, props ...*Props) *Element {
	p := &Props{}

	if len(props) > 0 && props[0] != nil {
		p = props[0]
	}

	p.Src = src
	p.Alt = alt

	return CreateElement("img", p)
}

func Input(props *Props) *Element {
	return CreateElement("input", props)
}

func Ins(props *Props, children ...interface{}) *Element {
	return CreateElement("ins", props, children...)
}

func Kbd(props *Props, children ...interface{}) *Element {
	return CreateElement("kbd", props, children...)
}

func Keygen(props *Props, children ...interface{}) *Element {
	return CreateElement("keygen", props, children...)
}

func Label(props *Props, children ...interface{}) *Element {
	return CreateElement("label", props, children...)
}

func Legend(props *Props, children ...interface{}) *Element {
	return CreateElement("legend", props, children...)
}

func Li(props *Props, children ...interface{}) *Element {
	return CreateElement("li", props, children...)
}

func Link(rel string, href string, props *Props) *Element {
	p := &Props{}

	if props != nil {
		p = props
	}

	if rel != "" {
		p.Rel = rel
	}

	if href != "" {
		p.Href = href
	}

	return CreateElement("link", props)
}

func Main(props *Props, children ...interface{}) *Element {
	return CreateElement("main", props, children...)
}

func Map(props *Props, children ...interface{}) *Element {
	return CreateElement("map", props, children...)
}

func Mark(props *Props, children ...interface{}) *Element {
	return CreateElement("mark", props, children...)
}

func Menu(props *Props, children ...interface{}) *Element {
	return CreateElement("menu", props, children...)
}

func Menuitem(props *Props, children ...interface{}) *Element {
	return CreateElement("menuitem", props, children...)
}

func Meter(props *Props, children ...interface{}) *Element {
	return CreateElement("meter", props, children...)
}

func Nav(props *Props, children ...interface{}) *Element {
	return CreateElement("nav", props, children...)
}

func Noscript(children ...interface{}) *Element {
	return CreateElement("noscript", nil, children...)
}

func Object(props *Props, children ...interface{}) *Element {
	return CreateElement("object", props, children...)
}

func Ol(props *Props, children ...interface{}) *Element {
	return CreateElement("ol", props, children...)
}

func Optgroup(props *Props, children ...interface{}) *Element {
	return CreateElement("optgroup", props, children...)
}

func Option(props *Props, children ...interface{}) *Element {
	return CreateElement("option", props, children...)
}

func Output(props *Props, children ...interface{}) *Element {
	return CreateElement("output", props, children...)
}

func P(props *Props, children ...interface{}) *Element {
	return CreateElement("p", props, children...)
}

func Param(props *Props, children ...interface{}) *Element {
	return CreateElement("param", props, children...)
}

func Pre(props *Props, children ...interface{}) *Element {
	return CreateElement("pre", props, children...)
}

func Progress(props *Props, children ...interface{}) *Element {
	return CreateElement("progress", props, children...)
}

func Q(props *Props, children ...interface{}) *Element {
	return CreateElement("q", props, children...)
}

func Rb(props *Props, children ...interface{}) *Element {
	return CreateElement("rb", props, children...)
}

func Rp(props *Props, children ...interface{}) *Element {
	return CreateElement("rp", props, children...)
}

func Rt(props *Props, children ...interface{}) *Element {
	return CreateElement("rt", props, children...)
}

func Rtc(props *Props, children ...interface{}) *Element {
	return CreateElement("rtc", props, children...)
}

func Ruby(props *Props, children ...interface{}) *Element {
	return CreateElement("ruby", props, children...)
}

func S(props *Props, children ...interface{}) *Element {
	return CreateElement("s", props, children...)
}

func Samp(props *Props, children ...interface{}) *Element {
	return CreateElement("samp", props, children...)
}

func Script(source interface{}, props *Props) *Element {
	var children interface{}
	p := &Props{}

	if props != nil {
		p = props
	}

	if source != nil && source != "" {
		switch src := source.(type) {
		case string:
			p.Src = src
			break
		case template.URL:
			p.Src = string(src)
			break

		case template.JS:
			children = src
			break
		}
	}

	if p.Type == nil {
		p.Type = "application/javascript"
	}

	return CreateElement("script", p, children)
}

func Section(props *Props, children ...interface{}) *Element {
	return CreateElement("section", props, children...)
}

func Select(props *Props, children ...interface{}) *Element {
	return CreateElement("select", props, children...)
}

func Small(props *Props, children ...interface{}) *Element {
	return CreateElement("small", props, children...)
}

func Source(props *Props, children ...interface{}) *Element {
	return CreateElement("source", props, children...)
}

func Span(props *Props, children ...interface{}) *Element {
	return CreateElement("span", props, children...)
}

func Strong(props *Props, children ...interface{}) *Element {
	return CreateElement("strong", props, children...)
}

func StyleTag(style template.CSS, props ...*Props) *Element {
	p := &Props{}

	if len(props) > 0 && props[0] != nil {
		p = props[0]
	}

	if p.Type == nil {
		p.Type = "text/css"
	}

	return CreateElement("style", p, style)
}

func Sub(props *Props, children ...interface{}) *Element {
	return CreateElement("sub", props, children...)
}

func Summary(props *Props, children ...interface{}) *Element {
	return CreateElement("summary", props, children...)
}

func Sup(props *Props, children ...interface{}) *Element {
	return CreateElement("sup", props, children...)
}

func Table(props *Props, children ...interface{}) *Element {
	return CreateElement("table", props, children...)
}

func Tbody(props *Props, children ...interface{}) *Element {
	return CreateElement("tbody", props, children...)
}

func Td(props *Props, children ...interface{}) *Element {
	return CreateElement("td", props, children...)
}

func Template(props *Props, children ...interface{}) *Element {
	return CreateElement("template", props, children...)
}

func Textarea(props *Props, children ...interface{}) *Element {
	return CreateElement("textarea", props, children...)
}

func Tfoot(props *Props, children ...interface{}) *Element {
	return CreateElement("tfoot", props, children...)
}

func Th(props *Props, children ...interface{}) *Element {
	return CreateElement("th", props, children...)
}

func Thead(props *Props, children ...interface{}) *Element {
	return CreateElement("thead", props, children...)
}

func Time(props *Props, children ...interface{}) *Element {
	return CreateElement("time", props, children...)
}

func Title(title string, args ...interface{}) *Element {
	return CreateElement("title", nil, fmt.Sprintf(title, args...))
}

func Tr(props *Props, children ...interface{}) *Element {
	return CreateElement("tr", props, children...)
}

func Track(props *Props, children ...interface{}) *Element {
	return CreateElement("track", props, children...)
}

func U(props *Props, children ...interface{}) *Element {
	return CreateElement("u", props, children...)
}

func Ul(props *Props, children ...interface{}) *Element {
	return CreateElement("ul", props, children...)
}

func Var(props *Props, children ...interface{}) *Element {
	return CreateElement("var", props, children...)
}

func Video(props *Props, children ...interface{}) *Element {
	return CreateElement("video", props, children...)
}

func Wbr(props *Props, children ...interface{}) *Element {
	return CreateElement("wbr", props, children...)
}
