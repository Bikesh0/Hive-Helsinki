package sprint


type Student struct {
	Name   string
	Grades []int
}

func TopStudent(students []Student) Student {
	if len(students) == 0 {
		return Student{}
	}

	top := students[0]
	topAvg := average(top.Grades)

	for _, student := range students[1:] {
		avg := average(student.Grades)
		if avg > topAvg {
			top = student
			topAvg = avg
		}
	}

	return top
}

func average(grades []int) float64 {
	if len(grades) == 0 {
		return 0
	}

	sum := 0
	for _, grade := range grades {
		sum += grade
	}

	return float64(sum) / float64(len(grades))
}
