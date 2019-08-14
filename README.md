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

## Safety

There are letters, and numerals that individual can confuse:

|   Character   | Confused With |
| ------------- |:------------- |
|         **0** | 8 (eight)     |
|         **0** | D (uppercase) |
|         **0** | O (uppercase) |
|         **0** | o (lowercase) |

|   Character   | Confused With |
| ------------- |:------------- |
|         **1** | 7 (seven)     |
|         **1** | I (uppercase) |
|         **1** | i (lowercase) |
|         **1** | j (lowercase) |
|         **1** | L (uppercase) |
|         **1** | l (lowercase) |

|   Character   | Confused With |
| ------------- |:------------- |
|         **2** | Z (uppercase) |
|         **2** | z (uppercase) |

|   Character   | Confused With |
| ------------- |:------------- |
|         **3** | 5 (five)      |
|         **3** | 8 (eight)     |
|         **3** | 9 (nine)      |

|   Character   | Confused With |
| ------------- |:------------- |
|         **4** | 9 (nine)      |

|   Character   | Confused With |
| ------------- |:------------- |
|         **5** | 3 (three)     |
|         **5** | 8 (eight)     |
|         **5** | S (uppercase) |
|         **5** | s (lowercase) |

|   Character   | Confused With |
| ------------- |:------------- |
|         **6** | 8 (eight)     |
|         **6** | b (lowercase) |
|         **6** | G (uppercase) |

|   Character   | Confused With |
| ------------- |:------------- |
|         **7** | 1 (one)       |
|         **7** | F (uppercase) |
|         **7** | T (uppercase) |
|         **7** | Z (uppercase) |
|         **7** | z (lowercase) |

|   Character   | Confused With |
| ------------- |:------------- |
|         **8** | 0 (zero)      |
|         **8** | 3 (three)     |
|         **8** | 5 (five)      |
|         **8** | 6 (six)       |
|         **8** | B (uppercase) |
|         **8** | S (uppercase) |
|         **8** | s (uppercase) |

|   Character   | Confused With |
| ------------- |:------------- |
|         **9** | 3 (three)     |
|         **9** | 4 (four)      |
|         **9** | g (lowercase) |
|         **9** | q (lowercase) |
