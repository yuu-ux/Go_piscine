package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	result := make([]int, min, max)
	j := 0
	for i := min; i < max; i++ {
		result[j] = i
		j++
	}
	return result
}
