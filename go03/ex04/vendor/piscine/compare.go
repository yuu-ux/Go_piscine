package piscine

func StrLen(s string) int {
        var cnt int
        for range s {
                cnt++
        }
        return cnt
}

func Compare(a, b string) int {
	aLen := StrLen(a)-1
	bLen := StrLen(b)-1
	j := 0
	i := 0
	if a == "" && b == "" {
		return 0
	} else if a == "" {
		return -1
	} else if b == "" {
		return 1
	}

	for i < aLen && j < bLen {
		if a[i] == b[j] {
			i++
			j++
		} else {
			break
		}
	}
	if i == aLen && j == bLen {
		return 0
	} else {
		if a[i] < b[j] || j < bLen {
			return -1
		} else if a[i] > b[j] || i < aLen {
			return 1
		}
	}
	return 0
}

