package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComment(t *testing.T) {
	t.Run("it should correctly return a simple comment", func(t *testing.T) {
		comment, err := Render(Comment("simple test comment"))

		assert.NoError(t, err)
		assert.Equal(t, "<!-- simple test comment -->", comment)
	})

	t.Run("it should correctly return a formatted comment", func(t *testing.T) {
		comment, err := Render(Comment("test %d %d %d", 1, 2, 3))

		assert.NoError(t, err)
		assert.Equal(t, "<!-- test 1 2 3 -->", comment)
	})
}

func TestDoctype(t *testing.T) {
	t.Run("it should correctly return default doctype", func(t *testing.T) {
		comment, err := Render(Doctype())

		assert.NoError(t, err)
		assert.Equal(t, "<!doctype html>", comment)
	})

	t.Run("it should correctly return a custom doctype", func(t *testing.T) {
		comment, err := Render(Doctype("HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\""))

		assert.NoError(t, err)
		assert.Equal(t, "<!doctype HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\">", comment)
	})

	t.Run("it should correctly return a custom full doctype", func(t *testing.T) {
		comment, err := Render(Doctype("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\">"))

		assert.NoError(t, err)
		assert.Equal(t, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\">", comment)
	})

	t.Run("it should correctly return right doctype on nil", func(t *testing.T) {
		comment, err := Render(Doctype(nil))

		assert.NoError(t, err)
		assert.Equal(t, "<!doctype html>", comment)
	})

	t.Run("it should correctly return right doctype on empty string", func(t *testing.T) {
		comment, err := Render(Doctype(""))

		assert.NoError(t, err)
		assert.Equal(t, "<!doctype html>", comment)
	})
}
