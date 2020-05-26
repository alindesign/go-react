package react

type DocumentProps struct {
	Head      *Element
	HeadProps *Props

	Body      *Element
	BodyProps *Props

	Doctype string
}

func Document(documentProps *DocumentProps, props *Props) *Element {
	dProps := &DocumentProps{}
	hProps := &Props{}

	if documentProps != nil {
		dProps = documentProps
	}

	if props != nil {
		hProps = props
	}

	if hProps.Lang == "" || hProps.Lang == nil {
		hProps.Lang = "en"
	}

	return CreateFragment(
		Doctype(dProps.Doctype),
		Html(hProps,
			Head(dProps.HeadProps, dProps.Head),
			Body(dProps.BodyProps, dProps.Body),
		),
	)
}
