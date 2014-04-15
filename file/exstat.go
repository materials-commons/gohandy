package file

import (
	"os"
	"time"
)

// ExFileInfo is an extended version of the os.FileInfo interface that
// includes additional information.
type ExFileInfo interface {
	os.FileInfo       // Support the os.FileInfo interface
	CTime() time.Time // Creation time
	ATime() time.Time // Last access time
	INode() uint64    // INode for systems that support it, otherwise 0
}

// ExStat is an extended version of the os.Stat() method.
func ExStat(path string) (fileInfo ExFileInfo, err error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	exfi := newExFileInfo(fi)
	return exfi, nil
}
