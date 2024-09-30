package sortex

import (
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
	Age  int
}

var persons = []*Person{
	{"Paul", 50},
	{"Mary", 30},
	{"John", 30},
	{"Alice", 25},
	{"Bob", 25},
}

type (
	ByAgeIncreasing []*Person // define a type for sorting
	ByAgeDecreasing []*Person // define a type for sorting
)

func (p ByAgeIncreasing) Len() int {
	return len(p)
}

func (p ByAgeIncreasing) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByAgeIncreasing) Less(i, j int) bool {
	if p[i].Age == p[j].Age {
		return p[i].Name < p[j].Name
	}
	return p[i].Age < p[j].Age
}

func (p ByAgeDecreasing) Len() int {
	return len(p)
}

func (p ByAgeDecreasing) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByAgeDecreasing) Less(i, j int) bool {
	if p[i].Age == p[j].Age {
		return p[i].Name < p[j].Name
	}
	return p[i].Age > p[j].Age
}

// use 1.8 sort.Slice
// if age is the same, sort by name in lexicographical order
func SortPeopleByAgeIncreasing(t *testing.T) {
	sort.Slice(persons, func(i, j int) bool {
		if persons[i].Age == persons[j].Age {
			return persons[i].Name < persons[j].Name
		}
		return persons[i].Age < persons[j].Age
	})
	assert.Equal(t, []*Person{
		{"Alice", 25},
		{"Bob", 25},
		{"Mary", 30},
		{"John", 30},
		{"Paul", 50},
	}, persons)
}

// use 1.21 slices.SortFunc
// if age is the same, sort by name in lexicographical order
func SliceSortPeopleByAgeIncreasing(t *testing.T) {
	slices.SortFunc(persons, func(a, b *Person) int {
		if a.Age == b.Age {
			return strings.Compare(a.Name, b.Name)
		}
		return a.Age - b.Age
	})
	assert.Equal(t, []*Person{
		{"Alice", 25},
		{"Bob", 25},
		{"Mary", 30},
		{"John", 30},
		{"Paul", 50},
	}, persons)
}

// use 1.8 sort.Slice
// if age is the same, sort by name in lexicographical order
func SortPeopleByAgeDecreasing(t *testing.T) {
	sort.Slice(persons, func(i, j int) bool {
		if persons[i].Age == persons[j].Age {
			return persons[i].Name < persons[j].Name
		}
		return persons[i].Age > persons[j].Age
	})
	assert.Equal(t, []*Person{
		{"Paul", 50},
		{"John", 30},
		{"Mary", 30},
		{"Alice", 25},
		{"Bob", 25},
	}, persons)
}

// use 1.21 slices.SortFunc
// if age is the same, sort by name in lexicographical order
func SliceSortPeopleByAgeDecreasing(t *testing.T) {
	slices.SortFunc(persons, func(a, b *Person) int {
		if a.Age == b.Age {
			return strings.Compare(a.Name, b.Name)
		}
		return b.Age - a.Age
	})
	assert.Equal(t, []*Person{
		{"Alice", 25},
		{"Bob", 25},
		{"Mary", 30},
		{"John", 30},
		{"Paul", 50},
	}, persons)
}

func SortPeopleByAgeIncreasingInterface(t *testing.T) {
	sort.Sort(ByAgeIncreasing(persons))
	assert.Equal(t, []*Person{
		{"Alice", 25},
		{"Bob", 25},
		{"Mary", 30},
		{"John", 30},
		{"Paul", 50},
	}, persons)
}

func SortPeopleByAgeDecreasingInterface(t *testing.T) {
	sort.Sort(ByAgeDecreasing(persons))
	assert.Equal(t, []*Person{
		{"Paul", 50},
		{"John", 30},
		{"Mary", 30},
		{"Alice", 25},
		{"Bob", 25},
	}, persons)
}
