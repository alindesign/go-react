package react

func IsFragment(value interface{}) bool {
	v, ok := value.(*Element)
	return ok && v.Fragment
}

func Fragment(childs ...interface{}) *Element {
	return &Element{
		Childs:   childs,
		Fragment: true,
	}
}

func CreateFragment(childs ...interface{}) *Element {
	return Fragment(childs...)
}
