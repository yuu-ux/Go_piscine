package piscine

func count(array []int) int {
	var len int = 0
	for range array {
		len++
	}
	return len
}
func SortIntegerTable(table []int) {
	var length = count(table)
	for i := 0; i < length - 1; i++ {
		for j := length - 1; i < j; j-- {
			if table[j-1] > table[j] {
				table[j-1], table[j] = table[j], table[j-1]
			}
		}
	}
}
