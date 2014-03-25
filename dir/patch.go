package dir

// PatchType denotes the kind of patch operation
type PatchType int

const (
	// PatchCreate created item
	PatchCreate PatchType = iota

	// PatchDelete deleted item
	PatchDelete
)

// Patch is an instance of a difference when comparing two directories.
type Patch struct {
	Path string
	File *File
	Type PatchType
}
