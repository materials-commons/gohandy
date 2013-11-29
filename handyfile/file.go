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
