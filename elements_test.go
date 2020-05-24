package react

import (
	"github.com/stretchr/testify/assert"
	"html/template"
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

func TestMeta(t *testing.T) {
	t.Run("it should correctly return charsetMeta", func(t *testing.T) {
		t.Run("empty call", func(t *testing.T) {
			meta, err := Render(MetaCharset())
			assert.NoError(t, err)
			assert.Equal(t, "<meta charset=\"UTF-8\" />", meta)
		})

		t.Run("custom charset", func(t *testing.T) {
			meta, err := Render(MetaCharset("ISO-8859-1"))
			assert.NoError(t, err)
			assert.Equal(t, "<meta charset=\"ISO-8859-1\" />", meta)
		})
	})

	t.Run("it should correctly return meta tag", func(t *testing.T) {
		t.Run("null call", func(t *testing.T) {
			meta, err := Render(Meta("", "", nil))
			assert.NoError(t, err)
			assert.Equal(t, "<meta />", meta)
		})

		t.Run("name call", func(t *testing.T) {
			meta, err := Render(Meta("propName", "", nil))
			assert.NoError(t, err)
			assert.Equal(t, "<meta name=\"propName\" />", meta)
		})

		t.Run("content call", func(t *testing.T) {
			meta, err := Render(Meta("", "content value", nil))
			assert.NoError(t, err)
			assert.Equal(t, "<meta content=\"content value\" />", meta)
		})

		t.Run("name and content call", func(t *testing.T) {
			meta, err := Render(Meta("propName", "content value", nil))
			assert.NoError(t, err)
			assert.Contains(t, meta, "content=\"content value\"")
			assert.Contains(t, meta, "name=\"propName\"")
		})
	})

	t.Run("it should correctly return meta itemprop tag", func(t *testing.T) {
		t.Run("null call", func(t *testing.T) {
			meta, err := Render(MetaItemProp("", "", nil))
			assert.NoError(t, err)
			assert.Equal(t, "<meta />", meta)
		})

		t.Run("itemprop call", func(t *testing.T) {
			meta, err := Render(MetaItemProp("propName", "", nil))
			assert.NoError(t, err)
			assert.Equal(t, "<meta itemprop=\"propName\" />", meta)
		})

		t.Run("content call", func(t *testing.T) {
			meta, err := Render(MetaItemProp("", "content value", nil))
			assert.NoError(t, err)
			assert.Equal(t, "<meta content=\"content value\" />", meta)
		})

		t.Run("itemprop and content call", func(t *testing.T) {
			meta, err := Render(MetaItemProp("propName", "content value", nil))
			assert.NoError(t, err)
			assert.Contains(t, meta, "content=\"content value\"")
			assert.Contains(t, meta, "itemprop=\"propName\"")
		})

		t.Run("it should correctly return meta property tag", func(t *testing.T) {
			t.Run("null call", func(t *testing.T) {
				meta, err := Render(MetaProperty("", "", nil))
				assert.NoError(t, err)
				assert.Equal(t, "<meta />", meta)
			})

			t.Run("property call", func(t *testing.T) {
				meta, err := Render(MetaProperty("propName", "", nil))
				assert.NoError(t, err)
				assert.Equal(t, "<meta property=\"propName\" />", meta)
			})

			t.Run("content call", func(t *testing.T) {
				meta, err := Render(MetaProperty("", "content value", nil))
				assert.NoError(t, err)
				assert.Equal(t, "<meta content=\"content value\" />", meta)
			})

			t.Run("property and content call", func(t *testing.T) {
				meta, err := Render(MetaProperty("propName", "content value", nil))
				assert.NoError(t, err)
				assert.Contains(t, meta, "content=\"content value\"")
				assert.Contains(t, meta, "property=\"propName\"")
			})
		})
	})
}

func TestJavascript(t *testing.T) {
	t.Run("it should correctly render script tag", func(t *testing.T) {
		t.Run("on empty", func(t *testing.T) {
			content, err := Render(Script("", nil))

			assert.NoError(t, err)
			assert.Equal(t, "<script type=\"application/javascript\"></script>", content)
		})

		t.Run("on src string", func(t *testing.T) {
			content, err := Render(Script("/local/file.js", nil))

			assert.NoError(t, err)
			assert.Equal(t, "<script src=\"/local/file.js\" type=\"application/javascript\"></script>", content)
		})

		t.Run("on src URL", func(t *testing.T) {
			content, err := Render(Script(template.URL("/local/file.js"), nil))

			assert.NoError(t, err)
			assert.Equal(t, "<script src=\"/local/file.js\" type=\"application/javascript\"></script>", content)
		})

		t.Run("on src JS", func(t *testing.T) {
			content, err := Render(Script(template.JS(`function a () { return "hello" } alert(a());`), nil))

			assert.NoError(t, err)
			assert.Equal(t, "<script type=\"application/javascript\">function a () { return \"hello\" } alert(a());</script>", content)
		})
	})
}
