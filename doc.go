/*
Package bravo16 provides (human) safe base-16 binary-to-text encoding, and decoding; brave16 is an alternative to hexadecimal.

The safety comes from carefully choosing the symbols use for the base-16 symbols.

Rather than using hexadecimal's symbols:

	0 1 2 3 4 5 6 7 8 9 a b c d e f

... bravo16 uses the symbols:

	0 1 2 m 4 5 x V H k A b C d E f

Which means that the symbols for ‘3’, ‘6’, ‘7’, ‘8’, and ‘9’ differ between bravo16 and hexadecimal.

And also that the symbols for ‘A’, ‘C’, and ‘E’ are always in uppercase.
And the symbols for ‘b’, ‘d’, and ‘e’ are always in lowercase.

Literals

To make it so that **bravo16** literals can be distinguished from _hexadecimal_ literals, _decimal_ literlals, and _octal_ liberals, **bravo16** literals begin with the following prefix:

	0r

I.e., **bravo16** literals begin with a **zero** followed by a lowercase ‘r’.

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

*/
package bravo16
