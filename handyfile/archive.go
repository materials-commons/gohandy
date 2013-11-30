package handyfile

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

type TarReader struct {
	file *os.File
	gz   *gzip.Reader
	tr   *tar.Reader
}

func NewTar(path string) (*TarReader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	tr := tar.NewReader(file)
	return &TarReader{file: file, gz: nil, tr: tr}, nil
}

func NewTarGz(path string) (*TarReader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	gz, _ := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	tr := tar.NewReader(gz)
	return &TarReader{file: file, gz: gz, tr: tr}, nil
}

func (tr *TarReader) Unpack(toPath string) error {
	defer tr.file.Close()
	if tr.gz != nil {
		defer tr.gz.Close()
	}

	r := tr.tr

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
