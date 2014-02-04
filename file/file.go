package file

import (
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func NormalizePath(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

func Checksum32(path string) uint32 {
	file, _ := os.Open(path)
	defer file.Close()
	c := crc32.NewIEEE()
	bytes, _ := ioutil.ReadAll(file)
	withcrc := c.Sum(bytes)
	return crc32.ChecksumIEEE(withcrc)
}

func Hash(hasher hash.Hash, path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := io.Copy(hasher, f); err != nil {
		return nil, err
	}
	return hasher.Sum(nil), nil
}

func HashStr(hasher hash.Hash, path string) (string, error) {
	if csum, err := Hash(hasher, path); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%x", csum), nil
	}
}

func IsDir(path string) bool {
	finfo, err := os.Stat(path)
	switch {
	case err != nil:
		return false
	case finfo.IsDir():
		return true
	default:
		return false
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Copy(src, dest string) error {
	fsrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fsrc.Close()

	fdest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer fdest.Close()

	if _, err := io.Copy(fdest, fsrc); err != nil {
		return err
	}

	return nil
}
