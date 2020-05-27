package react

type ComponentFunc func(ctx *Context) *Element
type ComponentStruct interface {
	Render(ctx *Context) *Element
}

func Component(component interface{}) *Element {
	switch x := component.(type) {
	case func(ctx *Context) *Element:
		return &Element{
			Type:   TYPE_COMPONENT,
			render: x,
		}
	case *ComponentStruct:
	case ComponentStruct:
		return &Element{
			Type:      TYPE_COMPONENT_STRUCT,
			component: x,
		}
	}

	return nil
}
