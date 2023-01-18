package main

import "fmt"

func main() {
	var digit, decimal, hundred int
	fmt.Scanln(&digit)
	if digit >= 100 {
		hundred = digit / 100
		digit %= hundred * 100
	}
	if digit >= 10 {
		decimal = digit / 10
		digit %= decimal * 10
	}

	hundredArr := [10]string{"", "сто ", "двісті ", "триста ", "чотириста ", "п'ятсот ", "шістсот ", "сімсот ", "вісімсот ", "дев'ятсот "}
	decimalAarr := [10]string{"", "десять ", "двадцять ", "тридцять ", "сорок ", "п'ятдесят ", "шістдесят ", "сімдесят ", "вісімдесят ", "дев'яносто "}
	digitArr := [10]string{"", "один", "два", "три", "чотири", "п'ять", "шість", "сім", "вісім", "дев'ять"}
	print(hundredArr[hundred], decimalAarr[decimal], digitArr[digit])
}
