package react

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"html/template"
	"strconv"
	"testing"
)

func List(num int) *Element {
	var items []interface{}
	for i := 0; i < num; i++ {
		items = append(items, CreateElement("li", &Props{
			ClassName: "list__item item-" + strconv.Itoa(i+1),
		}, fmt.Sprintf("Item %d", i+1)))
	}

	return CreateElement("ul", &Props{
		ClassName: "list__root",
	}, items...)
}

func App() *Element {
	return CreateElement("div", &Props{
		ClassName: "root",
	}, List(4))
}

func TestRender(t *testing.T) {
	t.Run("it should render an empty div", func(t *testing.T) {
		div, err := Render(CreateElement("div", nil))

		assert.NoError(t, err)
		assert.Equal(t, `<div></div>`, div)
	})

	t.Run("it should render a simple div", func(t *testing.T) {
		div, err := Render(CreateElement("div", &Props{
			ClassName: "ui__container",
		}))

		assert.NoError(t, err)
		assert.Equal(t, `<div class="ui__container"></div>`, div)
	})

	t.Run("it should render a small component", func(t *testing.T) {
		div, err := Render(App())

		assert.NoError(t, err)
		assert.Equal(t, `<div class="root"><ul class="list__root"><li class="list__item item-1">Item 1</li><li class="list__item item-2">Item 2</li><li class="list__item item-3">Item 3</li><li class="list__item item-4">Item 4</li></ul></div>`, div)
	})

	t.Run("it should render attrs correctly", func(t *testing.T) {
		div, err := Render(CreateElement("div", &Props{
			Data: `{"json": "true"}`,
		}))

		assert.NoError(t, err)
		assert.Equal(t, `<div data="{&#34;json&#34;: &#34;true&#34;}"></div>`, div)
	})

	t.Run("it should render children correctly", func(t *testing.T) {
		div, err := Render(CreateElement("div", &Props{
			Data: `{"json": "true"}`,
		}, "Text <script>function () { alert('broken'); }</script>"))

		assert.NoError(t, err)
		assert.Equal(t, `<div data="{&#34;json&#34;: &#34;true&#34;}">Text &lt;script&gt;function () { alert(&#39;broken&#39;); }&lt;/script&gt;</div>`, div)
	})

	t.Run("it should correctly run a void element", func(t *testing.T) {
		div, err := Render(CreateElement("img", &Props{
			Src: `url`,
		}, "Children elements text"))

		assert.NoError(t, err)
		assert.Equal(t, `<img src="url" />`, div)
	})

	t.Run("it should test on non valid element", func(t *testing.T) {
		content, err := Render(CreateElement("span", nil, 1024))

		assert.Error(t, err)
		assert.Equal(t, "", content)
	})

	t.Run("it should render fragment correctly", func(t *testing.T) {
		content, err := Render(CreateFragment(
			CreateElement("span", nil, "Item1"),
			CreateElement("span", nil, "Item2"),
		))

		assert.NoError(t, err)
		assert.Equal(t, "<span>Item1</span><span>Item2</span>", content)
	})
}

func TestRenderElement(t *testing.T) {
	t.Run("it should render pure html correctly", func(t *testing.T) {
		value := `<h1>Hello World!</h1>`
		str, err := renderElement(template.HTML(value), nil)

		assert.NoError(t, err)
		assert.Equal(t, value, str)
	})

	t.Run("it should render pure css correctly", func(t *testing.T) {
		value := `body {
	color: #fff;
	background: url(/path/to/image.png);
	font-size: 14px;
	font-family: 'Roboto', sans-serif;
}`
		str, err := renderElement(template.CSS(value), nil)

		assert.NoError(t, err)
		assert.Equal(t, value, str)
	})

	t.Run("it should render pure js correctly", func(t *testing.T) {
		value := `function a (b, c) { return b + c; }`
		str, err := renderElement(template.JS(value), nil)

		assert.NoError(t, err)
		assert.Equal(t, value, str)
	})

	t.Run("it should render pure string correctly", func(t *testing.T) {
		value := `Sample text`
		str, err := renderElement(value, nil)

		assert.NoError(t, err)
		assert.Equal(t, value, str)
	})

	t.Run("it should render combined html + string correctly", func(t *testing.T) {
		value := `Sample text <script type="text/javascript">alert("Hi!");</script>`
		str, err := renderElement(value, nil)

		assert.NoError(t, err)
		assert.Equal(t, `Sample text &lt;script type=&#34;text/javascript&#34;&gt;alert(&#34;Hi!&#34;);&lt;/script&gt;`, str)
	})
}