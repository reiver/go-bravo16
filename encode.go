package bravo16

import (
	"fmt"
)

// Encode encodes the binary data in ‘src’ into ‘dst’ as bravo16, and returns the number of bytes written to ‘dst’.
func Encode(dest []byte, src []byte) (int64, error) {
	lenSrc := len(src)

	if 1 > lenSrc {
		return 0, nil
	}
	if expectedAtLeast, actual := 2*lenSrc, len(dest); actual < expectedAtLeast {
		return 0, fmt.Errorf("bravo16: Destination Too Short: expected length at least %d but actually got length %d", expectedAtLeast, actual)
	}

	for i, b := range src {
		dest[i*2]   = encode[b >> 4]   //  most significant byte
		dest[i*2+1] = encode[b & 0x0f] // least significant byte
	}

	return int64(lenSrc*2), nil
}
