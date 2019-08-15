package bravo16

import (
	"testing"
)

func TestInternalInvalidAsError(t *testing.T) {

	var err error = internalInvalid{} // THIS IS WHAT ACTUALLY MATTERS!


	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalInvalidAsInvalid(t *testing.T) {

	var complainer Invalid = internalInvalid{} // THIS IS WHAT ACTUALLY MATTERS!


	if nil == complainer {
		t.Error("This should never happen.")
	}
}
