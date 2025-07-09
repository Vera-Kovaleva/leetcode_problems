package main

//https://leetcode.com/problems/reverse-vowels-of-a-string/description/

func isVowel(c rune) bool {
	return c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' || c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func reverseVowels(s string) string {
	result := []rune(s)

	iStart := 0
	iEnd := len(s) - 1
	for iStart < iEnd {
		for !isVowel(result[iStart]) && iStart < iEnd {
			iStart++
		}
		for !isVowel(result[iEnd]) && iStart < iEnd {
			iEnd--
		}
		result[iEnd], result[iStart] = result[iStart], result[iEnd]
		iStart++
		iEnd--
	}

	return string(result)
}
