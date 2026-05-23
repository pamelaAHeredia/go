package main

import "fmt"

func add_positive_numbers() {
	result := 0
	for i := 0; i < 252; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	fmt.Printf("%d\n", result)
}
