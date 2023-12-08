package react

func render(component *Element, data ...map[string]any) *Renderer {
	return NewRenderer(component).SetData(data...)
}

func Render(component *Element, data ...map[string]any) string {
	return render(component, data...).String()
}

func RenderWithStats(component *Element, data ...map[string]any) string {
	r := render(component, data...)
	str := r.String()
	r.Stats()
	return str
}

func RenderToBytes(component *Element, data ...map[string]any) []byte {
	return render(component, data...).Bytes()
}

func RenderToBytesWithStats(component *Element, data ...map[string]any) []byte {
	r := render(component, data...)
	bytes := r.Bytes()
	r.Stats()
	return bytes
}
