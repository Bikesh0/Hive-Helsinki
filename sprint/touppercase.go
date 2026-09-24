package sprint

func ToUpperCase(s string) string {
    r := []rune(s)

    for i, ch := range r {
        if ch >= 'a' && ch <= 'z' {
            r[i] = ch - 'a' + 'A'
        }
    }
    return string(r)
}