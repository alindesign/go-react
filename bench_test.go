package react

import (
	"strconv"
	"strings"
	"testing"
)

var (
	div *Element
	app *Element
	ela *Element
)

func init() {
	div = CreateElement("div", Props{"className": "ui__container"})
	app = App()
	ela = CreateElement("div", nil, List(4),
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
	)
}

func BenchmarkRender(b *testing.B) {
	b.Run("SimpleElement", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Render(div)
		}
	})
	b.Run("SmallElement", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Render(app)
		}
	})
	b.Run("TreeElement", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = NewRenderer(ela).String()
		}
	})
}

func BenchmarkJoins(b *testing.B) {
	var r []string

	for i := 0; i < 1024; i++ {
		r = append(r, "item "+strconv.Itoa(i))
	}

	b.Run("strings.Join", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = strings.Join(r, "")
		}
	})

	b.Run("string loop concat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			str := ""
			for _, e := range r {
				str += e
			}
		}
	})

	b.Run("string builder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var str strings.Builder
			for _, e := range r {
				str.WriteString(e)
			}
		}
	})
}

func BenchmarkStrings(b *testing.B) {
	var rs []string
	var rb [][]byte

	for i := 0; i < 1024; i++ {
		rs = append(rs, "item "+strconv.Itoa(i))
		rb = append(rb, []byte("item "+strconv.Itoa(i)))
	}

	b.Run("strings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var content strings.Builder
			for j := 0; j < 1024; j++ {
				content.WriteString(rs[j])
			}
		}
	})

	b.Run("bytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var content strings.Builder
			for j := 0; j < 1024; j++ {
				content.Write(rb[j])
			}
		}
	})
}
