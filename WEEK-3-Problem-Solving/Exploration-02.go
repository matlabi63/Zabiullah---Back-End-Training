package main

import "fmt"

func intToRoman(num int) string {
	
	vals := []struct {
		val  int
		sym  string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	
	roman := ""
	for _, v := range vals {
		for num >= v.val {
			roman += v.sym
			num -= v.val
		}
	}

	return roman
}

func main() {
	
	fmt.Println(intToRoman(4))    // IV
	fmt.Println(intToRoman(9))    // IX
	fmt.Println(intToRoman(23))   // XXIII
	fmt.Println(intToRoman(2021)) // MMXXI
	fmt.Println(intToRoman(1646)) // MDCXLVI
}
