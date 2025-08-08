package main

func longestCommonPrefix(strs []string) string {
	pref := ""
	for i := 0; i < len(strs[0]); i++ {
		curSymb := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 {
				return pref
			}
			if curSymb != strs[j][i] {
				return pref
			}
		}
		pref += string(curSymb)
	}
	return pref
}
