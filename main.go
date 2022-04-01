package main

import "fmt"

func GenericFunc[age int64 | float32](myAge age) {
	fmt.Printf("Age: %v, Type: %T\n", myAge, myAge)
}

func GenericFunc2[value any](v value) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

type SignedInteger interface {
	int8 | int16 | int32 | int64
}

type Float interface {
	float32 | float64
}

type Number interface {
	SignedInteger | Float
}

func BubbleSort[N Number](values []N) {
	n := len(values)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if values[i] > values[i+1] {
				values[i], values[i+1] = values[i+1], values[i]
				swapped = true
			}
		}
	}
}

func main() {
	var a int64 = 18
	var b float32 = 23
	var c string = "55"

	GenericFunc(a)
	GenericFunc(b)
	GenericFunc2(c)

	integers := []int16{10, 5, 2, 3, 11, 0}
	BubbleSort(integers)
	fmt.Println("Sorted integers: ", integers)

	floats := []float32{8.4, -1, 0.9, 3.1415, 10}
	BubbleSort(floats)
	fmt.Println("Sorted floats: ", floats)
}
