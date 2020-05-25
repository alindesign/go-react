package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyProps struct {
	MyCustomProp interface{}
	MySecondProp interface{}
}

type customThings struct {
	DataHello interface{}
}

func TestProps(t *testing.T) {
	t.Run("it should convert to empty", func(t *testing.T) {
		p := &Props{}
		assert.Equal(t, "", p.String())
	})

	t.Run("it should convert to string only added props", func(t *testing.T) {
		p := &Props{Width: 10, Height: 10}
		assert.Contains(t, p.String(), `height="10"`)
		assert.Contains(t, p.String(), `width="10"`)
	})

	t.Run("it should convert to string custom props structs correctly", func(t *testing.T) {
		mp := &MyProps{MyCustomProp: "test"}
		p := &Props{Width: 10, CustomProps: mp}

		assert.NotContains(t, p.String(), `my-second-prop`)
		assert.Contains(t, p.String(), `my-custom-prop="test"`)
		assert.Contains(t, p.String(), `width="10"`)
	})

	t.Run("it should convert class list correctly", func(t *testing.T) {
		p := &Props{ClassName: []string{"class-1", "class-2", "class-3"}}
		assert.Equal(t, `class="class-1 class-2 class-3"`, p.String())
	})

	t.Run("it should be able to add multiple custom props", func(t *testing.T) {
		mp := &MyProps{MyCustomProp: "test"}
		p := &Props{Width: 10, CustomProps: []interface{}{
			mp,
			&customThings{DataHello: "hi"},
			map[string]interface{}{"data-test": "test"},
		}}
		assert.Equal(t, p.String(), `data-hello="hi" data-test="test" my-custom-prop="test" width="10"`)
	})
}
