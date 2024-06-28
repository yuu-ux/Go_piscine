package piscine 

func count(table []int) int {
	var cnt int = 0
	for range table {
		cnt++
	}
	return cnt
}
func SortIntegerTable(table []int) {
	var len int = count(table)

	for i := 0; i < (len - 1); i++ {
		for j := (len - 1); i < j; j-- {
			if table[j - 1] > table[j] {
				table[j - 1], table[j] = table[j], table[j - 1]
			}
		}
	}
}