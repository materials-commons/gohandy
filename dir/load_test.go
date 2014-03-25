package dir

import (
	"fmt"
	"testing"
)

var _ = fmt.Println

func TestPopulate(t *testing.T) {
	d, err := Load("/tmp/testdir")
	if err != nil {
		t.Fatalf("Failed to create directory for /tmp/testdir: %s", err)
	}

	fmt.Printf("%#v\n", d)
	fmt.Println("")
	for _, dir := range d.SubDirectories {
		fmt.Printf("Subdir = %#v\n\n", dir)
		for _, d2 := range dir.SubDirectories {
			fmt.Printf("  Subdir for %s: %#v\n", dir.Path, d2)
		}
		fmt.Println("")
		for _, f2 := range dir.Files {
			fmt.Printf("  File for %s: %#v\n", dir.Path, f2)
		}
	}

	fmt.Println("")
	for _, f := range d.Files {
		fmt.Printf("File = %#v\n", f)
	}
}
