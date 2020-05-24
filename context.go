package react

type Context struct {
	data map[string]interface{}
}

func (c *Context) SetData(data map[string]interface{}) {
	c.data = data
}

func (c *Context) Has(key string) bool {
	_, ok := c.data[key]
	return ok
}

func (c *Context) Get(key string, defaultValue ...interface{}) interface{} {
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

func (c *Context) Set(key string, data map[string]interface{}) *Context {
	c.data[key] = data
	return c
}

func (c *Context) Unset(key string) *Context {
	delete(c.data, key)
	return c
}

func NewContextWithData(data map[string]interface{}) *Context {
	return &Context{data: data}
}

func NewContext() *Context {
	return &Context{data: map[string]interface{}{}}
}
