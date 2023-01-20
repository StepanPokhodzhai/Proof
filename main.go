package main

import (
	"fmt"
)

func main() {
	numbers := map[int]string{
		0:   "",
		1:   "один",
		2:   "два",
		3:   "три",
		4:   "чотири",
		5:   "п'ять",
		6:   "шість",
		7:   "сім",
		8:   "вісім",
		9:   "дев'ять",
		10:  "десять",
		11:  "одинадцять",
		12:  "дванадцять",
		13:  "тринадцять",
		14:  "чотирнадцять",
		15:  "п'ятнадцять",
		16:  "шістнадцять",
		17:  "сімнадцять",
		18:  "вісімнадцять",
		19:  "дев'ятнадцять",
		20:  "двадцять ",
		30:  "тридцять ",
		40:  "сорок ",
		50:  "п'ятдесят ",
		60:  "шістдесят ",
		70:  "сімдесят ",
		80:  "вісімдесят ",
		90:  "дев'яносто ",
		100: "сто ",
		200: "двісті ",
		300: "триста ",
		400: "чотириста ",
		500: "п'ятсот ",
		600: "шістсот ",
		700: "сімсот ",
		800: "вісімсот ",
		900: "дев'ятсот ",
	}

	var digit, decimal, hundred int
	fmt.Scan(&digit)
	if _, ok := numbers[digit]; ok {
		fmt.Println(numbers[digit])
	} else {
		hundred, digit = parse(100, digit)
		if _, ok := numbers[digit]; ok {
			fmt.Print(numbers[hundred], numbers[digit])
		} else {
			decimal, digit = parse(10, digit)
			fmt.Print(numbers[hundred], numbers[decimal], numbers[digit])
		}
	}
}
func parse(check int, number int) (int, int) {
	var res int
	if number >= check {
		res = number / check
		number %= res * check
		res *= check
	}
	return res, number
}
