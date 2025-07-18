package main

//https://leetcode.com/problems/isomorphic-strings/description/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	pairs1 := make(map[byte]byte)
	pairs2 := make(map[byte]byte)

	for i := 0; i < len(t); i++ {
		sValue := s[i]
		tValue := t[i]

		valueT, existT := pairs1[sValue]
		valueS, existS := pairs2[tValue]

		if (!existT) && (!existS) {
			pairs1[sValue] = tValue
			pairs2[tValue] = sValue
		} else if valueT != tValue || valueS != sValue {
			return false
		}

	}
	return true
}
