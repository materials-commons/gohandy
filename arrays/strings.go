package arrays

// Strings allows us to limit our filter to just arrays of strings
type strings struct{}

// Strings allows easy access to the functions that operate on a list of strings
var Strings strings

// Filter filters an array of strings.
func (s strings) Filter(in []string, keep func(item string) bool) []string {
	var out []string
	for _, item := range in {
		if keep(item) {
			out = append(out, item)
		}
	}

	return out
}

// Remove filters an array by removing all matching items
func (s strings) Remove(in []string, remove ...string) []string {
	return s.Filter(in, func(item string) bool {
		found := false
		for _, removeItem := range remove {
			if removeItem == item {
				found = true
			}
		}
		return !found
	})
}










