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
    var result []string
    var start int
    for {
        index := Index(s[start:], sep)
        if index == -1 {
            result = append(result, s[start:])
            break
        }
        result = append(result, s[start:start+index])
        start += index + StrLen(sep)
    }
    return result
}
