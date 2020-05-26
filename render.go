package react

func Render(component *Element, data ...map[string]interface{}) string {
	ctx := NewContext()
	if len(data) > 0 {
		ctx.SetData(data[0])
	}

	renderer := NewRenderer(component)
	return renderer.String()
}

func RenderBytes(component *Element, data ...map[string]interface{}) []byte {
	content := Render(component, data...)
	return []byte(content)
}
