package piscine

func CountElement(strs []string) int {
	var cnt int
	for i, _ :=  range strs {
		cnt = i
	}
	return cnt
}
func Join(strs []string, sep string) string {
	var result string
	ElementCount := CountElement(strs)
	flag := true
	var i int

	for i < ElementCount+1 {
		if flag {
			result += strs[i]
			flag = false
			i++
		} else {
			result += sep
			flag = true
		}
	}
	return result
}
