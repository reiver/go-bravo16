# go-bravo16

Package **bravo16** provides (human) **safe** base-16 binary-to-text encoding, and decoding;
**brave16** is an alternative to hexadecimal.

The **safety** comes from carefully choosing the symbols used for the base-16 symbols.

Rather than using hexadecimal's symbols:
```
0 1 2 3 4 5 6 7 8 9 a b c d e f
```
... **bravo16** uses the symbols:
```
0 1 2 3 4 5 6 7 k 9 a m c x w p
```

(Why these symbols were chosen is covered later in this document.)

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

## Safety: Alphanumeric Confusion

There are a number of ways that _safety_ was considered when choosing the symbols used for the **brave16** (base-16) encoding.

One of those ways was:...
> Which symbols do humans commonly confuse with other symbols?

There are a number of letter, and numeral symbols that are common for individuals to confuse.

This confusion can be due to different reasons, namely:

* **the font used** (on a computer, or in print) can make certain symbols look similar and cause humans to confuse them,
* the way that some humans **(hand) write** certain symbols can make them look similar to other symbols and cause humans to confuse them, and
* the way that some humans **handwrite (in cursive)** certain symbols an make them look similar to other symbols and cause humans to confuse them.

Here are a number of tables listing out the common symbol confusions:....

### Zero (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **0** | 8 (eight)     |
|         **0** | D (uppercase) |
|         **0** | O (uppercase) |
|         **0** | o (lowercase) |
|         **0** | Q (uppercase) |
|         **0** | U (uppercase) |

### One (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **1** | 7 (seven)     |
|         **1** | I (uppercase) |
|         **1** | i (lowercase) |
|         **1** | j (lowercase) |
|         **1** | L (uppercase) |
|         **1** | l (lowercase) |

### Two (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **2** | Z (uppercase) |
|         **2** | z (uppercase) |

### Three (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **3** | 5 (five)      |
|         **3** | 8 (eight)     |
|         **3** | 9 (nine)      |

### Four (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **4** | 9 (nine)      |
|         **4** | U (uppercase) |
|         **4** | u (lowercase) |

### Five (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **5** | 3 (three)     |
|         **5** | 8 (eight)     |
|         **5** | S (uppercase) |
|         **5** | s (lowercase) |
|         **5** | Y (uppercase) |
|         **5** | y (lowercase) |

### Six (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **6** | 8 (eight)     |
|         **6** | b (lowercase) |
|         **6** | G (uppercase) |

### Seven (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **7** | 1 (one)       |
|         **7** | F (uppercase) |
|         **7** | T (uppercase) |
|         **7** | Z (uppercase) |
|         **7** | z (lowercase) |

### Eight (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **8** | 0 (zero)      |
|         **8** | 3 (three)     |
|         **8** | 5 (five)      |
|         **8** | 6 (six)       |
|         **8** | B (uppercase) |
|         **8** | S (uppercase) |
|         **8** | s (uppercase) |

### Nine (Alphanumeric Confusion)

|   Character   | Confused With |
| ------------- |:------------- |
|         **9** | 3 (three)     |
|         **9** | 4 (four)      |
|         **9** | g (lowercase) |
|         **9** | q (lowercase) |

### A (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **a** (lowercase) | o (lowercase) |
| **a** (lowercase) | q (lowercase) |

**No common confusion for uppercase “A”.**

### B (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **B** (uppercase) | 8 (eight)     |

|     Character     | Confused With |
|:----------------- |:------------- |
| **b** (lowercase) | 6 (six)       |


