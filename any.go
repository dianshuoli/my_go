package main

import "fmt"

type Num interface {
	int | float64 | float32 | string
}

func sum[T Num](a, b T) T {
	return a + b
}

func main() {
	a, b := 3, 5
	c, d := 4., 67.
	e, f := "hello ", "golang"
	fmt.Println(sum(a, b))
	fmt.Println("sum of float", sum(c, d))
	fmt.Println("sum of string", sum(e, f))
}
