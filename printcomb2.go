package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var num1 rune
	var num2 rune
	var num3 rune
	var num4 rune

	for num1 = '0'; num1 <= '9'; num1++ {
		for num2 = '0'; num2 <= '9'; num2++ {
			for num3 = '0'; num3 <= '9'; num3++ {
				for num4 = '0'; num4 <= '9'; num4++ {
					if num1 == num3 && num2 == num4 {
					} else {
						z01.PrintRune(num1)
						z01.PrintRune(num2)
						z01.PrintRune(' ')
						z01.PrintRune(num3)
						z01.PrintRune(num4)
						if num1 == '9' && num2 == '9' && num3 == '9' && num4 == '8' {
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
