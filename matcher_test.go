package query

import "testing"

func Test_Matcher(t *testing.T) {
	t.Run("Does match works as intented", func(t *testing.T) {
		matcher := Match(func(item int, key int) bool {
			return item > 1
		})

		if !matcher.DoesMatch(2, 0) {
			t.Error("Matcher should match")
		}

		if matcher.DoesMatch(1, 0) {
			t.Error("Matcher should not match")
		}
	})

	t.Run("When nil, matches everything", func(t *testing.T) {
		matcher := Match[int, int](nil)

		if !matcher.DoesMatch(2, 0) {
			t.Error("Matcher should match")
		}

		if !matcher.DoesMatch(1, 0) {
			t.Error("Matcher should not match")
		}
	})

}
