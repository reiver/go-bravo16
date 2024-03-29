package bravo16

// Encode encodes the binary data in ‘src’ into ‘dst’ as bravo16, and returns the number of bytes written to ‘dst’.
func Encode(dst []byte, src []byte) (int64, error) {
	lenSrc := len(src)

	if 1 > lenSrc {
		return 0, nil
	}
	if expectedAtLeast, actual := EncodeLen(lenSrc), len(dst); actual < expectedAtLeast {
		return 0, errTooShort(expectedAtLeast, actual, lenSrc)
	}

	for i, b := range src {
		dst[i*2]   = encode[b >> 4]   //  most significant byte
		dst[i*2+1] = encode[b & 0x0f] // least significant byte
	}

	return int64(EncodeLen(lenSrc)), nil
}

// EncodeLen returns the length in bytes of the binary data encoded in bravo16.
func EncodeLen(n int) int {
	if 0 > n {
		return 0
	}

	return 2*n
}

// EncodeLiteral is similar to Encode, except it prefixes a “0r”.
func EncodeLiteral(dst []byte, src []byte) (int64, error) {
	lenSrc := len(src)

	if 1 > lenSrc {
		return 0, nil
	}
	if expectedAtLeast, actual := 2*lenSrc+2, len(dst); actual < expectedAtLeast {
		return 0, errTooShort(expectedAtLeast, actual, lenSrc)
	}

	dst[0] = '0'
	dst[1] = 'r'

	n64, err := Encode(dst[2:], src)
	n64 += 2
	return n64, err
}

// EncodeLiteralLen returns the length in bytes of the binary data encoded in bravo16 plus the length of the  bravo16 prefix (i.e., “0r”).
func EncodeLiteralLen(n int) int {
	if 1 > n {
		return 0
	}

	return 2+EncodeLen(n)
}
