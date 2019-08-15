package bravo16

import (
	"fmt"
)

// Invalid is an error that is returned by bravo16.Decode() is it receives an invalid symbol.
type Invalid interface {
	error
	Invalid() (byte, int64)
}

func errInvalid(character byte, index int64) error {
	var e Invalid = &internalInvalid{
		character: character,
		index: index,
	}

	return e
}

type internalInvalid struct {
	character byte
	index int64
}

func (receiver internalInvalid) Error() string {
	return fmt.Sprintf("bravo16: Invalid bravo16 Symbol: symbol=%q index=%d", receiver.character, receiver.index)
}

func (receiver internalInvalid) Invalid() (byte, int64) {
	return receiver.character, receiver.index
}
