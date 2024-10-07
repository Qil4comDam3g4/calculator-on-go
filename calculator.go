package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var romanMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"L":    50,
	"C":    100,
	"D":    500,
	"M":    1000,
}

func romanToArabic(roman string) int {
	roman = strings.ToUpper(roman)
	result := 0
	i := 0
	for i < len(roman) {
		s1 := romanMap[string(roman[i])]
		if i+1 < len(roman) {
			s2 := romanMap[string(roman[i+1])]
			if s1 >= s2 {
				result = result + s1
				i = i + 1
			} else {
				result = result + s2 - s1
				i = i + 2
			}
		} else {
			result = result + s1
			i = i + 1
		}
	}
	if result > 10 || result <= 0 {
		panic("roman numeral out of supported range (1-10)")
	}
	return result
}

func arabicToRoman(num int) string {
	if num < 1 {
		panic("roman numerals cannot be zero or negative")
	}
	conversions := []struct {
		value int
		digit string
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

	var result strings.Builder
	for _, conversion := range conversions {
		for num >= conversion.value {
			result.WriteString(conversion.digit)
			num -= conversion.value
		}
	}
	return result.String()
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("division by zero")
		}
		return a / b
	default:
		panic("invalid operator")
	}
}

func processInput(input string) interface{} {
	re := regexp.MustCompile(`^(\d+|[IVXLCDM]+)\s*([+\-*/])\s*(\d+|[IVXLCDM]+)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) == 0 {
		panic("invalid input format")
	}

	first, operator, second := matches[1], matches[2], matches[3]
	var a, b int
	var isRoman bool

	if _, err := strconv.Atoi(first); err == nil {
		a, _ = strconv.Atoi(first)
		if a < 1 || a > 10 {
			panic("numbers must be between 1 and 10")
		}
	} else {
		isRoman = true
		a = romanToArabic(first)
	}

	if _, err := strconv.Atoi(second); err == nil {
		if isRoman {
			panic("mixed number formats")
		}
		b, _ = strconv.Atoi(second)

		if b < 1 || b > 10 {
			panic("numbers must be between 1 and 10")
		}
	} else {
		if !isRoman {
			panic("mixed number formats")
		}
		b = romanToArabic(second)
	}

	result := calculate(a, b, operator)

	if isRoman {
		if result < 1 {
			panic("roman numeral result cannot be less than 1")
		}
		return arabicToRoman(result)
	} else {
		return result
	}
}

func main() {
	var input string
	fmt.Println("Enter an expression (e.g., I + II or 1 + 2):")

	for {
		fmt.Scanln(&input)
		input = strings.ReplaceAll(input, " ", "")

		if input == "exit" {
			break
		}

		result := processInput(input)
		fmt.Println("Result:", result)

		fmt.Print("Enter next expression (or 'exit' to quit): ")
	}
}
