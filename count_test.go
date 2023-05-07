package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Count(t *testing.T) {
	t.Run("Arrays", func(t *testing.T) {
		arr := RandomDataArray(10)

		count, err := OverArray(arr).Count()
		assert.NoError(t, err)
		assert.Equal(t, len(arr), count)

		count, err = Over[int, *TestData](arr).Count()
		assert.NoError(t, err)
		assert.Equal(t, len(arr), count)
	})

	t.Run("Name Maps", func(t *testing.T) {
		m := RandomNameDataMap(10)

		count, err := OverMap(m).Count()
		assert.NoError(t, err)
		assert.Equal(t, len(m), count)

		count, err = Over[string, *TestData](m).Count()
		assert.NoError(t, err)
		assert.Equal(t, len(m), count)
	})

	t.Run("Age Maps", func(t *testing.T) {
		m := RandomAgeDataMap(10)

		count, err := OverMap(m).Count()
		assert.NoError(t, err)
		assert.Equal(t, len(m), count)

		count, err = Over[int, *TestData](m).Count()
		assert.NoError(t, err)
		assert.Equal(t, len(m), count)
	})
}
