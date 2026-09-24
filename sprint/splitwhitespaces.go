package sprint 
func SplitWhitespaces(s string) []string {
	var result []string
	word := ""

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}
