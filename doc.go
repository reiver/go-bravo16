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
*/
package bravo16
