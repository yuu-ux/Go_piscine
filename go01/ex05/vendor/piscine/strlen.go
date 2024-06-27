package piscine

func StrLen(s string) int {
	var cnt int
	for _ = range s {
		cnt++
	}
	return cnt
}