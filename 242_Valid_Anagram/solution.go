package main

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

