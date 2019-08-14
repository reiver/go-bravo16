# go-bravo16

Package **bravo16** provides **safe** base-16 binary-to-text encoding, and decoding;
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

### Zero

|   Character   | Confused With |
| ------------- |:------------- |
|         **0** | 8 (eight)     |
|         **0** | D (uppercase) |
|         **0** | O (uppercase) |
|         **0** | o (lowercase) |

### One

|   Character   | Confused With |
| ------------- |:------------- |
|         **1** | 7 (seven)     |
|         **1** | I (uppercase) |
|         **1** | i (lowercase) |
|         **1** | j (lowercase) |
|         **1** | L (uppercase) |
|         **1** | l (lowercase) |

### Two

|   Character   | Confused With |
| ------------- |:------------- |
|         **2** | Z (uppercase) |
|         **2** | z (uppercase) |

### Three

|   Character   | Confused With |
| ------------- |:------------- |
|         **3** | 5 (five)      |
|         **3** | 8 (eight)     |
|         **3** | 9 (nine)      |

### Four

|   Character   | Confused With |
| ------------- |:------------- |
|         **4** | 9 (nine)      |

### Five

|   Character   | Confused With |
| ------------- |:------------- |
|         **5** | 3 (three)     |
|         **5** | 8 (eight)     |
|         **5** | S (uppercase) |
|         **5** | s (lowercase) |

### Six

|   Character   | Confused With |
| ------------- |:------------- |
|         **6** | 8 (eight)     |
|         **6** | b (lowercase) |
|         **6** | G (uppercase) |

### Seven

|   Character   | Confused With |
| ------------- |:------------- |
|         **7** | 1 (one)       |
|         **7** | F (uppercase) |
|         **7** | T (uppercase) |
|         **7** | Z (uppercase) |
|         **7** | z (lowercase) |

### Eight

|   Character   | Confused With |
| ------------- |:------------- |
|         **8** | 0 (zero)      |
|         **8** | 3 (three)     |
|         **8** | 5 (five)      |
|         **8** | 6 (six)       |
|         **8** | B (uppercase) |
|         **8** | S (uppercase) |
|         **8** | s (uppercase) |

### Nine

|   Character   | Confused With |
| ------------- |:------------- |
|         **9** | 3 (three)     |
|         **9** | 4 (four)      |
|         **9** | g (lowercase) |
|         **9** | q (lowercase) |

### A

|     Character     | Confused With |
|:----------------- |:------------- |
| **a** (lowercase) | o (lowercase) |
| **a** (lowercase) | q (lowercase) |

**No common confusion for uppercase “A”.**

### B

|     Character     | Confused With |
|:----------------- |:------------- |
| **B** (uppercase) | 8 (eight)     |

|     Character     | Confused With |
|:----------------- |:------------- |
| **b** (lowercase) | 6 (six)       |


### C

|     Character     | Confused With |
|:----------------- |:------------- |
| **C** (uppercase) | e (lowercase) | 
| **C** (uppercase) | G (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **c** (lowercase) | e (lowercase) |
| **c** (lowercase) | G (uppercase) |

### D

|     Character     | Confused With |
|:----------------- |:------------- |
| **D** (uppercase) | 0 (zero)      |
| **D** (uppercase) | O (uppercase) |
| **D** (uppercase) | o (lowercase) |

**No common confusion for lowercase “d”.**

### E

|     Character     | Confused With |
|:----------------- |:------------- |
| **E** (uppercase) | F (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **e** (lowercase) | C (uppercase) |
| **e** (lowercase) | c (lowercase) |

### F

|     Character     | Confused With |
|:----------------- |:------------- |
| **F** (uppercase) | E (uppercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **f** (lowercase) | t (lowercase) |

### G

|     Character     | Confused With |
|:----------------- |:------------- |
| **G** (uppercase) | 6 (six)       |
| **G** (uppercase) | C (uppercase) |
| **G** (uppercase) | c (lowercase) |

|     Character     | Confused With |
|:----------------- |:------------- |
| **g** (lowercase) | 9 (nine)      |
| **g** (lowercase) | q (lowercase) |

### H

**No common confusion for uppercase “H”.**

**No common confusion for lowercase “h”.**

### I

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

### J

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
