package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	
	if nb == 0 || nb == 1 {
		return nb
	}
	var ret int
	for i := 1; i < nb; i++ {
		if nb / i == i {
			ret =  i
		}
	}
	return ret
}