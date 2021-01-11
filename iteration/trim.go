package iteration

func TrimFunc(s string, f func(rune) bool) string {
	var trimed string
	for _, char := range s {
		if f(char) {
			continue
		}
		trimed += string(char)
	}

	return trimed
}
