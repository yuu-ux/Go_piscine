package piscine

func StrLen(s string) int {
	var cnt int
	for _ = range s {
		cnt++
	}
	return cnt
}

func LastRune(s string) rune {
	len := StrLen(s)
	return rune(s[len - 1])
}
