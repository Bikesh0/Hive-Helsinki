package sprint

func ShiftBy(r rune, step int) rune {

	return 'a' + (r - 'a' + rune(step)) % 26

	// rune ruins 

}
