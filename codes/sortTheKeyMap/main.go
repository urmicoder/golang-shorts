package main

import (
	"fmt"
	"sort"
)

// How do you sort the keys of a map?
// map ko sort nhi kar shakte par slice ko sort kar shakte hai, yisiliye pahele map ko slice me convert karke slice ko key ke base par sort kar lege aur fir slice se map me convert kar lege.
// we cannot sort a map, but we can sort a slice. Therefore, first, we will convert the map into a slice, sort the slice based on the key, and then convert it back to a map
func main() {
	m := map[string]int{"a": 1, "c": 2, "d": 4, "b": 3}
	arrayOfSlice := make([]string, 0, len(m))

	for k := range m {
		arrayOfSlice = append(arrayOfSlice, k)
		fmt.Println(arrayOfSlice)
	}
	sort.Strings(arrayOfSlice)
	for _, k := range arrayOfSlice {
		fmt.Println(k, m[k])
	}
}
