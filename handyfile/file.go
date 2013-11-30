package handyfile

import (
	"hash/crc32"
	"io/ioutil"
	"os"
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
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
