package bravo16_test

import (
	"github.com/reiver/go-bravo16"

	"fmt"
)

func ExampleEncode() {

	// 0x000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f
	var src []byte = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, 0x9c, 0x08, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1, 0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f}

	lenDst := bravo16.EncodeLen(len(src))
	var dst []byte = make([]byte, lenDst)

	_, err := bravo16.Encode(dst, src)

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Println("SRC: 000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f (hexadecimal)",)
	fmt.Printf("DST: %s (bravo16)\n", dst)

	// Output:
	// SRC: 000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f (hexadecimal)
	// DST: 00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf (bravo16)
}
