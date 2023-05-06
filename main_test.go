package query

import "testing"

func Test_Arrays(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		count := 0

		arr := []int{
			1, 2, 3,
		}

		OverArray(arr).Walk(func(item int, key int) error {
			t.Logf("key: %v, item: %v", key, item)
			count++
			return nil
		})

		Over[int, int](arr).Walk(func(item int, key int) error {
			t.Logf("key: %v, item: %v", key, item)
			count++
			return nil
		})

		if count != 6 {
			t.Errorf("expected 6, got %v", count)
		}
	})
}

func Test_Maps(t *testing.T) {
	t.Run("Ints", func(t *testing.T) {
		count := 0

		m := map[int]int{
			1: 2, 2: 3, 3: 4,
		}

		OverMap(m).Walk(func(item int, key int) error {
			t.Logf("key: %v, item: %v", key, item)
			count++
			return nil
		})

		Over[int, int](m).Walk(func(item int, key int) error {
			t.Logf("key: %v, item: %v", key, item)
			count++
			return nil
		})

		if count != 6 {
			t.Errorf("expected 6, got %v", count)
		}
	})
}
