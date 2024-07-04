package piscine

func StrLen(s string) int {
        var cnt int
        for _ = range s {
                cnt++
        }
        return cnt
}

func NRune(s string, n int) rune {
	cnt := StrLen(s)
	index := 0

	if cnt < n || n < 1 {
		return 0
	}

	for i := 0; i < n; i++ {
		index++
	}
	return rune(s[index-1])
}
