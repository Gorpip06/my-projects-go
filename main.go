package main

import (
	"fmt"
)
func main() {
	ish, a, celev := vodim()
	raschet(ish, a, celev)
}
func vodim() (string, float64, string) {
	var b float64
	var ish1 string
	var celev1 string
	
	for  {
	fmt.Println("Введите исходную валюту(USD, EUR, RUB)")
	fmt.Scan(&ish1)
	if ish1 != "USD" && ish1 != "EUR" && ish1 != "RUB" {
		fmt.Println("Введено некорректное значение")
		continue
	}
	break
	}
	for {
	fmt.Println("Введите число для конвертации")
	_, err := fmt.Scan(&b)
	if err != nil {
		fmt.Println("Нужно ввести число")
		continue
	}
	break
}	
	for {
	fmt.Println("Введите целевую валюту, кроме", ish1)
	fmt.Scan(&celev1)
	if celev1 == ish1 || (celev1 != "USD" && celev1 != "EUR" && celev1 != "RUB") {
		fmt.Println("Введено некорректное значение")
		continue
	}
	break
	}
	return ish1, b, celev1
}

func raschet(ish string, a float64, celev string) {
	switch ish {
	case "USD":
		switch celev {
		case "RUB":
			fmt.Print(a * 80.30)
		case "EUR":
			fmt.Print(a * 0.89)
		}
	case "RUB":
		switch celev {
		case "EUR":
			fmt.Print(a * 0.011)
		case "USD":
			fmt.Print(a * 0.012)
		}
	case "EUR":
		switch celev {
		case "RUB":
			fmt.Print(a * 90.18)
		case "USD":
			fmt.Print(a * 1.12)
	}
		}
}
