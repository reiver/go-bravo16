package bravo16

import (
	"testing"
)

func TestInternalTooShortAsError(t *testing.T) {

	var err error = internalTooShort{} // THIS IS WHAT ACTUALLY MATTERS!


	if nil == err {
		t.Error("This should never happen.")
	}
}

func TestInternalTooShortAsTooShort(t *testing.T) {

	var complainer TooShort = internalTooShort{} // THIS IS WHAT ACTUALLY MATTERS!


	if nil == complainer {
		t.Error("This should never happen.")
	}
}
