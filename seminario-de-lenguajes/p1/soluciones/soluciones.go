package main

import (
	"fmt"
)

func main() {
	menu()
}

func menu() {
	for {
		var input int

		fmt.Println("Lista de opciones:")
		fmt.Println(" 1 - Hello world\n 2- Correción sintáctica\n 3- Sumar números positivos\n 4- Números enteros\n 5- Números flotantes\n 0- salir")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Número inválido")
			continue
		}

		if input == 0 {
			fmt.Println("BAI")
			break
		}

		choose_option(input)
	}
}

func choose_option(opt int) {
	switch opt {
	case 1:
		hello_world()
	case 2:
		exercise_3()
	case 3:
		add_positive_numbers()
	case 4:
		process_int_numbers()
	case 5:
		process_float_numbers()
	default:
		fmt.Println("Número inválido")
	}
}
