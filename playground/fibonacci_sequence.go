package playground

// Fibonacci 
// 1,1,2,3,5,8...n
// return sum of numbers to n (inclusive)

func FibonacciSequence(n int) int {
    if n <= 0 {
        return 0
    }

    a, b := 0, 1 // first two Fibonacci numbers
    sum := 0

    for i := 1; i <= n; i++ {
        sum += b    // add current Fibonacci number
        a, b = b, a+b // shift to next Fibonacci
    }

    return sum
}
