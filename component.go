package react

type ComponentClass interface {
	Render(ctx *Context) interface{}
}
