package main

func add_positive_numbers() int {
	result := 0
	for i := 0; i < 252; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	return result
}
