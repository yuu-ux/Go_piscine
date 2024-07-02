package piscine

func StrLen(s string) int {
        var cnt int
        for range s {
                cnt++
        }
        return cnt
}

func Compare(a, b string) int {
	aLen = StrLen(a)
	bLen = StrLen(b)
	j := 0

	for i := 0; i < aLen; i++ {
		if a[i] < b[j] {
			return -1
		} else if a[i] > b[j] {
			return 1
		} else {
			for j := 0; j < bLen; j++ {
						
			}
	}	
}
// 完全一致したら0を返す
// s1-s2でマイナスであれば-1を返す
// s1-s2でプラスであれば1を返す
