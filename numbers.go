package rei

var (
	// numToRomanMap is used by `numberToRoman` to build Roman numerals.
	numToRomanMap = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	// romanNumOrder is used by `numberToRoman` to build Roman numerals.
	romanNumOrder = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	// numToRomanSmall is used by `numberToRoman` for fast Roman footnote generation.
	numToRomanSmall = map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
	}
)

// NumberToRoman converts an integer to a roman numeral
// adapted from https://pencilprogrammer.com/python-programs/convert-integer-to-roman-numerals/
func NumberToRoman(num int) string {
	if num <= 14 {
		return numToRomanSmall[num]
	}
	res := ""
	for _, v := range romanNumOrder {
		if num != 0 {
			quot := num / v
			if quot != 0 {
				for x := 0; x < quot; x++ {
					res += numToRomanMap[v]
				}
			}
			num %= v
		}
	}
	return res
}
