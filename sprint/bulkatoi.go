package sprint 

func StrToInt(s string) int {
	sign := 1
	result := 0

	if s == "" {
		return 0
	}

	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

		result = result*10 + int(s[i]-'0')
	}

	return result * sign
}

func BulkAtoi(arr []string) []int {
	result := []int{}

	for _, value := range arr {
		number := StrToInt(value)
		result = append(result, number)
	}

	return result
}
