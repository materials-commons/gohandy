package dir

// flattenState stores the state of the flatten progress.
type flattenState struct {
	all []*FileInfo
}

// Create a type for sorting at the end.
type files []*FileInfo

// Flatten takes a Directory and flattens it into a list of file objects
// sorted by full path. It does this for the entire set of files, including
// files in sub directories.
func (d *Directory) Flatten() []*FileInfo {
	state := &flattenState{
		all: []*FileInfo{},
	}
	state.flatten(d)
	return state.all
}

// flatten does the actual work of flattening the directory files into
// a list and descending down through all the sub directories.
func (s *flattenState) flatten(d *Directory) {
	// The entries from the walk are sorted lexographically. So,
	// we first add our directory on, then we add the files, then
	// we recurse for each subdirectory doing the same.
	s.all = append(s.all, &d.Info)
	for _, file := range d.Files {
		s.all = append(s.all, file)
	}
	for _, dir := range d.SubDirectories {
		s.flatten(dir)
	}
}
