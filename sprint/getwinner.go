package sprint

type Contestant struct {
	Name string
	Scores []int
}

func GetWinner(c1, c2 Contestant) string {

	s1 := 0
	s2:= 0

	for _, a:= range c1.Scores {
		s1 += a

	}
	for _, b := range c2.Scores {
		s2 += b
	}
	if s1>s2 {
		return c1.Name
	} else {
		return c2.Name
	}


}