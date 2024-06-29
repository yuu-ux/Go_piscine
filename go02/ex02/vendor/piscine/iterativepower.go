package piscine

func IterativePower(nb int, power int) int {
	var result int = 1
	if power < 0 {
		return 0
	}
	
	if power == 0 {
		return 1
	}

	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}
