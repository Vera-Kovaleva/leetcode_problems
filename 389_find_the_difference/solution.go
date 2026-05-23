package main

func findTheDifference(s string, t string) byte {
	out := t[len(t)-1]

	for i := 0; i < len(s); i++ {
		out ^= s[i] ^ t[i]
	}
	return out
}
