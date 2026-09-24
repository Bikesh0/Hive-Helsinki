package sprint 


func NbrBase(n int, base string) string {
	// Validate the base
	if len(base) < 2 {
		return "NV"
	}

	seen := make(map[rune]bool)

	for _, c := range base {
		if c == '+' || c == '-' {
			return "NV"
		}
		if seen[c] {
			return "NV"
		}
		seen[c] = true
	}

	// Handle zero
	if n == 0 {
		return string(base[0])
	}

	negative := n < 0
	if negative {
		n = -n
	}

	baseLen := len([]rune(base))
	digits := []rune(base)

	var result []rune

	for n > 0 {
		remainder := n % baseLen
		result = append(result, digits[remainder])
		n /= baseLen
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	if negative {
		return "-" + string(result)
	}

	return string(result)
}
