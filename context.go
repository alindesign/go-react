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

func (c *Context) Get(key string, defaultValue interface{}) interface{} {
	value, ok := c.data[key]

	if ok {
		return value
	}

	return defaultValue
}

func (c *Context) GetString(key string, defaultValue string) string {
	return c.Get(key, defaultValue).(string)
}

func (c *Context) GetBool(key string, defaultValue bool) bool {
	return c.Get(key, defaultValue).(bool)
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
