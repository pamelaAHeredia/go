package main

import (
	"fmt"
	"math"
)

func int_numbers(num int) int {
	switch {
	case num < -18:
		return num
	case num >= -18 && num <= -1:
		return num % 4
	case num >= 1 && num < 20:
		return num * num
	case num >= 20:
		return num * -1
	default:
		return 0
	}
}

func float_numbers(num float64) float64 {
	switch {
	case num < -18:
		return num
	case num >= -18 && num <= -1:
		return math.Mod(num, 4)
	case num >= 1 && num < 20:
		return num * num
	case num >= 20:
		return num * -1
	default:
		return 0
	}
}

func process_float_numbers() {
	for {
		var num float64

		fmt.Println("Ingresá un número entero distinto de cero: ")
		_, err := fmt.Scan(&num)

		if err != nil {
			fmt.Println("Número inválido")
			continue
		}

		if num == 0 {
			fmt.Println("Bai")
			break
		}

		fmt.Println(float_numbers(num))
	}
}

func process_int_numbers() {
	for {
		var num int

		fmt.Println("Ingresá un número entero distinto de cero: ")
		_, err := fmt.Scan(&num)

		if err != nil {
			fmt.Println("Número inválido")
			continue
		}

		if num == 0 {
			fmt.Println("Bai")
			break
		}

		fmt.Println(int_numbers(num))
	}
}
