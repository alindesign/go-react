package react

func Meta(name, content string, props Props) *Element {
	p := Props{}

	if props != nil {
		p = props
	}

	if name != "" {
		p["name"] = name
	}

	if content != "" {
		p["content"] = content
	}

	return CreateElement("meta", p)
}

func MetaItemProp(name, content string, props Props) *Element {
	p := Props{}

	if props != nil {
		p = props
	}

	if name != "" {
		p["itemprop"] = name
	}

	return Meta("", content, p)
}

func MetaProperty(name, content string, props Props) *Element {
	p := Props{}

	if props != nil {
		p = props
	}

	if name != "" {
		p["property"] = name
	}

	return Meta("", content, p)
}

func MetaCharset(charset ...string) *Element {
	cset := "UTF-8"

	if len(charset) > 0 {
		if charset[0] != "" {
			cset = charset[0]
		}
	}

	return Meta("", "", Props{"charset": cset})
}
