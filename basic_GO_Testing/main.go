package main

import "fmt"

// Calculate return num + 2
func Calculate(num int) (result int) {
	result = num + 2
	return result
}

func main() {
	fmt.Print("The Output : ")
	fmt.Println(Calculate(2))

}
