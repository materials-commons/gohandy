package dir

// PatchType denotes the kind of patch operation
type PatchType int

const (
	// PatchCreate created item
	PatchCreate PatchType = iota

	// PatchDelete deleted item
	PatchDelete
)

// PatchApplyTo tells us which directory set to apply the patch to.
type PatchApplyTo int

const (
	// OriginalDirectory apply the patch to the original directory set.
	OriginalDirectory PatchApplyTo = iota

	// NewDirectory apply the patch to the new directory set.
	NewDirectory
)

// Patch is an instance of a difference when comparing two directories. It specifies
// the kind of change to apply and where to apply it to.
type Patch struct {
	File    *FileInfo
	Type    PatchType
	ApplyTo PatchApplyTo
}
