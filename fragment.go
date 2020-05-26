package react

func Fragment(childs ...*Element) *Element {
	return &Element{
		Type:   TYPE_FRAGMENT,
		Childs: childs,
	}
}

func CreateFragment(childs ...*Element) *Element {
	return Fragment(childs...)
}
