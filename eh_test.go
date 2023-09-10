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
