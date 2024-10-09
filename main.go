package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(arr []string) [][]string {
	groupedAnagrams := make([][]string, 0)
	visited := make([]bool, len(arr))

	for i := 0; i < len(arr); i++ {
		if !visited[i] {
			anagramGroup := []string{arr[i]}
			visited[i] = true

			for j := i + 1; j < len(arr); j++ {
				if !visited[j] && isAnagram(arr[i], arr[j]) {
					anagramGroup = append(anagramGroup, arr[j])
					visited[j] = true
				}
			}

			groupedAnagrams = append(groupedAnagrams, anagramGroup)
		}
	}

	return groupedAnagrams
}

func isAnagram(s1, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	return string(r1) == string(r2)
}

func main() {
	arr := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := groupAnagrams(arr)
	fmt.Println(result)
}
