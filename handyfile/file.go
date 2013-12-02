package handyfile

import (
	"hash/crc32"
	"io/ioutil"
	"os"
	"io"
)

func Checksum32(path string) uint32 {
	file, _ := os.Open(path)
	defer file.Close()
	c := crc32.NewIEEE()
	bytes, _ := ioutil.ReadAll(file)
	withcrc := c.Sum(bytes)
	return crc32.ChecksumIEEE(withcrc)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
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
