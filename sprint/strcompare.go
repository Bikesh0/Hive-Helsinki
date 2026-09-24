package sprint

func StrCompare(a, b string) int {
	min := len(a)
	if len(b) < min {
		min = len(b)
	}

	for i := 0; i < min; i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}

	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}

	return 0
}
