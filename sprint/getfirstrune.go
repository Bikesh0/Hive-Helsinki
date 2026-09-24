package sprint

func GetFirstRune(s string) rune {
	for _, i := range s {
		return i
	}

	return 0
}
