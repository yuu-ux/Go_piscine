package piscine

func IsPrintable(s string) bool {
    for _, c := range s {
        if !(32 <= c && 126 >= c) {
            return false
        }
    }
    return true
}
