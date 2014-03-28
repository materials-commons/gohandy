package arrays

// Strings allows us to limit our filter to just arrays of strings
type Strings struct{}

// Filter filters an array of strings.
func (s Strings) Filter(in []string, keep func(item string) bool) []string {
	var out []string
	for _, item := range in {
		if keep(item) {
			out = append(out, item)
		}
	}

	return out
}










