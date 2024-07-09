package piscine

func StrLen(s string) int {
    var cnt int
    for range s {
        cnt++
    }
    return cnt
}

func Index(s string, toFind string) int {
    i := 0
    j := 0
    start := 0
    sLen := StrLen(s)

    if toFind == "" {
        return 0
    }
    for ; i < sLen; i++{
        if s[i] == toFind[0] {
            tLen := StrLen(toFind)
            for start, j = i, 0; start != sLen && j != tLen;{
                if s[start] == toFind[j] && j == tLen-1 {
                    return i
                } else if s[start] != toFind[j] {
                    break
                }
                start++
                j++
            }
        }
    }
    return -1
}

func Split(s, sep string) []string {
	cnt := StrLen(s)
	var temp string
	var result []string
	var index int
	var i int

	for index+i < cnt {
		i = 0
		index = Index(s, sep)
		if index == -1 {
			break
		}
		for _, c := range []rune(s)[:index] {
			temp += string(c)
			i++
		}
		result = append(result, temp)
		temp = ""
		s = string([]rune(s)[index:])
	}
	return result
}
