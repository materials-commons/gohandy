package dir

// flattenState stores the state of the flatten progress.
type flattenState struct {
	all []*File
}

// Flatten takes a Directory and flattens it into a list of file objects
// sorted by full path. It does this for the entire set of files, including
// files in sub directories.
func (d *Directory) Flatten() []*File {
	state := &flattenState{
		all: []*File{},
	}
	state.flatten(d)
	return state.all
}

// flatten does the actual work of flattening the directory files into
// a list and descending down through all the sub directories.
func (s *flattenState) flatten(d *Directory) {
	for _, file := range d.Files {
		s.all = append(s.all, file)
	}
	for _, dir := range d.SubDirectories {
		s.flatten(dir)
	}
}
