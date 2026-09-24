package sprint 

func ToCapitalCase(s string) string {
	r := []rune(s)
	firstletter := true

	for i, ch := range r {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			if firstletter {
				if ch >= 'a' && ch <= 'z' {
					r[i] = ch - ('a' - 'A')
				}
				firstletter = false
			} else {
				if ch >= 'A' && ch <= 'Z' {
					r[i] = ch + ('a' - 'A')
				}
			}
		} else {
			firstletter = true
		}
	}

	return string(r)
}


	


