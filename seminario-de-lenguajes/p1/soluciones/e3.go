package main

import "fmt"

func exercise_3() {
	/* integers */
	var zz int = 0
	x := 10
	var z int = x
	var y int = x + 1
	const n = 5001
	/* float */
	var e float32 = 6
	var f float32 = e

	fmt.Printf("zz: %d\n", zz)
	fmt.Printf("z: %d\n", z)
	fmt.Printf("y: %d\n", y)
	fmt.Printf("f: %f\n", f)
	fmt.Printf("n: %d\n", n)
}

// output
// zz: 0
// z: 10
// y: 11
// f: 6.000000
// n: 5001
