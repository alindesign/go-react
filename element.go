package react

type Element struct {
	Tag    string
	Props  *Props
	Childs []interface{}
}

func CreateElement(tag string, props *Props, childs ...interface{}) *Element {
	return &Element{tag, props, childs}
}
