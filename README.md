# multipart creator

Utility to create the mulitpart content that could be used as the body of an HTTP POST request or email.

### Usage

This package is fully go-getable. So just type `go get github.com/rosbit/multipart-creator` to install.

```go
package main

import (
	"github.com/rosbit/multipart-creator"
	"fmt"
	"os"
	"bytes"
)

func main() {
	// sample 1:
	params := []multipart.Param{
		multipart.Param{"name", "rosbit", nil},
		multipart.Param{"age", 10, nil},
		multipart.Param{"file", "this/is/filename", bytes.NewBuffer([]byte("the content of filename"))},
	}

	contentType, err := multipart.Create(os.Stdout, "", params)

	// sample 2:
	contentType, err := multipart.CreateMultiPart(os.Stdout, "-- this is my boundary --",
		multipart.KeyVal("name", "rosbit"),
		multipart.KeyVal("age", 10),
		multipart.Reader("file", "file/name", bytes.NewBuffer([]byte("the content of optional file name"))),
	)

	// output result
	if err != nil {
		fmt.Printf("failed to create multipart: %v\n", err)
		return
	}
	fmt.Printf("Content-Type: %s\n", contentType)
}
```

### Status

The package is fully tested, so be happy to use it.

### Contribution

Pull requests are welcome! Also, if you want to discuss something send a pull request with proposal and changes.
__Convention:__ fork the repository and make changes on your fork in a feature branch.
