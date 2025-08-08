package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}

	for i := 0; i < len(s); i++ {
		if counter[s[i]-'a'] != 0 {
			return false
		}
	}
	return true
}

/*
import (
    "sort"
    "strings"
)

func SortString(given string) string {
    result := strings.Split(given, "")
    sort.Strings(result)
    return strings.Join(result, "")
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
        return false
    }
	return SortString(s) == SortString(t)
}
*/
