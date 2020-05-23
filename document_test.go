package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocument(t *testing.T) {
	t.Run("it should render Document correctly", func(t *testing.T) {
		doc, err := Render(Document(nil, nil))

		assert.NoError(t, err)
		assert.Equal(t, `<!doctype html><html lang="en"><head></head><body></body></html>`, doc)
	})

	t.Run("it should render Document correctly", func(t *testing.T) {
		doc, err := Render(Document(&DocumentProps{Body: H1(nil, "test")}, nil))

		assert.NoError(t, err)
		assert.Equal(t, `<!doctype html><html lang="en"><head></head><body><h1>test</h1></body></html>`, doc)
	})
}
