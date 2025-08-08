package playground

import "fmt"

// FizzBuzz generates a string representation of the FizzBuzz sequence
// For numbers from 1 to n:
//   - If the number is divisible by 3, print "Fizz"
//   - If the number is divisible by 5, print "Buzz"
//   - If the number is divisible by both 3 and 5, print "FizzBuzz"
//   - Otherwise, print the number itself

func FizzBuzz(n int) {
	for value := 1; value < n; value++ {
		if value%15 == 0 {
			fmt.Print("FizzBuzz", ", ")
		} else if value%3 == 0 {
			fmt.Print("Fizz", ", ")
		} else if value%5 == 0 {
			fmt.Print("Buzz", ", ")
		} else {
			fmt.Print(value, ", ")
		}
	}
	fmt.Print(n) // Print the last number n
}
