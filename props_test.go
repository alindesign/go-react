package react

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestProps(t *testing.T) {
	t.Run("it should convert to empty", func(t *testing.T) {
		p := Props{}
		assert.Equal(t, "", p.String())
	})

	t.Run("it should convert to string only added props", func(t *testing.T) {
		p := Props{"width": "10", "height": "10"}
		assert.Contains(t, p.String(), `height="10"`)
		assert.Contains(t, p.String(), `width="10"`)
	})

	t.Run("it should convert class list correctly", func(t *testing.T) {
		p := Props{"className": []string{"class-1", "class-2", "class-3"}}
		assert.Equal(t, `class="class-1 class-2 class-3"`, strings.TrimSpace(p.String()))
	})
}
