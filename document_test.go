package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocument(t *testing.T) {
	t.Run("it should render Document correctly", func(t *testing.T) {
		doc := Render(Document(nil, nil))
		assert.Equal(t, `<!doctype html><html lang="en"><head></head><body></body></html>`, doc)
	})

	t.Run("it should render Document correctly", func(t *testing.T) {
		doc := Render(Document(&DocumentProps{Body: H1(nil, Text("test"))}, nil))

		assert.Equal(t, `<!doctype html><html lang="en"><head></head><body><h1>test</h1></body></html>`, doc)
	})
}
