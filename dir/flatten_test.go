package dir

import (
	"fmt"
	"testing"
)

var _ = fmt.Println

func TestFlatten(t *testing.T) {
	d, _ := Load("/tmp/testdir")
	files := d.Flatten()
	for _, f := range files {
		fmt.Printf("%#v\n", f)
	}
}
