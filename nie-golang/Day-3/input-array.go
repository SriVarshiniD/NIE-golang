package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter the number of array elements: ")
	fmt.Scan(&num)

	even := make([]int, num)
	fmt.Println("Enter ", num, "even numbers:")

	for i := 0; i < num; i++ {
		fmt.Scan(&even[i])
	}

	fmt.Println("Even Numbers are as follows: ", even)
}
