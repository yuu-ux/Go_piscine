package piscine

func Index(s string, toFind string) int {
	var index int = 0
	for _, c := range s {
		if c == toFind[0] {
			for _, c := range c {
				if c == toFind[i] {
					return index
				}
				break
			}
		}
	}
}
