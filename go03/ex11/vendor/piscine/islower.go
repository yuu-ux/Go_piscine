package piscine

func IsLower(s string) bool {
    for _, c := range s {
        if !('a' <= c && 'z' >= c) {
            return false
        }
    }
    return true
}
