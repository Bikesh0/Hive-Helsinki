package sprint

func IntVsFloat(i int, f float32) string {
	if float32(i) > f {
		return "Integer"
	}
	if float32(i) < f {
		return "Float"
	}
	return "Same"
}
