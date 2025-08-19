package main

import (
	"fmt"
	"github.com/rocklessg/go-algo/playground"
)

func main() {
	// num_in_list test
	//result := add(3, 4)
	// fmt.Println(playground.NumInList([]int{1, 2, 3, 4, 5}, 3)) // should return true
	// fmt.Println(playground.NumInList([]int{1, 2, 3, 4, 5}, 6)) // should return false

	// sum test
	// fmt.Println(playground.Sum([]int {1, 2, 3, 4, 5})) // should return 15
	// fmt.Println(playground.Sum([]int {10, 20, 30})) // should return 60

	// ReverseString test
	// fmt.Println(playground.ReverseString3("cat"))      // should return "tac"
	// fmt.Println(playground.ReverseString3("alphabet")) // should return "tebahpla"

	// FizzBuzz test
	// fmt.Println("sequence:")
	// playground.FizzBuzz(100)

	// NumberToBase test
	// fmt.Println("Number to Base:")
	// fmt.Println(playground.NumberToBase(10, 2))  // should return "1010"
	// fmt.Println(playground.NumberToBase(255, 16)) // should return "FF"
	// fmt.Println(playground.NumberToBase(100, 8))  // should return "144"
	// fmt.Println(playground.NumberToBase(0, 10))   // should return "0"
	// fmt.Println(playground.NumberToBase(123, 36)) // should return "3F

	// New logic test
	// fmt.Println("New logic Number to Base:")
	// fmt.Println(playground.NumberToBaseNew(10, 2))  // should return "1010"
	// fmt.Println(playground.NumberToBaseNew(255, 16)) // should return "FF"
	// fmt.Println(playground.NumberToBaseNew(100, 8))  // should return "144"
	// fmt.Println(playground.NumberToBaseNew(0, 10))   // should return "0"
	// fmt.Println(playground.NumberToBaseNew(123, 36)) // should return "3F

	// Fibonacci test
	fmt.Println("Fibonacci Sequence:")
	fmt.Println(playground.FibonacciSequence(5)) // should return 12
	fmt.Println(playground.FibonacciSequence(6)) // should return 20
	fmt.Println(playground.FibonacciSequence(7)) // should return 33
	fmt.Println(playground.FibonacciSequence(8)) // should return 54
	fmt.Println(playground.FibonacciSequence(9)) // should return 88
	fmt.Println(playground.FibonacciSequence(10)) // should return 143

}