package piscine

func IsUpper(s string) bool {
    for _, c := range s {
        if !('A' <= c && 'Z' >= c) {
            return false
        }
    }
    return true
}
