package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

// walkerContext keeps contextual information around as a directory tree is walked.
// This allows the directory walk function to add items to the correct directory
// object.
type walkerContext struct {
	baseDir string
	current *Directory
	all     map[string]*Directory
}

// Load walks a given directory path creating a Directory object with all files and
// sub directories filled out.
func Load(path string) (*Directory, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	dir := &Directory{
		Info: FileInfo{
			Path:  path,
			IsDir: true,
			MTime: fi.ModTime(),
			Mode:  fi.Mode(),
		},
	}
	context := &walkerContext{
		current: dir,
		baseDir: path,
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
	if path != c.current.Info.Path && dirpath != c.current.Info.Path {
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
		if c.baseDir != path {
			c.addSubdir(path, info)
		}
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
	f := newFileInfo(path, info)
	c.current.Files = append(c.current.Files, f)
}

// newDirectory creates a new Directory entry.
func newDirectory(path string, info os.FileInfo) *Directory {
	return &Directory{
		Info: FileInfo{
			Path:  path,
			Mode:  info.Mode(),
			MTime: info.ModTime(),
			IsDir: true,
		},
	}
}

// newFile creates a new File entry.
func newFileInfo(path string, info os.FileInfo) *FileInfo {
	return &FileInfo{
		Path:  path,
		Mode:  info.Mode(),
		MTime: info.ModTime(),
		Size:  info.Size(),
	}
}
