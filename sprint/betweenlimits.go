package sprint

func BetweenLimits(from, to rune) string {

	if from > to {
		from, to = to , from // swapping
	}
	result := ""

	for i := from + 1; i < to; i++ {

		result += string (i)
	}
    return result
}