package sorter

import (
	"sort"
)

type IntSorter []int

func (s IntSorter) Len() int           { return len(s) }
func (s IntSorter) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type FloatSorter []float64

func (s FloatSorter) Len() int           { return len(s) }
func (s FloatSorter) Less(i, j int) bool { return s[i] < s[j] }
func (s FloatSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type StringSorter []string

func (s StringSorter) Len() int           { return len(s) }
func (s StringSorter) Less(i, j int) bool { return s[i] < s[j] }
func (s StringSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type MixedSorter []interface{}

func SortInt(input []int) {
	sort.Sort(IntSorter(input))
}

func SortFloat(input []float64) {
	sort.Sort(FloatSorter(input))
}

func SortString(input []string) {
	sort.Sort(StringSorter(input))
}

func SortMixedElements(input MixedSorter) {
	var intSlice []int
	var floatSlice []float64
	var stringSlice []string

	for _, elem := range input {
		switch v := elem.(type) {
		case int:
			intSlice = append(intSlice, v)
		case float64:
			floatSlice = append(floatSlice, v)
		case string:
			stringSlice = append(stringSlice, v)
		}
	}

	sort.Ints(intSlice)
	sort.Float64s(floatSlice)
	sort.Strings(stringSlice)

	var result MixedSorter
	for _, val := range intSlice {
		result = append(result, val)
	}
	for _, val := range floatSlice {
		result = append(result, val)
	}
	for _, val := range stringSlice {
		result = append(result, val)
	}

	copy(input, result)
}
