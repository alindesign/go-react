package react

import (
	"reflect"
)

type ComponentStruct interface {
	Render(ctx *Context) interface{}
}

func IsComponent(value interface{}) bool {
	ok := false

	if value != nil {
		rVal := reflect.TypeOf(value)
		_, ok = rVal.MethodByName("Render")
	}

	return ok
}

func renderComponent(value interface{}, ctx interface{}) interface{} {
	if !IsComponent(value) {
		return nil
	}

	rVal := reflect.ValueOf(value)
	method := rVal.MethodByName("Render")

	if method.IsZero() || method.IsNil() {
		return nil
	}

	var result []reflect.Value

	if ctx != nil {
		result = method.Call([]reflect.Value{reflect.ValueOf(ctx)})
	} else {
		result = method.Call([]reflect.Value{reflect.ValueOf(NewContext())})
	}

	return result[0].Interface()
}
