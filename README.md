# eh
Error handling library. regenerates error message alongside the line of code that caused it.

Think of it as adding more information to an error.

### Installation

To add `eh` to your project, run the following command:

```
go get github.com/cheikh2shift/eh
```

### Usage
Here is a code sample of how to use the project:

main.go
```
package main

import (
	"errors"
	"log"

	"github.com/cheikh2shift/eh"
)

func main() {

	if err := run(); err != nil {
		log.Println(
			err,
		)
	}
}

func run() error {
	err1 := errors.New("Error 1")

	log.Println(
		eh.Err(err1),
	)

	eh.Log(err1)

	err2 := eh.Err(errors.New("Error 2"))

	return err2
}
```


output
```
$ go run main.go 
2023/09/10 17:16:03 error : Error 1 [~/eh/sample/main.go:23 [main.run]]
2023/09/10 17:16:03 error : Error 1 [~/eh/sample/main.go:26 [main.run]]
2023/09/10 17:16:03 error : Error 2 [~/eh/sample/main.go:28 [main.run]]
```

### How it works?
After discovering slog's `AddSource` flag, I decided to look around the code base to determine how they were doing it. You can determine how this is done by looking at the source code in this repository; the idea is simple, the line number where you call the helper functions will be determined as the origin of the call. 

Please note that placing this within an abstraction will not work as intended; the line number within the abstracted function will be called each time. I plan on solving this in the future by making the depth level configurable.


 I plan on using this package for future projects because there are times where my functions have multiple errors and I need a way to determine which line is setting the error. 

Check out the documentation [here](https://pkg.go.dev/github.com/cheikh2shift/eh)