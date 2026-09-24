package sprint

func FactorialIterative(n int) int {
     factor := 0
	if n < 0 {

		return 0
	}

	for i = n; i > 0; i ++ {
		factor *= i
       


	}
	return factor

}