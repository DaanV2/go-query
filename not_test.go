package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Not(t *testing.T) {

	t.Run("If matchers return false, then Not should return true", func(t *testing.T) {
		arr := RandomDataArray(10)

		amount, err := OverArray(arr).
			Filter(func(item *TestData, key int) bool {
				return false
			}).
			Not().Count()

		require.NoError(t, err)
		assert.Len(t, arr, amount)
	})
}
