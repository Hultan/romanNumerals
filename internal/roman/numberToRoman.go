package roman

// I = 1 V = 5 X = 10 L = 50 C = 100 D = 500 M = 1000

type Digits int

const (
	Ones Digits = iota
	Tens
	Hundreds
	Thousands
)

func ToRoman(num int) string {
	d := Ones
	roman := ""
	for num > 0 {
		// Get the remainder when dividing with ten
		r := num % 10
		// Get roman numerals for the remainder, and add it to the result
		roman = getRomanFor(d, r) + roman
		// Divide by ten
		num = num / 10
		// Increase digit value
		d++
	}

	return roman
}

func getRomanFor(d Digits, num int) string {
	var low, middle, high string

	switch d {
	case Ones:
		low = "I"
		middle = "V"
		high = "X"
	case Tens:
		low = "X"
		middle = "L"
		high = "C"
	case Hundreds:
		low = "C"
		middle = "D"
		high = "M"
	case Thousands:
		result := ""
		for i := 0; i < num; i++ {
			result += "M"
		}
		return result
	}

	switch num {
	case 0:
		return ""
	case 1:
		return low
	case 2:
		return low + low
	case 3:
		return low + low + low
	case 4:
		return low + middle
	case 5:
		return middle
	case 6:
		return middle + low
	case 7:
		return middle + low + low
	case 8:
		return middle + low + low + low
	case 9:
		return low + high
	default:
		panic("ERROR")
	}
}
