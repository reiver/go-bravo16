package bravo16_test

import (
	"github.com/reiver/go-bravo16"

	"fmt"
)

func ExampleDecode() {

	// 0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf
	var src []byte = []byte("00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf")

	lenDst := bravo16.DecodeLen(len(src))
	var dst []byte = make([]byte, lenDst)

	_, err := bravo16.Decode(dst, src)

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("SRC: %s (bravo16)\n", src)
	fmt.Printf("DST: %x (hexadecimal)\n", dst)

	// Output:
	// SRC: 00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf (bravo16)
	// DST: 000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f (hexadecimal)
}
