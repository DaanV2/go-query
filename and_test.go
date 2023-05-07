package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_And(t *testing.T) {
	t.Run("If both return true, and should return true", func(t *testing.T) {
		arr := RandomDataArray(10)

		var (
			triggerd1 = false
			triggerd2 = false
		)

		amount, err := OverArray(arr).
			And(func(item *TestData, key int) bool {
				triggerd1 = true
				return true
			}).
			And(func(item *TestData, key int) bool {
				triggerd2 = true
				return true
			}).Count()

		require.NoError(t, err)
		assert.True(t, triggerd1)
		assert.True(t, triggerd2)
		assert.Len(t, arr, amount)
	})

	t.Run("If one return false, and result should be false", func(t *testing.T) {
		arr := RandomDataArray(10)

		var (
			triggerd1 = false
			triggerd2 = false
		)

		amount, err := OverArray(arr).
			And(func(item *TestData, key int) bool {
				triggerd1 = true
				return false
			}).
			And(func(item *TestData, key int) bool {
				triggerd2 = true
				return true
			}).Count()

		require.NoError(t, err)
		assert.True(t, triggerd1)
		assert.False(t, triggerd2)
		assert.Equal(t, 0, amount)
	})

	t.Run("If one is nil, and should return false", func(t *testing.T) {
		arr := RandomDataArray(10)

		var (
			triggerd1 = false
		)

		amount, err := OverArray(arr).
			And(func(item *TestData, key int) bool {
				triggerd1 = true
				return false
			}).
			And(nil).Count()

		require.NoError(t, err)
		assert.True(t, triggerd1)
		assert.Equal(t, 0, amount)
	})
}
