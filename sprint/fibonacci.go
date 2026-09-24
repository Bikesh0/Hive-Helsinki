/*Write a function that takes an integer n and returns the 
Fibonacci number at position n using a loop. The Fibonacci 
sequence starts with 0 and 1, and each subsequent number is 
the sum of the two preceding ones. For any negative input, 
return 0.*/
package sprint

func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}
