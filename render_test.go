package react

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func List(num int) *Element {
	var items []interface{}
	for i := 0; i < num; i++ {
		items = append(items, CreateElement("li", Props{
			"className": "list__item item-" + strconv.Itoa(i+1),
		}, fmt.Sprintf("Item %d", i+1)))
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
		div, err := Render(CreateElement("div", nil))

		assert.NoError(t, err)
		assert.Equal(t, `<div></div>`, div)
	})

	t.Run("it should render a simple div", func(t *testing.T) {
		div, err := Render(CreateElement("div", Props{
			"className": "ui__container",
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
		div, err := Render(CreateElement("div", Props{
			"data": `{"json": "true"}`,
		}))

		assert.NoError(t, err)
		assert.Equal(t, `<div data="{&#34;json&#34;: &#34;true&#34;}"></div>`, div)
	})

	t.Run("it should render children correctly", func(t *testing.T) {
		div, err := Render(CreateElement("div", Props{
			"data": `{"json": "true"}`,
		}, "Text <script>function () { alert('broken'); }</script>"))

		assert.NoError(t, err)
		assert.Equal(t, `<div data="{&#34;json&#34;: &#34;true&#34;}">Text &lt;script&gt;function () { alert(&#39;broken&#39;); }&lt;/script&gt;</div>`, div)
	})

	t.Run("it should correctly run a void element", func(t *testing.T) {
		div, err := Render(CreateElement("img", Props{
			"src": `url`,
		}, "Children elements text"))

		assert.NoError(t, err)
		assert.Equal(t, `<img src="url" />`, div)
	})
}

func BenchmarkSimpleElement(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Render(CreateElement("div", Props{
			"className": "ui__container",
		}))
	}
}

func BenchmarkSmallElement(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Render(App())
	}
}

func BenchmarkTreeElement(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Render(CreateElement("div", nil, List(4),
			CreateElement("div", nil, List(4),
				CreateElement("div", nil, List(4)),
				CreateElement("div", nil, List(4)),
				CreateElement("div", nil, List(4),
					CreateElement("div", nil, List(4),
						CreateElement("div", nil, List(4)),
					),
					CreateElement("div", nil, List(4)),
				),
			),
		))
	}
}
