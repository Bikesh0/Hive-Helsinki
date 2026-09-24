package sprint

func AlphaNumber(n int) string {
	if n < 0 {
		return "-" + AlphaNumber(-n)
	}

	if n == 0 {
		return "a"
	}

	result := ""

	for n > 0 {
		digit := n % 10
		result = string(rune('a'+digit)) + result
		n /= 10
	}

	return result
}
