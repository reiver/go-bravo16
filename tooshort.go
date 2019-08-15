package bravo16

import (
	"fmt"
)

// TooShort is an error that is returned by bravo16.Decode() is it receives an invalid symbol.
type TooShort interface {
	error
	TooShort() (expected int, actual int, src int)
}

func errTooShort(expected int, actual int, src int) error {
	var e TooShort = &internalTooShort{
		expected: expected,
		actual:   actual,
		src:      src,
	}

	return e
}

type internalTooShort struct {
	expected int
	actual   int
	src      int
}

func (receiver internalTooShort) Error() string {
	return fmt.Sprintf("bravo16: Destination Too Short: for source of length %d bytes, expected destination length to be at least %d bytes, actually was %d bytes long", receiver.src, receiver.expected, receiver.actual)
}

func (receiver internalTooShort) TooShort() (expected int, actual int, src int) {
	return receiver.expected, receiver.actual, receiver.src
}
