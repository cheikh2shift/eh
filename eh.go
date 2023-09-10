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

func (f Error) Error() string {
	return fmt.Sprintf(
		"error : %s [%s:%v [%s]]",
		f.err.Error(),
		f.file,
		f.line,
		f.function,
	)
}
