package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyProps struct {
	MyCustomProp interface{} `structs:",omitempty"`
	MySecondProp interface{} `structs:",omitempty"`
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
}
