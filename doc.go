/*
Package bravo16 provides safe base-16 binary-to-text encoding, and decoding; brave16 is an alternative to hexadecimal.

The safety comes from carefully choosing the symbols use for the base-16 symbols.

Rather than using hexadecimal's symbols:

0 1 2 3 4 5 6 7 8 9 a b c d e f

... bravo16 uses the symbols:

0 1 2 3 4 5 6 7 k 9 a m c x w p
*/
package bravo16
