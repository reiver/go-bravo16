package bravo16

const (
	encode string = "012m45xVHkAbCdEf"
	// hexadecimal:  0123456789abcdef
	//                  |  ||||^ ^ ^
)

func decode(b byte) (byte, bool) {
	switch b {
	case '0':
		return 0, true
	case '1':
		return 1, true
	case '2':
		return 2, true
	case 'm':
		return 3, true
	case '4':
		return 4, true
	case '5':
		return 5, true
	case 'x','X':
		return 6, true
	case 'V','v':
		return 7, true
	case 'H':
		return 8, true
	case 'k','K':
		return 9, true
	case 'A':
		return 10, true
	case 'b':
		return 11, true
	case 'C':
		return 12, true
	case 'd':
		return 13, true
	case 'E':
		return 14, true
	case 'f':
		return 15, true
	default:
		return 0, false
	}
}
