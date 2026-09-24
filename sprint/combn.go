package sprint



import "fmt"

func CombN(n int) []string {
	result := []string{}

	if n < 1 || n > 10 {
		return result
	}

	var generate func(start int, current string)

	generate = func(start int, current string) {
		if len(current) == n {
			result = append(result, current)
			return
		}

		for digit := start; digit <= 9; digit++ {
			generate(digit+1, current+fmt.Sprint(digit))
		}
	}

	generate(0, "")

	return result
}
