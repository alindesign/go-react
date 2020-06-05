package react

func render(component *Element, data ...map[string]interface{}) *Renderer {
	return NewRenderer(component).SetData(data...)
}

func Render(component *Element, data ...map[string]interface{}) string {
	return render(component, data...).String()
}

func RenderWithStats(component *Element, data ...map[string]interface{}) string {
	r := render(component, data...)
	str := r.String()
	r.Stats()
	return str
}

func RenderToBytes(component *Element, data ...map[string]interface{}) []byte {
	return render(component, data...).Bytes()
}

func RenderToBytesWithStats(component *Element, data ...map[string]interface{}) []byte {
	r := render(component, data...)
	bytes := r.Bytes()
	r.Stats()
	return bytes
}
