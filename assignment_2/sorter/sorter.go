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

func SortInt(slice []int) {
	sort.Sort(IntSorter(slice))
}

func SortFloat(slice []float64) {
	sort.Sort(FloatSorter(slice))
}

func SortString(slice []string) {
	sort.Sort(StringSorter(slice))
}

func SortMixedElements(slice []interface{}) {
	sort.Slice(slice, func(i, j int) bool {
		switch slice[i].(type) {
		case int:
			return slice[i].(int) < slice[j].(int)
		case float64:
			return slice[i].(float64) < slice[j].(float64)
		case string:
			return slice[i].(string) < slice[j].(string)
		}
		return false
	})
}
