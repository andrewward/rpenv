package main

import (
	"sort"
)

func mapAsSortedSlice(m map[string]string) [][2]string {
	s := make([][2]string, 0)
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		s = append(s, [2]string{key, m[key]})
	}
	return s
}

func updateMap(originalMap map[string]string, newMap map[string]string) map[string]string {
	for key, value := range newMap {
		_, present := originalMap[key]
		if !present {
			originalMap[key] = value
		}
	}

	return originalMap
}
