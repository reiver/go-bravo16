package bravo16_test

import (
	"github.com/reiver/go-bravo16"

	"fmt"
)

func ExampleInvalid() {

	// This is not valid bravo16
	var src []byte = []byte("AA3678")

	lenDst := bravo16.DecodeLen(len(src))
	var dst []byte = make([]byte, lenDst)

	_, err := bravo16.Decode(dst, src)

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)

		switch casted := err.(type) {
		case bravo16.Invalid: // <-------------- We are detecting a bravo16.Invalid error here.
			symbol, index := casted.Invalid()

			fmt.Printf("Detected a bravo16.Invalid error for symbol=%q, and index=%d.\n", symbol, index)
		default:
			fmt.Printf("Unhandled error: (%T)\n", err)
		}

		return
	}

	fmt.Printf("SRC: 0r%s (bravo16)\n", src)
	fmt.Printf("DST: 0x%x (hexadecimal)\n", dst)
	fmt.Println("      .")

	// Output:
	// ERROR: bravo16: Invalid bravo16 Symbol: symbol='3' index=2
	// Detected a bravo16.Invalid error for symbol='3', and index=2.
}
