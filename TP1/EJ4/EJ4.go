package main

import "fmt"

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func main() {
	var n int
	fmt.Println("Imgrese un nuemero:")
	fmt.Scan(&n)
	fmt.Println("La sucecion Fibonacci del numero ingresado es: ", Fibonacci(n))
}
