package react

import (
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

func TestComment(t *testing.T) {
	t.Run("it should correctly return a simple comment", func(t *testing.T) {
		comment := Render(Comment("simple test comment"))

		assert.Equal(t, "<!-- simple test comment -->", comment)
	})

	t.Run("it should correctly return a formatted comment", func(t *testing.T) {
		comment := Render(Comment("test %d %d %d", 1, 2, 3))

		assert.Equal(t, "<!-- test 1 2 3 -->", comment)
	})
}

func TestDoctype(t *testing.T) {
	t.Run("it should correctly return default doctype", func(t *testing.T) {
		comment := Render(Doctype())

		assert.Equal(t, "<!doctype html>", comment)
	})

	t.Run("it should correctly return a custom doctype", func(t *testing.T) {
		comment := Render(Doctype("HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\""))

		assert.Equal(t, "<!doctype HTML PUBLIC \"-//W3C//DTD HTML 4.01//EN\" \"http://www.w3.org/TR/html4/strict.dtd\">", comment)
	})

	t.Run("it should correctly return a custom full doctype", func(t *testing.T) {
		comment := Render(Doctype("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\">"))

		assert.Equal(t, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\">", comment)
	})

	t.Run("it should correctly return right doctype on nil", func(t *testing.T) {
		comment := Render(Doctype(nil))

		assert.Equal(t, "<!doctype html>", comment)
	})

	t.Run("it should correctly return right doctype on empty string", func(t *testing.T) {
		comment := Render(Doctype(""))

		assert.Equal(t, "<!doctype html>", comment)
	})
}

func TestMeta(t *testing.T) {
	t.Run("it should correctly return charsetMeta", func(t *testing.T) {
		t.Run("empty call", func(t *testing.T) {
			meta := Render(MetaCharset())

			assert.Equal(t, "<meta charset=\"UTF-8\" />", meta)
		})

		t.Run("custom charset", func(t *testing.T) {
			meta := Render(MetaCharset("ISO-8859-1"))

			assert.Equal(t, "<meta charset=\"ISO-8859-1\" />", meta)
		})
	})

	t.Run("it should correctly return meta tag", func(t *testing.T) {
		t.Run("null call", func(t *testing.T) {
			meta := Render(Meta("", "", nil))
			assert.Equal(t, "<meta />", meta)
		})

		t.Run("name call", func(t *testing.T) {
			meta := Render(Meta("propName", "", nil))
			assert.Equal(t, "<meta name=\"propName\" />", meta)
		})

		t.Run("content call", func(t *testing.T) {
			meta := Render(Meta("", "content value", nil))
			assert.Equal(t, "<meta content=\"content value\" />", meta)
		})

		t.Run("name and content call", func(t *testing.T) {
			meta := Render(Meta("propName", "content value", nil))
			assert.Contains(t, meta, "content=\"content value\"")
			assert.Contains(t, meta, "name=\"propName\"")
		})
	})

	t.Run("it should correctly return meta itemprop tag", func(t *testing.T) {
		t.Run("null call", func(t *testing.T) {
			meta := Render(MetaItemProp("", "", nil))
			assert.Equal(t, "<meta />", meta)
		})

		t.Run("itemprop call", func(t *testing.T) {
			meta := Render(MetaItemProp("propName", "", nil))
			assert.Equal(t, "<meta itemprop=\"propName\" />", meta)
		})

		t.Run("content call", func(t *testing.T) {
			meta := Render(MetaItemProp("", "content value", nil))
			assert.Equal(t, "<meta content=\"content value\" />", meta)
		})

		t.Run("itemprop and content call", func(t *testing.T) {
			meta := Render(MetaItemProp("propName", "content value", nil))
			assert.Contains(t, meta, "content=\"content value\"")
			assert.Contains(t, meta, "itemprop=\"propName\"")
		})

		t.Run("it should correctly return meta property tag", func(t *testing.T) {
			t.Run("null call", func(t *testing.T) {
				meta := Render(MetaProperty("", "", nil))
				assert.Equal(t, "<meta />", meta)
			})

			t.Run("property call", func(t *testing.T) {
				meta := Render(MetaProperty("propName", "", nil))
				assert.Equal(t, "<meta property=\"propName\" />", meta)
			})

			t.Run("content call", func(t *testing.T) {
				meta := Render(MetaProperty("", "content value", nil))
				assert.Equal(t, "<meta content=\"content value\" />", meta)
			})

			t.Run("property and content call", func(t *testing.T) {
				meta := Render(MetaProperty("propName", "content value", nil))
				assert.Contains(t, meta, "content=\"content value\"")
				assert.Contains(t, meta, "property=\"propName\"")
			})
		})
	})

	t.Run("it should correctly return http-equiv meta", func(t *testing.T) {
		meta := Render(Meta("", "ie=edge", &Props{HttpEquiv: "X-UA-Compatible"}))
		assert.Equal(t, "<meta content=\"ie=edge\" http-equiv=\"X-UA-Compatible\" />", meta)
	})
}

func TestJavascript(t *testing.T) {
	t.Run("it should correctly render script tag", func(t *testing.T) {
		t.Run("on empty", func(t *testing.T) {
			content := Render(Script("", nil))

			assert.Equal(t, "<script type=\"application/javascript\"></script>", content)
		})

		t.Run("on src string", func(t *testing.T) {
			content := Render(Script("/local/file.js", nil))

			assert.Equal(t, "<script src=\"/local/file.js\" type=\"application/javascript\"></script>", content)
		})

		t.Run("on src URL", func(t *testing.T) {
			content := Render(Script(template.URL("/local/file.js"), nil))

			assert.Equal(t, "<script src=\"/local/file.js\" type=\"application/javascript\"></script>", content)
		})

		t.Run("on src JS", func(t *testing.T) {
			content := Render(Script(template.JS(`function a () { return "hello" } alert(a());`), nil))

			assert.Equal(t, "<script type=\"application/javascript\">function a () { return \"hello\" } alert(a());</script>", content)
		})
	})
}
