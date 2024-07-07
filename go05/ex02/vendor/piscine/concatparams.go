package piscine

func CountElement(s []string) int {
	var cnt int
	for i, _ := range s {
		cnt = i
	}
	return cnt
}

func ConcatParams(args []string) string {
	ElementNum := CountElement(args)
	var result string
	var i int
	flag := true

	for i < ElementNum+1 {
		if flag {
			result += args[i]
			flag=false
			i++
		} else {
			result += "\n"
			flag=true
		}
	}
	return result
}
