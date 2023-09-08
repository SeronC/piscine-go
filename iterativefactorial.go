package piscine

func IterativeFactorial(nb int) int {
	result := nb
	if nb >= 0 && nb < 26 {
		if nb == 0 {
			result = 1
		} else {
			for i := nb - 1; i > 0; i-- {
				result = result * i
			}
		}
	} else {
		result = 0
	}
	return result
}
