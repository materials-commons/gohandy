package dir

import (
	"os"
	"time"
)

// FileInfo describes a file or directory entry
type FileInfo struct {
	Path     string      // Full path including name
	Size     int64       // Size valid only for file
	Checksum string      // MD5 Hash - valid only for file
	MTime    time.Time   // Modification time
	Mode     os.FileMode // Permissions
	IsDir    bool        // True if this entry represents a directory
}

// Directory is a container for the files and sub directories in a single directory.
// Each sub directory will itself contain the same list.
type Directory struct {
	Info           FileInfo
	Files          []*FileInfo  // List of files in this directory
	SubDirectories []*Directory // List of directories in this directory
}
