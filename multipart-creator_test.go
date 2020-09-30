package multipart

import (
	"testing"
	"bytes"
	"os"
	"fmt"
)

var (
	params = []Param{
		Param{"name", "rosbit", nil},
		Param{"age", 10, nil},
		Param{"file", "this/is/filename.png", bytes.NewBuffer([]byte("the content of filename"))},
	}
)

func Test_1_with_generated_boundary(t *testing.T) {
	fmt.Printf("////////////////////////////////////////////////////////////\n")
	fmt.Printf("///// BEGIN to create multipart with generated boundary/////\n")
	contentType, err := Create(os.Stdout, "", params)
	if err != nil {
		t.Fatalf("failed to create multipart: %v\n", err)
	}

	fmt.Printf("///// END to create multipart with generated boundary/////\n")
	fmt.Printf("//////////////////////////////////////////////////////////\n")
	fmt.Printf("Content-Type: %s\n\n\n", contentType)
}

func Test_2_with_given_boundary(t *testing.T) {
	fmt.Printf("////////////////////////////////////////////////////////\n")
	fmt.Printf("///// BEGIN to create multipart with given boundary/////\n")
	params[2].Reader = bytes.NewBuffer([]byte("a new content of filename"))
	contentType, err := Create(os.Stdout, "-----this is my boundary -----", params)
	if err != nil {
		t.Fatalf("failed to create multipart: %v\n", err)
	}
	fmt.Printf("///// END to create multipart with given boundary/////\n")
	fmt.Printf("//////////////////////////////////////////////////////\n")
	fmt.Printf("Content-Type: %s\n\n\n", contentType)
}
