# Connect package to project

First, you need to have Go installed. If Go hasn't installed, then go to [their website](https://go.dev/) and install it.
When you installed Go, create directory and init module:
```
go mod init example.com/example
```

Now, after this, you can connect package:
```
go get github.com/aleksnew2/fsutilss
```

# Example code

```go
package main

import (
	"fmt"
	"time"

	fsutilss "github.com/aleksnew2/fsutilss"
)

func main() {
	// Creating file
	file, nerr := fsutilss.CreateFile("hi.txt", "HI!", "BYE!", "YAY")
	if nerr != nil {
		fmt.Println(nerr)
	}

	file.Output()

	// Waiting for five seconds
	time.Sleep(5 * time.Second)

	// Removing file
	serr := fsutilss.RemoveFile("hi.txt")
	if serr != nil {
		fmt.Println(serr)
	}
}
```
