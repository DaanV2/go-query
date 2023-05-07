package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Arrays(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		count := 0

		arr := []int{
			1, 2, 3,
		}

		OverArray(arr).Walk(func(item int, key int) error {
			count++
			return nil
		})

		Over[int, int](arr).Walk(func(item int, key int) error {
			count++
			return nil
		})

		assert.Equal(t, 6, count)
	})
}

func Test_Maps(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		count := 0

		m := map[int]int{
			1: 2, 2: 3, 3: 4,
		}

		OverMap(m).Walk(func(item int, key int) error {
			count++
			return nil
		})

		Over[int, int](m).Walk(func(item int, key int) error {
			count++
			return nil
		})

		assert.Equal(t, 6, count)
	})
}

func Test_Channel(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		count := 0

		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)

		OverChannel(ch).Walk(func(item int, key int) error {
			count++
			return nil
		})

		ch = make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)

		Over[int, int](ch).Walk(func(item int, key int) error {
			count++
			return nil
		})

		assert.Equal(t, 6, count)
	})
}

func Test_Custom(t *testing.T) {
	collection := &TestCollection{
		Items: RandomDataArray(10),
	}

	t.Run("Correct generic types will be able to walk over the collection", func(t *testing.T) {
		count := 0
		_, err := Over[int, *TestData](collection).Walk(func(item *TestData, key int) error {
			count++
			return nil
		})

		assert.Equal(t, 10, count)
		assert.NoError(t, err)
	})
}

func Test_Over_Errors(t *testing.T) {
	t.Run("When not a collection, an error is returned", func(t *testing.T) {
		data := &TestData{
			Name:  "Test",
			Age:   10,
			Flag1: false,
			Flag2: true,
		}

		_, err := Over[int, int](data).Count()
		assert.Error(t, err)
		assert.Equal(t, ErrWrongWalkableType, err)
	})

	t.Run("When key type is not int for Arrays, an error is returned", func(t *testing.T) {
		data := []int{
			1, 2, 3,
		}

		_, err := Over[string, int](data).Count()
		assert.Error(t, err)
	})

	t.Run("When key type is not int for Channels, an error is returned", func(t *testing.T) {
		data := make(chan int, 3)
		data <- 1
		data <- 2
		data <- 3
		close(data)

		_, err := Over[string, int](data).Count()
		assert.Error(t, err)
	})
}
