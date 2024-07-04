package piscine

func IsUpper(c rune) bool {
    return ('A' <= c && 'Z' >= c)
}
func ToLower(s string) string {
    var result string

    for _, c := range s {
        if IsUpper(c) {
            c += ('a' - 'A')
        }
        result += string(c)
    }
    return result
}
