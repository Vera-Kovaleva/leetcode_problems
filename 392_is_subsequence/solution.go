package main

func isSubsequence(s string, t string) bool {
	tidx, sindx := 0, 0
    for tidx<len(t) && sindx<len(s) {
        if s[sindx] == t[tidx] {
            sindx++
            tidx++
        } else {
            tidx++
        }
    }
    if sindx==len(s) {
        return true
    }
	return false
}

/*
func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		if s[0] == t[i] {
			return isSubsequence(s[1:], t[i+1:])
		}
	}

	return false
}*/
