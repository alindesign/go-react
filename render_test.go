package react

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func List(num int) *Element {
	var items []*Element
	for i := 0; i < num; i++ {
		items = append(items, CreateElement("li", Props{
			"className": "list__item item-" + strconv.Itoa(i+1),
		}, Text("Item %d", i+1)))
	}

	return CreateElement("ul", Props{
		"className": "list__root",
	}, items...)
}

func App() *Element {
	return CreateElement("div", Props{
		"className": "root",
	}, List(4))
}

func TestRender(t *testing.T) {
	t.Run("it should render an empty div", func(t *testing.T) {
		div := Render(CreateElement("div", nil))

		assert.Equal(t, `<div></div>`, div)
	})

	t.Run("it should render to bytes an empty div", func(t *testing.T) {
		div := RenderBytes(CreateElement("div", nil))

		assert.Equal(t, []byte("<div></div>"), div)
	})

	t.Run("it should render a simple div", func(t *testing.T) {
		div := Render(CreateElement("div", Props{
			"className": "container",
		}))

		assert.Equal(t, `<div class="container"></div>`, div)
	})

	t.Run("it should render a small Component", func(t *testing.T) {
		div := Render(App())

		assert.Equal(t, `<div class="root"><ul class="list__root"><li class="list__item item-1">Item 1</li><li class="list__item item-2">Item 2</li><li class="list__item item-3">Item 3</li><li class="list__item item-4">Item 4</li></ul></div>`, div)
	})

	t.Run("it should render attrs correctly", func(t *testing.T) {
		div := Render(CreateElement("div", Props{
			"data": `{"json": "true"}`,
		}))

		assert.Equal(t, `<div data="{&#34;json&#34;: &#34;true&#34;}"></div>`, div)
	})

	t.Run("it should render children correctly", func(t *testing.T) {
		div := Render(CreateElement("div", Props{
			"data": `{"json": "true"}`,
		}, Text("Text <script>function () { alert('broken'); }</script>")))

		assert.Equal(t, `<div data="{&#34;json&#34;: &#34;true&#34;}">Text &lt;script&gt;function () { alert(&#39;broken&#39;); }&lt;/script&gt;</div>`, div)
	})

	t.Run("it should correctly run a void element", func(t *testing.T) {
		div := Render(CreateElement("img", Props{
			"src": `url`,
		}, Text("Children elements text")))

		assert.Equal(t, `<img src="url"/>`, div)
	})

	t.Run("it should render fragment correctly", func(t *testing.T) {
		content := Render(CreateFragment(
			CreateElement("span", nil, Text("Item1")),
			CreateElement("span", nil, Text("Item2")),
		))

		assert.Equal(t, "<span>Item1</span><span>Item2</span>", content)
	})
}

func TestRenderElement(t *testing.T) {
	t.Run("it should render pure html correctly", func(t *testing.T) {
		value := `<h1>Hello World!</h1>`
		str := Render(HTMLString(value))

		assert.Equal(t, value, str)
	})

	t.Run("it should render pure css correctly", func(t *testing.T) {
		value := `body {
	color: #fff;
	background: url(/path/to/image.png);
	font-size: 14px;
	font-family: 'Roboto', sans-serif;
}`
		str := Render(CSSString(value))

		assert.Equal(t, value, str)
	})

	t.Run("it should render pure js correctly", func(t *testing.T) {
		value := `function a (b, c) { return b + c; }`
		str := Render(JSString(value))

		assert.Equal(t, value, str)
	})

	t.Run("it should render pure string correctly", func(t *testing.T) {
		value := `Sample text`
		str := Render(Text(value))

		assert.Equal(t, value, str)
	})

	t.Run("it should render combined html + string correctly", func(t *testing.T) {
		value := `Sample text <script type="text/javascript">alert("Hi!");</script>`
		str := Render(Text(value))

		assert.Equal(t, `Sample text &lt;script type=&#34;text/javascript&#34;&gt;alert(&#34;Hi!&#34;);&lt;/script&gt;`, str)
	})
}
