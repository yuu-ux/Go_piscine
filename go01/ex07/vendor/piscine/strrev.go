
package piscine 

func StrLen(s string) int {
	var cnt int
	for _ = range s {
		cnt++
	}
	return cnt
}

func StrRev(s string) string {
	runes := []rune(s)
	for i, j := 0, StrLen(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] 
	}
	return string(runes)
}