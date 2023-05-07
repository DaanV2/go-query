package query

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const RandomSeed = 0

type TestData struct {
	Age   int
	Name  string
	Flag1 bool
	Flag2 bool
}

var _ IWalkable[int, *TestData] = &TestCollection{}

type TestCollection struct {
	Items []*TestData
}

func (c *TestCollection) Walk(fn func(item *TestData, key int) error) (IWalkable[int, *TestData], error) {
	for i, item := range c.Items {
		if err := fn(item, i); err != nil {
			return c, err
		}
	}

	return c, nil
}


func Randomizer() *rand.Rand {
	return rand.New(rand.NewSource(RandomSeed))
}

func RandomData(r *rand.Rand) *TestData {
	age := r.Intn(100)

	nameL := r.Intn(10) + 3
	name := make([]byte, nameL)
	for i := 0; i < nameL; i++ {
		name[i] = byte(r.Intn(26) + 'a')
	}

	return &TestData{
		Age:   age,
		Name:  string(name),
		Flag1: r.Intn(2) == 0,
		Flag2: r.Intn(2) == 0,
	}
}

func RandomDataArray(n int) []*TestData {
	rand := Randomizer()
	data := make([]*TestData, 0, n)

	for i := 0; i < n; i++ {
		data = append(data, RandomData(rand))
	}

	return data
}

func RandomAgeDataMap(n int) map[int]*TestData {
	rand := Randomizer()
	data := make(map[int]*TestData, n)

	for i := 0; i < n; i++ {
		item := RandomData(rand)
		data[item.Age] = item
	}

	return data
}

func RandomNameDataMap(n int) map[string]*TestData {
	rand := Randomizer()
	data := make(map[string]*TestData, n)

	for i := 0; i < n; i++ {
		item := RandomData(rand)
		data[item.Name] = item
	}

	return data
}

func WalkArray[T any](arr []T, fn func(item T, key int) error) error {
	for key, item := range arr {
		if err := fn(item, key); err != nil {
			return err
		}
	}

	return nil
}

func WalkMap[K comparable, V any](m map[K]V, fn func(item V, key K) error) error {
	for key, item := range m {
		if err := fn(item, key); err != nil {
			return err
		}
	}

	return nil
}

func Test_Walk(t *testing.T) {
	t.Run("Array", func(t *testing.T) {
		count := 0
		arr := RandomDataArray(10)
		assert.Len(t, arr, 10)

		WalkArray(arr, func(item *TestData, key int) error {
			count++
			return nil
		})

		assert.Equal(t, 10, count)
	})

	t.Run("Name Map", func(t *testing.T) {
		count := 0
		m := RandomNameDataMap(10)

		WalkMap(m, func(item *TestData, key string) error {
			count++
			return nil
		})

		assert.Len(t, m, count)
	})

	t.Run("Age Map", func(t *testing.T) {
		count := 0
		m := RandomAgeDataMap(10)

		WalkMap(m, func(item *TestData, key int) error {
			count++
			return nil
		})

		assert.Len(t, m, count)
	})
}
