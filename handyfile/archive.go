package futil

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

func Unpack(r *tar.Reader, toPath string) error {
	for {
		hdr, err := r.Next()
		switch {
		case err == io.EOF:
			break
		case err != nil:
			return err
		default:
			if err := doOnType(hdr.Typeflag, toPath, hdr.Name, r); err != nil {
				return err
			}
		}

	}

	return nil
}

func doOnType(typeFlag byte, toPath string, name string, r *tar.Reader) error {
	fullpath := filepath.Join(toPath, name)
	switch typeFlag {
	case tar.TypeReg, tar.TypeRegA:
		return writeFile(fullpath, r)
	case tar.TypeDir:
		return os.MkdirAll(fullpath, 0777)
	default:
		return nil
	}
}

func writeFile(path string, r *tar.Reader) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, r)
	return err
}
