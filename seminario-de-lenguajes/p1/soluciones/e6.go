package main

import (
	"fmt"
)

func ask_int_numbers() (int, int, error) {
	var a, b int
	fmt.Println("Ingrese un número entero: ")
	_, err := fmt.Scan(&a)

	if err != nil {
		fmt.Println("Número inválido")
		return 0, 0, err
	}

	fmt.Println("Ingrese otro número entero: ")
	_, err = fmt.Scan(&b)

	if err != nil {
		fmt.Println("Número inválido")
		return 0, 0, err
	}

	return a, b, nil
}

func divide_int() {
	for {
		a, b, err := ask_int_numbers()

		if err != nil {
			fmt.Println("Número inválido")
			return
		}

		var max = a
		var min = b

		if b > a {
			max = b
			min = a
		}

		if min == 0 {
			fmt.Println("No se puede dividir por cero.")
			return
		}

		fmt.Printf("%d / %d = %d\n", max, min, max/min)
	}
}

func ask_float_numbers() (float64, float64, error) {
	var a, b float64
	fmt.Println("Ingrese un número real: ")
	_, err := fmt.Scan(&a)

	if err != nil {
		fmt.Println("Número inválido")
		return 0, 0, err
	}

	fmt.Println("Ingrese otro número real: ")
	_, err = fmt.Scan(&b)

	if err != nil {
		fmt.Println("Número inválido")
		return 0, 0, err
	}

	return a, b, nil
}

func divide_float() {
	for {
		a, b, err := ask_float_numbers()

		if err != nil {
			fmt.Println("Número inválido")
			return
		}

		var max = a
		var min = b

		if b > a {
			max = b
			min = a
		}

		if min == 0 {
			fmt.Println("No se puede dividir por cero.")
			return
		}

		fmt.Printf("%v / %v = %v\n", max, min, max/min)
	}

}
