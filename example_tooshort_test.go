package bravo16_test

import (
	"github.com/reiver/go-bravo16"

	"fmt"
)

func ExampleTooShort() {

	var src []byte = []byte("0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf")

	lenDst := 4 // <--------- Normally we would call bravo16.DecodeLiteralLen() here, but we wanted bravo16.DecodeLiteral() to return a bravo16.TooShort error, so we manually set it.
	var dst []byte = make([]byte, lenDst)

	_, err := bravo16.DecodeLiteral(dst, src)

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)

		switch casted := err.(type) {
		case bravo16.TooShort: // <-------------- We are detecting a bravo16.TooShort error here.
			expected, actual, lenSrc := casted.TooShort()

			fmt.Printf("Detected a bravo16.TooShort error expected=%d actual=%d lenSrc=%d\n", expected, actual, lenSrc)
		default:
			fmt.Printf("Unhandled error: (%T)\n", err)
		}

		return
	}

	fmt.Printf("SRC: 0r%s (bravo16)\n", src)
	fmt.Printf("DST: 0x%x (hexadecimal)\n", dst)
	fmt.Println("      .")

	// Output:
	// ERROR: bravo16: Destination Too Short: for source of length 66 bytes, expected destination length to be at least 32 bytes, but actually was 4 bytes long
	// Detected a bravo16.TooShort error expected=32 actual=4 lenSrc=66
}
