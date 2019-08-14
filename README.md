# go-bravo16

Package **bravo16** provides **safe** base-16 binary-to-text encoding, and decoding;
**brave16** is an alternative to hexadecimal.

The **safety** comes from carefully choosing the symbols use for the base-16 symbols.

Rather than using hexadecimal's symbols:
```
0 1 2 3 4 5 6 7 8 9 a b c d e f
```
... **bravo16** uses the symbols:
```
0 1 2 3 4 5 6 7 k 9 a m c x w p
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-bravo16

[![GoDoc](https://godoc.org/github.com/reiver/go-bravo16?status.svg)](https://godoc.org/github.com/reiver/go-bravo16)

## Symbols

The following symbols are used for the **bravo16** encoding:

|  bravo16  | hexadecimal | decimal | binary |
| ---------:| -----------:| -------:| ------:|
|       0   |           0 |       0 |   0000 |
|       1   |           1 |       1 |   0001 |
|       2   |           2 |       2 |   0010 |
|       3   |           3 |       3 |   0011 |
|       4   |           4 |       4 |   0100 |
|       5   |           5 |       5 |   0101 |
|       6   |           6 |       6 |   0110 |
|       7   |           7 |       7 |   0111 |
|     **k** |           8 |       8 |   1000 |
|       9   |           9 |       9 |   1001 |
|       a   |           a |      10 |   1010 |
|     **m** |           b |      11 |   1011 |
|       c   |           c |      12 |   1100 |
|     **x** |           d |      13 |   1101 |
|     **w** |           e |      14 |   1110 |
|     **p** |           f |      15 |   1111 |

Which means that the symbols for 8, 11, 13, 14, and 15 differ between **bravo16** and _hexadecimal_.
