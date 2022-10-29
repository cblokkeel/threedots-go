package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func Sum(n1, n2, n3, n4, n5 int) int {
	return n1 + n2 + n3 + n4 + n5
}
