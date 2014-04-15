package file

import (
	"os"
	"time"
)

// winExFileInfo stores windows specific file information.
// At the moment all the information we need is available
// through the Sys() interface.
type winExFileInfo struct {
	os.FileInfo
}

// CTime returns the CreationTime from Win32FileAttributeData.
func (fi *winExFileInfo) CTime() time.Time {
	return time.Unix(0, fi.Sys().(syscall.Win32FileAttributeData).CreationTime)
}

// ATime returns the LastAccessTime from Win32FileAttributeData.
func (fi *winExFileInfo) ATime() time.Time {
	return time.Unix(0, fi.Sys().(syscall.Win32FileAttributeData).LastAccessTime)
}

// INode Windows doesn't directly support inodes. At the moment this just returns
// 0. Since windows files don't change their underlying file id when they are
// renamed, they can't be used to identify new files.
func (fi *winExFileInfo) INode() uint64 {
	return 0
}

// newExFileInfo creates a new winExFileInfo from a os.FileInfo.
func newExFileInfo(fi os.FileInfo) *winExFileInfo {
	return &winExFileInfo{
		FileInfo: fi,
	}
}
