package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Filter(t *testing.T) {
	t.Run("Array", func(t *testing.T) {
		arr := RandomDataArray(10)

		count := 0

		WalkArray(arr, func(item *TestData, key int) error {
			if item.Age > 20 {
				count++
			}
			return nil
		})

		w, err := Filter(arr, func(item *TestData, key int) bool {
			return item.Age > 20
		}).Count()

		assert.NoError(t, err)
		assert.Equal(t, count, w)
	})

	t.Run("Name Map", func(t *testing.T) {
		m := RandomNameDataMap(10)
		count := 0

		WalkMap(m, func(item *TestData, key string) error {
			if item.Age > 20 {
				count++
			}
			return nil
		})

		w, err := Filter(m, func(item *TestData, key string) bool {
			return item.Age > 20
		}).Count()

		assert.NoError(t, err)
		assert.Equal(t, count, w)
	})

	t.Run("Age Map", func(t *testing.T) {
		m := RandomAgeDataMap(10)
		count := 0

		WalkMap(m, func(item *TestData, key int) error {
			if item.Age > 20 {
				count++
			}
			return nil
		})

		w, err := Filter(m, func(item *TestData, key int) bool {
			return item.Age > 20
		}).Count()

		assert.NoError(t, err)
		assert.Equal(t, count, w)
	})
}
