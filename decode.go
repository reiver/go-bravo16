package bravo16

import (
	"fmt"
)

// Decode decodes the bravo16 encoded data in ‘src’ into ‘dst’ (as binary), and returns the number of bytes written to ‘dst’.
func Decode(dst []byte, src []byte) (int64, error) {
	lenSrc := len(src)

	if 1 > lenSrc {
		return 0, nil
	}
	if length := lenSrc; 0 != length % 2 {
		return 0, fmt.Errorf("bravo16: Source Odd Length: %d", length)
	}
	if expectedAtLeast, actual := DecodeLen(lenSrc), len(dst); actual < expectedAtLeast {
		return 0, errTooShort(expectedAtLeast, actual, lenSrc)
	}

	var n64 int64
	{
		limit := DecodeLen(lenSrc)

		for i:=0; i<limit; i++ {
			ii := i*2

			mostAt  := ii
			leastAt := ii+1

			mostCharacter  := src[mostAt]
			leastCharacter := src[leastAt]

			var mostNibble byte
			{
				var found bool
				mostNibble, found = decode(mostCharacter)
				if !found {
					return n64, errInvalid(mostCharacter, int64(mostAt))
				}
			}

			var leastNibble byte
			{
				var found bool
				leastNibble, found = decode(leastCharacter)
				if !found {
					return n64, errInvalid(leastCharacter, int64(leastAt))
				}
			}

			dst[i] = (mostNibble<<4) | leastNibble
			n64++
		}
	}

	return n64, nil
}

// DecodeLen returns the length in bytes of the bravo16 data decoded into binary data.
func DecodeLen(n int) int {
	return n/2
}

// DecodeLiteral is similar to Decode, except expects ‘src’ to start with a “0r” prefix.
func DecodeLiteral(dst []byte, src []byte) (int64, error) {
	lenSrc := len(src)

	if 1 > lenSrc {
		return 0, nil
	}
	if length := lenSrc; 0 != length % 2 {
		return 0, fmt.Errorf("bravo16: Source Odd Length: %d", length)
	}
	if expectedAtLeast, actual := DecodeLiteralLen(lenSrc), len(dst); actual < expectedAtLeast {
		return 0, errTooShort(expectedAtLeast, actual, lenSrc)
	}

	if '0' != src[0] || 'r' != src[1] {
		return 0, fmt.Errorf("bravo16: Not bravo16 Literal: does not begin with “0r”, instead begins with %q.", src[:2])
	}

	return Decode(dst, src[2:])
}

// DecodeLiteralLen returns the length in bytes of the bravo16 literal decoded into binary data.
func DecodeLiteralLen(n int) int {
	n -= 2

	if 0 > n {
		return 0
	}

	return DecodeLen(n)
}
