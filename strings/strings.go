package strings

// Filter filters an array of strings.
func Filter(in []string, keep func(item string) bool) []string {
	var out []string
	for _, item := range in {
		if keep(item) {
			out = append(out, item)
		}
	}

	return out
}
