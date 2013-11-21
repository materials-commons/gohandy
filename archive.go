package gohandy

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

func Unpack(r *tar.Reader, toPath string) error {
	for {
		hdr, err := r.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		err = nil

		if hdr.Typeflag == tar.TypeDir {
			dirpath := filepath.Join(toPath, hdr.Name)
			err = os.MkdirAll(dirpath, 0777)
		} else if hdr.Typeflag == tar.TypeReg || hdr.Typeflag == tar.TypeRegA {
			err = writeFile(toPath, hdr.Name, r)
		}

		if err != nil {
			return nil
		}

	}

	return nil
}

func writeFile(toPath, filename string, r *tar.Reader) error {
	path := filepath.Join(toPath, filename)
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, r); err != nil {
		return err
	}

	return nil
}
