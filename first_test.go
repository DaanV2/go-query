package query

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_First(t *testing.T) {
	t.Run("Arrays", func(t *testing.T) {
		arr := RandomDataArray(10)

		key, item, err := OverArray(arr).First()
		assert.NoError(t, err)
		assert.Equal(t, arr[0], item)
		assert.Equal(t, 0, key)

		key, item, err = Over[int, *TestData](arr).First()
		assert.NoError(t, err)
		assert.Equal(t, arr[0], item)
		assert.Equal(t, 0, key)

		key, item, err = First[int, *TestData](arr, nil)
		assert.NoError(t, err)
		assert.Equal(t, arr[0], item)
		assert.Equal(t, 0, key)
	})

	t.Run("Name Maps", func(t *testing.T) {
		m := RandomNameDataMap(10)

		key, item, err := OverMap(m).First()
		assert.NoError(t, err)
		assert.Equal(t, m[key], item)

		key, item, err = Over[string, *TestData](m).First()
		assert.NoError(t, err)
		assert.Equal(t, m[key], item)
	})

	t.Run("Age Maps", func(t *testing.T) {
		m := RandomAgeDataMap(10)

		key, item, err := OverMap(m).First()
		assert.NoError(t, err)
		assert.Equal(t, m[key], item)

		key, item, err = Over[int, *TestData](m).First()
		assert.NoError(t, err)
		assert.Equal(t, m[key], item)
	})

	t.Run("When no matches, returns ErrNoMatches", func(t *testing.T) {
		arr := RandomDataArray(10)

		_, _, err := OverArray(arr).Filter(func(item *TestData, key int) bool {
			return false
		}).First()
		assert.Equal(t, ErrNoMatches, err)
	})

	t.Run("When error walking, returns error", func(t *testing.T) {
		arr := RandomDataArray(10)

		_, _, err := OverArray(arr).Filter(func(item *TestData, key int) bool {
			panic(errors.New("test"))
		}).First()
		assert.Error(t, err)
	})
}
