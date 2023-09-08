package piscine

func RecursiveFactorial(nb int) int {
	result := nb
	if nb >= 0 && nb < 26 {
		if nb == 0 || nb == 1 {
			result = 1
		}
		if nb > 1 {
			return nb * RecursiveFactorial(nb-1)
		}
	} else {
		result = 0
	}
	return result
}
