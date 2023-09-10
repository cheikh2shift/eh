package eh_test

import (
	"errors"
	"testing"

	"github.com/cheikh2shift/eh"
)

func TestErr(t *testing.T) {

	err1 := eh.Err(
		errors.New("Test err"),
	)

	if err1.GetLine() != 12 {
		t.Fatalf("Expecting line number %v got %v", 12, err1.GetLine())
	}

	err1 = eh.Err(
		errors.New("Test err"),
	)

	if err1.GetLine() != 20 {
		t.Fatalf("Expecting line number %v got %v", 20, err1.GetLine())
	}

}

func TestErrUseCase(t *testing.T) {

	var err2, err3, err4 error

	if err2 != nil {
		t.Log(eh.Err(err2))
	}

	if err3 != nil {
		t.Log(eh.Err(err3))
	}

	err1 := eh.Err(
		errors.New("Test err"),
	)

	if err1.GetLine() != 42 {
		t.Fatalf("Expecting line number %v got %v", 42, err1.GetLine())
	}

	if err4 != nil {
		t.Log(eh.Err(err4))
	}

}
