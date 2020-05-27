package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStyle(t *testing.T) {
	t.Run("it should convert to empty", func(t *testing.T) {
		s := &Style{}
		assert.Equal(t, "", s.String())
	})

	t.Run("it should convert to string only added styles", func(t *testing.T) {
		s := &Style{"width": "10", "height": "10px"}
		assert.Contains(t, s.String(), `height:10px`)
		assert.Contains(t, s.String(), `width:10`)
	})
}
