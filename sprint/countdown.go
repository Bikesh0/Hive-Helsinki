package sprint

func Countdown(n int) string {
	result := ""

	for i := n; i > 0; i -= 2 {
		if result != "" {
			result += ", "
		}
		result += string(rune('0' + i))
	}

	if result != "" {
		result += ", "
	}

	return result + "0!"
}
