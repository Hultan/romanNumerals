package roman

var romanNumerals = map[byte]int{0: 0, 73: 1, 86: 5, 88: 10, 76: 50, 67: 100, 68: 500, 77: 1000}

func ToArabic(roman string) int {
	var num int
	for j := 0; j < len(roman); j++ {
		// Get the next "digit" and the "digit" after that (if it exists)
		n1 := romanNumerals[getNumber(roman, j)]
		n2 := romanNumerals[getNumber(roman, j+1)]

		// Check if n2 is bigger than n1, if so, we need to
		// subtract the numbers, and skip an extra step
		if n1 >= n2 || n2 == 0 {
			num += n1
		} else {
			num += n2 - n1
			j++
		}
	}
	return num
}

func getNumber(roman string, num int) byte {
	if num > len(roman)-1 {
		return 0
	}
	return roman[num]
}
