package file

import (
	"os"
	"syscall"
	"time"
)

// darwinExFileInfo stores windows specific file information.
// At the moment all the information we need is available
// through the Sys() interface.
type darwinExFileInfo struct {
	os.FileInfo
}

// timespecToTime converts a unix timespec into a time.Time. This was
// copied from os/stat_linux.go in the Go source.
func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

// CTime returns the creation time (ctime) from stat_t.
func (fi *darwinExFileInfo) CTime() time.Time {
	return timespecToTime(fi.Sys().(*syscall.Stat_t).Ctimespec)
}

// ATime returns the access time (atime) from stat_t
func (fi *darwinExFileInfo) ATime() time.Time {
	return timespecToTime(fi.Sys().(*syscall.Stat_t).Atimespec)
}

// INode returns the files inode.
func (fi *darwinExFileInfo) INode() uint64 {
	return fi.Sys().(*syscall.Stat_t).Ino
}

// newExFileInfo creates a new winExFileInfo from a os.FileInfo.
func newExFileInfo(fi os.FileInfo) *darwinExFileInfo {
	return &darwinExFileInfo{
		FileInfo: fi,
	}
}
