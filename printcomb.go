package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var first_number rune
	var second_number rune
	var third_number rune

	for first_number = '0'; first_number <= '9'; first_number++ {
		for second_number = first_number + 1; second_number <= '9'; second_number++ {
			for third_number = second_number + 1; third_number <= '9'; third_number++ {

				z01.PrintRune(first_number)
				z01.PrintRune(second_number)
				z01.PrintRune(third_number)

				if first_number == '7' && second_number == '8' && third_number == '9' {
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
