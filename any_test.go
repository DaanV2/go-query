package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Any(t *testing.T) {
	t.Run("If any return true, any should return true", func(t *testing.T) {
		arr := RandomDataArray(10)

		var (
			triggerd1 = false
			triggerd2 = false
		)

		amount, err := OverArray(arr).Any(
			func(item *TestData, key int) bool {
				triggerd1 = true
				return true
			}, func(item *TestData, key int) bool {
				triggerd2 = true
				return true
			}).Count()

		require.Nil(t, err)
		assert.True(t, triggerd1)
		assert.False(t, triggerd2)
		assert.Len(t, arr, amount)
	})

	t.Run("If any return false, no items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		amount, err := OverArray(arr).Any(
			func(item *TestData, key int) bool {
				return false
			}, func(item *TestData, key int) bool {
				return true
			}).Count()

		require.NoError(t, err)
		assert.Len(t, arr, amount)
	})

	t.Run("If all return false, no items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		amount, err := OverArray(arr).Any(
			func(item *TestData, key int) bool {
				return false
			}, func(item *TestData, key int) bool {
				return false
			}).Count()

		require.NoError(t, err)
		assert.Equal(t, 0, amount)
	})
}

func Test_HasAny(t *testing.T) {
	t.Run("If any return true, result should be true", func(t *testing.T) {
		arr := RandomDataArray(10)

		var (
			triggerd1 = false
			triggerd2 = false
		)

		result, err := HasAny(arr,
			func(item *TestData, key int) bool {
				triggerd1 = true
				return true
			}, func(item *TestData, key int) bool {
				triggerd2 = true
				return true
			})

		require.NoError(t, err)
		assert.True(t, triggerd1)
		assert.False(t, triggerd2)
		assert.True(t, result)
	})

	t.Run("If any return false, but some still return true, items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		result, err := HasAny(arr,
			func(item *TestData, key int) bool {
				return false
			}, func(item *TestData, key int) bool {
				return true
			})

		require.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("If all return false, no items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		result, err := HasAny(arr,
			func(item *TestData, key int) bool {
				return false
			}, func(item *TestData, key int) bool {
				return false
			})

		require.NoError(t, err)
		assert.True(t, result)
	})
}
