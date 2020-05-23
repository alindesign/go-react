package react

import (
	"testing"
)

func BenchmarkSimpleElement(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Render(CreateElement("div", &Props{
			ClassName: "ui__container",
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
