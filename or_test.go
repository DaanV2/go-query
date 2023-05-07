package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Or(t *testing.T) {

	t.Run("If any return true, any should return true", func(t *testing.T) {
		arr := RandomDataArray(10)

		var (
			triggerd1 = false
			triggerd2 = false
		)

		amount, err := OverArray(arr).
			Or(func(item *TestData, key int) bool {
				triggerd1 = true
				return true
			}).
			Or(func(item *TestData, key int) bool {
				triggerd2 = true
				return true
			}).Count()

		require.Nil(t, err)
		assert.True(t, triggerd1)
		assert.False(t, triggerd2)
		assert.Len(t, arr, amount)
	})

	t.Run("If any return false, but other true, items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		amount, err := OverArray(arr).
			Or(func(item *TestData, key int) bool {
				return false
			}).
			Or(func(item *TestData, key int) bool {
				return true
			}).Count()

		require.NoError(t, err)
		assert.Len(t, arr, amount)
	})

	t.Run("If one is nil, items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		amount, err := OverArray(arr).
			Or(func(item *TestData, key int) bool {
				return false
			}).
			Or(nil).Count()

		require.NoError(t, err)
		assert.Len(t, arr, amount)
	})

	t.Run("If all are false, no items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		amount, err := OverArray(arr).
			Or(func(item *TestData, key int) bool {
				return false
			}).
			Or(func(item *TestData, key int) bool {
				return false
			}).Count()

		require.NoError(t, err)
		assert.Equal(t, 0, amount)
	})
}
