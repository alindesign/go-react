package react

type ComponentStruct interface {
	Render(ctx *Context) interface{}
}

func IsComponent(value interface{}) bool {
	_, ok := value.(ComponentStruct)

	return ok
}
