package query

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Contain(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	t.Run("Contains should return true if there is a match", func(t *testing.T) {
		contains, err := Contains(arr, func(item int, key int) bool {
			return item == 3
		})

		require.NoError(t, err)
		require.True(t, contains)
	})

	t.Run("Contains should return false if there is no match", func(t *testing.T) {
		contains, err := Contains(arr, func(item int, key int) bool {
			return item == 10
		})

		require.NoError(t, err)
		require.False(t, contains)
	})

	t.Run("Multiple matches should return true", func(t *testing.T) {
		var (
			triggerd1 = false
			triggerd2 = false
		)

		contains, err := Contains(arr,
			Or(func(item int, key int) bool {
				triggerd1 = true
				return item == 10
			}, func(item int, key int) bool {
				triggerd2 = true
				return item == 3
			}))

		require.NoError(t, err)
		require.True(t, contains)
		require.True(t, triggerd1)
		require.True(t, triggerd2)
	})
}
