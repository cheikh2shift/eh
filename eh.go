package eh

import (
	"fmt"
	"log"
	"runtime"
)

// ripped from the Go
// src.
func callerPC(depth int) uintptr {
	var pcs [1]uintptr
	runtime.Callers(depth, pcs[:])
	return pcs[0]
}

// Err accepts an error and
// returns a variable with type `Error`.
// Use this method to get the location
// where an error is constructed/initialized
func Err(err error) Error {

	u := callerPC(3)
	fs := runtime.CallersFrames([]uintptr{u})
	f, _ := fs.Next()

	return Error{
		CodeReference: CodeReference{
			line:     f.Line,
			file:     f.File,
			function: f.Function,
		},
		err: err,
	}
}

func Log(err error) {

	// return early,
	// enable function to
	// be used as a utility.
	if err == nil {
		return
	}

	u := callerPC(3)
	fs := runtime.CallersFrames([]uintptr{u})
	f, _ := fs.Next()

	err = Error{
		CodeReference: CodeReference{
			line:     f.Line,
			file:     f.File,
			function: f.Function,
		},
		err: err,
	}

	log.Println(err.Error())
}

// CodeReference is embedded
// within type error. It stores
// the line, file and function
// that invoked the `Err` function
type CodeReference struct {
	line     int
	file     string
	function string
}

func (c CodeReference) GetLine() int {
	return c.line
}

func (c CodeReference) GetFile() string {
	return c.file
}

func (c CodeReference) GetFunction() string {
	return c.function
}

type Error struct {
	CodeReference
	err error
}

func (f Error) GetError() error {
	return f.err
}

// Error helps the `Error` type
// satisfy Go's error interface.
func (f Error) Error() string {
	return fmt.Sprintf(
		"error : %s [%s:%v [%s]]",
		f.err.Error(),
		f.file,
		f.line,
		f.function,
	)
}
