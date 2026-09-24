/* Instructions
Create a function that takes a positive integer as 
input and returns the corresponding number of letters 
from the Latin alphabet. The input integer won't be 
larger than the alphabet's length. */

package sprint

func AlphabetMastery(n int) string{
	a := ""


	for i := 0; i < n; i++ {

		a += string (rune('a' + i))

		
	}
	return a
}