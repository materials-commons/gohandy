package dir

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// File is a file entry in one of the directories
type File struct {
	Path     string      // Fullname including path
	Size     int64       // Size of file
	Checksum string      // MD5 Hash - optional
	MTime    time.Time   // Modification Time
	Mode     os.FileMode // Permissions
}

// Directory is a container for the files and sub directories in a single directory.
// Each sub directory will itself contain the same list.
type Directory struct {
	Path           string       // Full path
	Mode           os.FileMode  // Permissions
	MTime          time.Time    // Modification Time
	Files          []*File      // List of files in this directory
	SubDirectories []*Directory // List of directories in this directory
}

// walkerContext keeps contextual information around as a directory tree is walked.
// This allows the directory walk function to add items to the correct directory
// object.
type walkerContext struct {
	current *Directory
	all     map[string]*Directory
}

// Populate walks a given directory path creating a Directory object with all files and
// sub directories filled out.
func Load(path string) (*Directory, error) {
	dir := &Directory{
		Path: path,
	}
	context := &walkerContext{
		current: dir,
		all:     make(map[string]*Directory),
	}
	context.all[path] = dir
	if err := filepath.Walk(path, context.directoryWalker); err != nil {
		return dir, err
	}
	return dir, nil
}

// directoryWalker is the function called by filepath.Walk as it walks a directory tree.
// It builds up the list of entries, both files and directories, in each directory that
// is visited populating a top level Directory.
func (c *walkerContext) directoryWalker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	dirpath := filepath.Dir(path)
	if path != c.current.Path && dirpath != c.current.Path {
		// We have descended into a new directory, so set current to the
		// new entry.
		dir, ok := c.all[dirpath]
		if !ok {
			panic(fmt.Sprintf("Fatal: Could not find directory: %s", dirpath))
		}
		c.current = dir
	}

	switch {
	case info.IsDir():
		c.addSubdir(path, info)
	default:
		c.addFile(path, info)
	}

	return nil
}

// addSubdir adds a new sub directory to the current directory.
func (c *walkerContext) addSubdir(path string, info os.FileInfo) {
	dir := newDirectory(path, info)
	c.current.SubDirectories = append(c.current.SubDirectories, dir)
	c.all[path] = dir
}

// addFile adds a new file to the current directory
func (c *walkerContext) addFile(path string, info os.FileInfo) {
	f := newFile(path, info)
	c.current.Files = append(c.current.Files, f)
}

// newDirectory creates a new Directory entry.
func newDirectory(path string, info os.FileInfo) *Directory {
	return &Directory{
		Path:  path,
		Mode:  info.Mode(),
		MTime: info.ModTime(),
	}
}

// newFile creates a new File entry.
func newFile(path string, info os.FileInfo) *File {
	return &File{
		Path:  path,
		Mode:  info.Mode(),
		MTime: info.ModTime(),
		Size:  info.Size(),
	}
}
