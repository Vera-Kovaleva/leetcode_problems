package main

/*
func canConstruct(ransomNote string, magazine string) bool {
	helpMap := make(map[rune]int)
	for _, haveSymb := range magazine {
		helpMap[haveSymb]++
	}

	for _, needSymb := range ransomNote {
		if helpMap[needSymb] == 0 {
			return false
		}
		helpMap[needSymb] --
	}

	return true
}
*/

func canConstruct(ransomNote string, magazine string) bool {
	letterCount := make([]int, 26)
	for _, ch := range magazine {
		letterCount[ch - 'a'] ++
	}
	
	for _, ch := range ransomNote {
		letterCount[ch - 'a'] --
		if letterCount[ch - 'a'] < 0 {
			return false
		}
	}
	return true
	
}