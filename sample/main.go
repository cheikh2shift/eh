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

	logError(errors.New("Error 5"))

	return err2
}

func logError(err error) {
	log.Println(
		//
		eh.Err(err, 4),
	)
}
