package piscine 

func IsPrime(nb int) bool {
	var flag int = 0 

	if nb <= 0 {
		return false
	}

	if nb == 1 {
		return false
	}
	
	for i := 2; i < nb; i++ {
		if nb % i == 0 {
			flag += 1
		}
	}
	return flag < 1
}