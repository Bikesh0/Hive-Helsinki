package sprint

func BalanceOut(arr []bool) []bool {
	count := 0

	for _, value := range arr {
		if value {
			count++
		} else {
			count--
		}
	}

	if count > 0 {
		for i := 0; i < count; i++ {
			arr = append(arr, false)
		}
	} else {
		for i := 0; i < -count; i++ {
			arr = append(arr, true)
		}
	}

	return arr
}
