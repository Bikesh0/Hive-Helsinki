package sprint 

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	return a * b / GCD(a, b)
}
