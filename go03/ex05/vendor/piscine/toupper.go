package piscine

func ToUpper(s string) string {
	var result []string

	for _, c := range s {
		if 'a' <= c && 'z' >= c {
			result = c - 32
		}
	}
	return s
}

