package react

func Meta(name, content string, props *Props) *Element {
	p := &Props{}

	if props != nil {
		p = props
	}

	if name != "" {
		p.Name = name
	}

	if content != "" {
		p.Content = content
	}

	return CreateElement("meta", p)
}

func MetaItemProp(name, content string, props *Props) *Element {
	p := &Props{}

	if props != nil {
		p = props
	}

	if name != "" {
		p.Itemprop = name
	}

	return Meta("", content, p)
}

func MetaProperty(name, content string, props *Props) *Element {
	p := &Props{}

	if props != nil {
		p = props
	}

	if name != "" {
		p.Property = name
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

	return Meta("", "", &Props{Charset: cset})
}
