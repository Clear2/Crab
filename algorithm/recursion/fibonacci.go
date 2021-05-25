package main

func fib(n int) int{
	var result int
	if n < 2 {
		result = n
	}
	result = fib(n-1) + fib(n-2)
	return result
}
func main() {

}