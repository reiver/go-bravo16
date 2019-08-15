/*
Package bravo16 provides a more (human) safe base-16 binary-to-text encoding, and decoding; brave16 is an alternative to hexadecimal.

The safety comes from carefully choosing the symbols use for the base-16 symbols.

Rather than using hexadecimal's symbols:

	0 1 2 3 4 5 6 7 8 9 a b c d e f

... bravo16 uses the symbols:

	0 1 2 m 4 5 x V H k A b C d E f

Which means that the symbols for ‘3’, ‘6’, ‘7’, ‘8’, and ‘9’ differ between bravo16 and hexadecimal.

And also that the symbols for ‘A’, ‘C’, and ‘E’ are always in uppercase.
And the symbols for ‘b’, ‘d’, and ‘e’ are always in lowercase.

Literals

To make it so that bravo16 literals can be distinguished from hexadecimal literals, decimal literlals, and octal liberals, bravo16 literals begin with the following prefix:

	0r

I.e., bravo16 literals begin with a zero followed by a lowercase ‘r’.

Here are some examples of bravo16 literals:

	0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf
	0r00
	0r11
	0r22
	0rmm
	0r44
	0r55
	0rxx
	0rVV
	0rHH
	0rkk
	0rAA
	0rbb
	0rCC
	0rdd
	0rEE
	0rff

Example Encode

Here is how you encode binary data as bravo16:

	// The ‘src’ variable contains our binary data.
	var src []byte = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, 0x9c, 0x08, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1, 0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f}
	
	// ...
	
	// The ‘dst’ variable will contain the encoded bravo16 data, after we call bravo16.Encode()
	lenDst := bravo16.EncodeLen(len(src))
	var dst []byte = make([]byte, lenDst)
	
	n, err := bravo16.Encode(dst, src)

Example Encode Literal

Here is how you encode binary data as a bravo16 literal:

	// The ‘src’ variable contains our binary data.
	var src []byte = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, 0x9c, 0x08, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1, 0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f}
	
	// ...
	
	// The ‘dst’ variable will contain the encoded bravo16 data, after we call bravo16.Encode()
	lenDst := bravo16.EncodeLiteralLen(len(src))
	var dst []byte = make([]byte, lenDst)
	
	n, err := bravo16.EncodeLiteral(dst, src)

Example Encode

Here is how you decode bravo16 data:

	// The ‘src’ variable contains our binary data.
	var src []byte = []byte("00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf")
	
	// ...
	
	// The ‘dst’ variable will contain the encoded bravo16 data, after we call bravo16.Encode()
	lenDst := bravo16.DecodeLen(len(src))
	var dst []byte = make([]byte, lenDst)
	
	n, err := bravo16.Decode(dst, src)

Example Encode Literal

Here is how you decode a bravo16 literal:

	// The ‘src’ variable contains our binary data.
	var src []byte = []byte("0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf")
	
	// ...
	
	// The ‘dst’ variable will contain the encoded bravo16 data, after we call bravo16.Encode()
	lenDst := bravo16.DecodeLiteralLen(len(src))
	var dst []byte = make([]byte, lenDst)
	
	n, err := bravo16.DecodeLiteral(dst, src)
*/
package bravo16
