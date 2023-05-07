package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_All(t *testing.T) {
	t.Run("If all return true, all should have been called", func(t *testing.T) {
		arr := RandomDataArray(10)

		var (
			triggerd1 = false
			triggerd2 = false
		)

		amount, err := OverArray(arr).All(
			func(item *TestData, key int) bool {
				triggerd1 = true
				return true
			}, func(item *TestData, key int) bool {
				triggerd2 = true
				return true
			}).Count()

		require.NoError(t, err)
		assert.True(t, triggerd1)
		assert.True(t, triggerd2)
		assert.Len(t, arr, amount)
	})

	t.Run("If any return false, no items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		amount, err := OverArray(arr).All(
			func(item *TestData, key int) bool {
				return false
			}, func(item *TestData, key int) bool {
				return true
			}).Count()

		require.NoError(t, err)
		assert.Equal(t, 0, amount)
	})
}

func Test_HasAll(t *testing.T) {
	t.Run("If all return true, all should have been called", func(t *testing.T) {
		arr := RandomDataArray(10)

		var (
			triggerd1 = false
			triggerd2 = false
		)

		result, err := HasAll(arr,
			func(item *TestData, key int) bool {
				triggerd1 = true
				return true
			}, func(item *TestData, key int) bool {
				triggerd2 = true
				return true
			})

		require.NoError(t, err)
		assert.True(t, triggerd1)
		assert.True(t, triggerd2)
		assert.True(t, result)
	})

	t.Run("If any return false, no items should be returned", func(t *testing.T) {
		arr := RandomDataArray(10)

		result, err := HasAll(arr,
			func(item *TestData, key int) bool {
				return false
			}, func(item *TestData, key int) bool {
				return true
			})

		require.NoError(t, err)
		assert.True(t, result)
	})
}
