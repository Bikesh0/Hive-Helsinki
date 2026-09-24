package sprint

func ToThePowerIterative(n int, power int) int {

	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
    a := 1
	for i := power ; i > 0 ; i-- {

		a *= n
	}

 return a

}