### C (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **C** (uppercase) | e (lowercase) | 
| **C** (uppercase) | G (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **c** (lowercase) | e (lowercase) |
| **c** (lowercase) | G (uppercase) |

### D (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **D** (uppercase) | 0 (zero)      |
| **D** (uppercase) | O (uppercase) |
| **D** (uppercase) | o (lowercase) |
| **D** (uppercase) | Q (uppercase) |


**No common confusion for lowercase “d”.**

### E (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **E** (uppercase) | F (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **e** (lowercase) | C (uppercase) |
| **e** (lowercase) | c (lowercase) |

### F (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **F** (uppercase) | E (uppercase) |
| **F** (uppercase) | R (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **f** (lowercase) | t (lowercase) |

### G (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **G** (uppercase) | 6 (six)       |
| **G** (uppercase) | C (uppercase) |
| **G** (uppercase) | c (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **g** (lowercase) | 9 (nine)      |
| **g** (lowercase) | q (lowercase) |

### H (Alphanumeric Confusion)

**No common confusion for uppercase “H”.**

**No common confusion for lowercase “h”.**

### I (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **I** (uppercase) | 1 (one)       |
| **I** (uppercase) | i (lowercase) |
| **I** (uppercase) | J (uppercase) |
| **I** (uppercase) | j (lowercase) |
| **I** (uppercase) | L (uppercase) |
| **I** (uppercase) | l (lowercase) |
| **I** (uppercase) | T (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **i** (lowercase) | 1 (one)       |
| **i** (lowercase) | I (uppercase) |
| **i** (lowercase) | J (uppercase) |
| **i** (lowercase) | j (lowercase) |
| **i** (lowercase) | l (lowercase) |

### J (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **J** (uppercase) | I (uppercase) |
| **J** (uppercase) | i (lowercase) |
| **J** (uppercase) | j (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **j** (lowercase) | 1 (one)       |
| **j** (lowercase) | I (uppercase) |
| **j** (lowercase) | i (lowercase) |
| **j** (lowercase) | J (uppercase) |

### K (Alphanumeric Confusion)

**No common confusion for uppercase “K”.**

**No common confusion for lowercase “k”.**

### L (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **L** (uppercase) | I (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **l** (lowercase) | 1 (one)       |
| **l** (lowercase) | I (uppercase) |
| **l** (lowercase) | i (lowercase) |

### M (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **M** (uppercase) | N (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **m** (lowercase) | n (lowercase) |

### N (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **N** (uppercase) | M (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **n** (lowercase) | m (lowercase) |

### O (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **O** (uppercase) | 0 (zero)      |
| **O** (uppercase) | D (uppercase) |
| **O** (uppercase) | o (lowercase) |
| **O** (uppercase) | Q (uppercase) |
| **O** (uppercase) | U (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **o** (lowercase) | 0 (zero)      |
| **O** (lowercase) | D (uppercase) |
| **O** (lowercase) | O (uppercase) |
| **O** (lowercase) | Q (uppercase) |
| **O** (lowercase) | U (uppercase) |

### P (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **P** (uppercase) | B (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **p** (lowercase) | n (lowercase) |

### Q (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **Q** (uppercase) | 0 (zero)      |
| **Q** (uppercase) | D (uppercase) |
| **Q** (uppercase) | O (uppercase) |
| **Q** (uppercase) | o (lowercase) |
| **Q** (uppercase) | U (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **q** (lowercase) | 9 (nine)      |
| **q** (lowercase) | a (lowercase) |
| **q** (lowercase) | g (lowercase) |

### R (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **R** (uppercase) | F (uppercase) |

**No common confusion for lowercase “r”.**

### S (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **S** (uppercase) | 5 (five)      |
| **S** (uppercase) | 8 (eight)     |
| **S** (uppercase) | s (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **s** (lowercase) | 5 (five)      |
| **s** (lowercase) | 8 (eight)     |
| **s** (lowercase) | s (lowercase) |

### T (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **T** (uppercase) | 7 (seven)     |
| **T** (uppercase) | I (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **t** (lowercase) | f (lowercase) |

### U (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **U** (uppercase) | 0 (zero)      |
| **U** (uppercase) | 4 (four)      |
| **U** (uppercase) | O (uppercase) |
| **U** (uppercase) | o (lowercase) |
| **U** (uppercase) | Q (uppercase) |
| **U** (uppercase) | u (lowercase) |
| **U** (uppercase) | V (uppercase) |
| **U** (uppercase) | v (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **u** (lowercase) | 0 (zero)      |
| **u** (lowercase) | 4 (four)      |
| **u** (lowercase) | O (uppercase) |
| **u** (lowercase) | o (lowercase) |
| **u** (lowercase) | Q (uppercase) |
| **u** (lowercase) | u (lowercase) |
| **u** (lowercase) | V (uppercase) |
| **u** (lowercase) | v (lowercase) |

### V (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **V** (uppercase) | U (uppercase) |
| **V** (uppercase) | u (lowercase) |
| **V** (uppercase) | v (lowercase) |
| **V** (uppercase) | W (uppercase) |
| **V** (uppercase) | w (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **v** (lowercase) | U (uppercase) |
| **v** (lowercase) | u (lowercase) |
| **v** (lowercase) | V (uppercase) |
| **v** (lowercase) | W (uppercase) |
| **v** (lowercase) | w (lowercase) |

### W (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **W** (uppercase) | V (uppercase) |
| **W** (uppercase) | v (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **w** (lowercase) | V (uppercase) |
| **w** (lowercase) | v (lowercase) |

### X (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **X** (uppercase) | x (lowercase) |
| **X** (uppercase) | Y (uppercase) |
| **X** (uppercase) | y (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **x** (lowercase) | X (uppercase) |
| **x** (lowercase) | Y (uppercase) |
| **x** (lowercase) | y (lowercase) |

### Y (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **Y** (uppercase) | 5 (five)      |
| **Y** (uppercase) | X (uppercase) |
| **Y** (uppercase) | x (lowercase) |
| **Y** (uppercase) | y (lowercase) |
| **Y** (uppercase) | Z (uppercase) |
| **Y** (uppercase) | z (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **y** (lowercase) | 5 (five)      |
| **y** (lowercase) | X (uppercase) |
| **y** (lowercase) | x (lowercase) |
| **y** (lowercase) | Y (uppercase) |
| **y** (lowercase) | Z (uppercase) |
| **y** (lowercase) | z (lowercase) |

### Z (Alphanumeric Confusion)

|     Character     | Confused With |
|:----------------- |:------------- |
| **Z** (uppercase) | 2 (two)       |
| **Z** (uppercase) | 7 (seven)     |
| **Z** (uppercase) | Y (uppercase) |
| **Z** (uppercase) | y (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **z** (lowercase) | 2 (two)       |
| **z** (lowercase) | 7 (seven)     |
| **z** (lowercase) | Y (uppercase) |
| **z** (lowercase) | y (lowercase) |

### Golang (Alphanumeric Confusion)

We could encode these _alphanumeric confusion_ relations in Golang with a structure such as:
```go
var alphanumericConfusion map[rune]map[rune]struct{}{
  '0': map[rune]struct{}{},
  '1': map[rune]struct{}{},
  '2': map[rune]struct{}{},
  '3': map[rune]struct{}{},
  '4': map[rune]struct{}{},
  '5': map[rune]struct{}{},
  '6': map[rune]struct{}{},
  '7': map[rune]struct{}{},
  '8': map[rune]struct{}{},
  '9': map[rune]struct{}{},
  'A': map[rune]struct{}{},
  'a': map[rune]struct{}{},
  'B': map[rune]struct{}{},
  'C': map[rune]struct{}{},
  'c': map[rune]struct{}{},
  'D': map[rune]struct{}{},
  'd': map[rune]struct{}{},
  'E': map[rune]struct{}{},
  'e': map[rune]struct{}{},
  'F': map[rune]struct{}{},
  'f': map[rune]struct{}{},
  'G': map[rune]struct{}{},
  'g': map[rune]struct{}{},
  'H': map[rune]struct{}{},
  'h': map[rune]struct{}{},
  'I': map[rune]struct{}{},
  'i': map[rune]struct{}{},
  'J': map[rune]struct{}{},
  'j': map[rune]struct{}{},
  'K': map[rune]struct{}{},
  'k': map[rune]struct{}{},
  'L': map[rune]struct{}{},
  'l': map[rune]struct{}{},
  'M': map[rune]struct{}{},
  'm': map[rune]struct{}{},
  'N': map[rune]struct{}{},
  'n': map[rune]struct{}{},
  'O': map[rune]struct{}{},
  'o': map[rune]struct{}{},
  'P': map[rune]struct{}{},
  'p': map[rune]struct{}{},
  'Q': map[rune]struct{}{},
  'q': map[rune]struct{}{},
  'R': map[rune]struct{}{},
  'r': map[rune]struct{}{},
  'S': map[rune]struct{}{},
  's': map[rune]struct{}{},
  'T': map[rune]struct{}{},
  't': map[rune]struct{}{},
  'U': map[rune]struct{}{},
  'u': map[rune]struct{}{},
  'V': map[rune]struct{}{},
  'v': map[rune]struct{}{},
  'W': map[rune]struct{}{},
  'w': map[rune]struct{}{},
  'X': map[rune]struct{}{},
  'x': map[rune]struct{}{},
  'Y': map[rune]struct{}{},
  'y': map[rune]struct{}{},
  'Z': map[rune]struct{}{},
  'z': map[rune]struct{}{},
}
```
Where the _keys_ of the outer `map` is the value from the _character_ column in our _alphanumeric confusion_ tables.
