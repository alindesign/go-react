package react

func IsFragment(value interface{}) bool {
	if value == nil {
		return false
	}
	v, ok := value.(*Element)

	if v == nil {
		return false
	}

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
