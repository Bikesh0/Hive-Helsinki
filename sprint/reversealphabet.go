package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}

	result := ""

	for i := int('z'); i >= int('a'); i -= step {
		result += string(rune(i))
	}

	return result
}
