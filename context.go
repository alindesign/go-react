package react

import (
	"log"
	"sync"
	"time"
)

var ctxl sync.Mutex

type Context struct {
	data  map[string]any
	nodes int
	start time.Time
}

func (c *Context) CountNode() {
	ctxl.Lock()
	c.nodes = c.nodes + 1
	ctxl.Unlock()
}

func (c *Context) SetData(data map[string]any) {
	c.data = data
}

func (c *Context) Has(key string) bool {
	_, ok := c.data[key]
	return ok
}

func (c *Context) Get(key string, defaultValue ...any) any {
	value, ok := c.data[key]

	if ok {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return nil
}

func (c *Context) GetString(key string, defaultValue ...string) string {
	defVal := ""

	if len(defaultValue) > 0 {
		defVal = defaultValue[0]
	}

	return c.Get(key, defVal).(string)
}

func (c *Context) GetBool(key string, defaultValue ...bool) bool {
	defVal := false

	if len(defaultValue) > 0 {
		defVal = defaultValue[0]
	}

	return c.Get(key, defVal).(bool)
}

func (c *Context) Set(key string, data any) *Context {
	c.data[key] = data
	return c
}

func (c *Context) Unset(key string) *Context {
	delete(c.data, key)
	return c
}

func (c *Context) Stats() {
	diff := time.Now().Sub(c.start)
	log.Printf("go-react: %d nodes rendered in %s", c.nodes, diff)
}

func NewContextWithData(data map[string]any) *Context {
	return &Context{data: data}
}

func NewContext() *Context {
	return &Context{
		data:  map[string]any{},
		nodes: 0,
		start: time.Now(),
	}
}
