package main

import (
	"fmt"
)

func main() {
	numbers := map[int]string{
		0:  "",
		1:  "сто ",
		2:  "двісті ",
		3:  "триста ",
		4:  "чотириста ",
		5:  "п'ятсот ",
		6:  "шістсот ",
		7:  "сімсот ",
		8:  "вісімсот ",
		9:  "дев'ятсот ",
		10: "",
		11: "десять ",
		12: "двадцять ",
		13: "тридцять ",
		14: "сорок ",
		15: "п'ятдесят ",
		16: "шістдесят ",
		17: "сімдесят ",
		18: "вісімдесят ",
		19: "дев'яносто ",
		20: "",
		21: "один",
		22: "два",
		23: "три",
		24: "чотири",
		25: "п'ять",
		26: "шість",
		27: "сім",
		28: "вісім",
		29: "дев'ять",
	}

	var digit, decimal, hundred int
	fmt.Scan(&digit)
	hundred, digit = parse(100, digit)
	decimal, digit = parse(10, digit)
	fmt.Print(numbers[hundred], numbers[decimal+10], numbers[digit+20])
}

func parse(check int, number int) (int, int) {
	var res int
	if number >= check {
		res = number / check
		number %= res * check
	}
	return res, number
}
