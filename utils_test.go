package react

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type StructA struct {
	B string
	C string
	D string
}

func TestToMap(t *testing.T) {
	t.Run("it should convert struct to map correctly", func(t *testing.T) {
		t.Run("empty map", func(t *testing.T) {
			m := toMap(&StructA{})

			mb, b := m["B"]
			mc, c := m["C"]
			md, d := m["D"]

			assert.Equal(t, mb, nil)
			assert.Equal(t, mc, nil)
			assert.Equal(t, md, nil)

			assert.False(t, b)
			assert.False(t, c)
			assert.False(t, d)
		})

		t.Run("non empty message", func(t *testing.T) {
			m := toMap(&StructA{B: "b", C: "c"})

			mb, b := m["B"]
			mc, c := m["C"]
			md, d := m["D"]

			assert.Equal(t, mb, "b")
			assert.Equal(t, mc, "c")
			assert.Equal(t, md, nil)

			assert.True(t, b)
			assert.True(t, c)
			assert.False(t, d)
		})

		t.Run("non empty but defined after", func(t *testing.T) {
			s := &StructA{B: "b"}
			s.C = "c"

			m := toMap(s)

			mb, b := m["B"]
			mc, c := m["C"]
			md, d := m["D"]

			assert.Equal(t, mb, "b")
			assert.Equal(t, mc, "c")
			assert.Equal(t, md, nil)

			assert.True(t, b)
			assert.True(t, c)
			assert.False(t, d)
		})

		t.Run("empty but defined after", func(t *testing.T) {
			s := &StructA{}
			s.B = "b"
			s.C = "c"
			s.D = "d"

			m := toMap(s)

			mb, b := m["B"]
			mc, c := m["C"]
			md, d := m["D"]

			assert.Equal(t, mb, "b")
			assert.Equal(t, mc, "c")
			assert.Equal(t, md, "d")

			assert.True(t, b)
			assert.True(t, c)
			assert.True(t, d)
		})

		t.Run("empty but full replaced", func(t *testing.T) {
			s := &StructA{}
			s2 := &StructA{
				B: "b",
				C: "c",
				D: "d",
			}

			s = s2

			m := toMap(s)

			mb, b := m["B"]
			mc, c := m["C"]
			md, d := m["D"]

			assert.Equal(t, mb, "b")
			assert.Equal(t, mc, "c")
			assert.Equal(t, md, "d")

			assert.True(t, b)
			assert.True(t, c)
			assert.True(t, d)
		})

		t.Run("convert map to map", func(t *testing.T) {
			m := toMap(map[string]interface{}{
				"B": "b",
				"C": "c",
				"D": "d",
			})

			mb, b := m["B"]
			mc, c := m["C"]
			md, d := m["D"]

			assert.Equal(t, mb, "b")
			assert.Equal(t, mc, "c")
			assert.Equal(t, md, "d")

			assert.True(t, b)
			assert.True(t, c)
			assert.True(t, d)
		})

		t.Run("convert slice of map to map", func(t *testing.T) {
			m := toMap([]map[string]interface{}{
				{"B": "b"},
				{"C": "c"},
				{"D": "d"},
			})

			mb, b := m["B"]
			mc, c := m["C"]
			md, d := m["D"]

			assert.Equal(t, mb, "b")
			assert.Equal(t, mc, "c")
			assert.Equal(t, md, "d")

			assert.True(t, b)
			assert.True(t, c)
			assert.True(t, d)
		})

		t.Run("convert slice of map and struct to map", func(t *testing.T) {
			m := toMap([]interface{}{
				map[string]interface{}{"B": "b"},
				&StructA{C: "c"},
				StructA{D: "d"},
			})

			mb, b := m["B"]
			mc, c := m["C"]
			md, d := m["D"]

			assert.Equal(t, mb, "b")
			assert.Equal(t, mc, "c")
			assert.Equal(t, md, "d")

			assert.True(t, b)
			assert.True(t, c)
			assert.True(t, d)
		})
	})
}
