package bravo16_test

import (
	"github.com/reiver/go-bravo16"

	"bytes"

	"testing"
)

func TestDecode(t *testing.T) {

	tests := []struct{
		Src []byte
		Expected []byte
	}{
		{
			Src:      []byte(nil),
			Expected: []byte(nil),
		},
		{
			Src:      []byte{},
			Expected: []byte{},
		},



		{
			Src:      []byte("x1V0V0xCx5"),
			Expected: []byte("apple"),
		},
		{
			Src:      []byte("42414E414E41"),
			Expected: []byte("BANANA"),
		},
		{
			Src:      []byte("4mxHx5V2V2Vk"),
			Expected: []byte("Cherry"),
		},
		{
			Src:      []byte("x4415445"),
			Expected: []byte("dATE"),
		},



		{
			Src:      []byte("4Hx5xCxCxf20VVxfV2xCx421"),
			Expected: []byte("Hello world!"),
		},




		{
			Src:      []byte( "C0"+ "5m"+ "5E"+ "4b"+ "E2"+ "bV"+ "kf"+ "fd"+ "km"+ "2k"+ "1m"+ "05"+ "4m"+ "xb"+ "fH"+ "Hk"+ "m1"+ "4E"+ "4A"+ "mf"+ "AE"+ "C0"+ "5E"+ "Cf"+ "fC"+ "bb"+ "Vd"+ "fm"+ "1A"+ "dk"+ "E5"+ "1A"),
			Expected: []byte{0xc0, 0x53, 0x5e, 0x4b, 0xe2, 0xb7, 0x9f, 0xfd, 0x93, 0x29, 0x13, 0x05, 0x43, 0x6b, 0xf8, 0x89, 0x31, 0x4e, 0x4a, 0x3f, 0xae, 0xc0, 0x5e, 0xcf, 0xfc, 0xbb, 0x7d, 0xf3, 0x1a, 0xd9, 0xe5, 0x1a},
		},



		{
			Src:      []byte( "00"+ "00"+ "00"+ "00"+ "00"+ "1k"+ "dx"+ "xH"+ "kC"+ "0H"+ "5A"+ "E1"+ "x5"+ "Hm"+ "1E"+ "km"+ "4f"+ "fV"+ "xm"+ "AE"+ "4x"+ "A2"+ "Ax"+ "C1"+ "V2"+ "bm"+ "f1"+ "bx"+ "0A"+ "HC"+ "E2"+ "xf"),
			Expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, 0x9c, 0x08, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1, 0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f},
		},



		{
			Src:      []byte( "00"),
			Expected: []byte{0x00},
		},
		{
			Src:      []byte( "01"),
			Expected: []byte{0x01},
		},
		{
			Src:      []byte( "02"),
			Expected: []byte{0x02},
		},
		{
			Src:      []byte( "0m"),
			Expected: []byte{0x03},
		},
		{
			Src:      []byte( "04"),
			Expected: []byte{0x04},
		},
		{
			Src:      []byte( "05"),
			Expected: []byte{0x05},
		},
		{
			Src:      []byte( "0x"),
			Expected: []byte{0x06},
		},
		{
			Src:      []byte( "0V"),
			Expected: []byte{0x07},
		},
		{
			Src:      []byte( "0H"),
			Expected: []byte{0x08},
		},
		{
			Src:      []byte( "0k"),
			Expected: []byte{0x09},
		},
		{
			Src:      []byte( "0A"),
			Expected: []byte{0x0a},
		},
		{
			Src:      []byte( "0b"),
			Expected: []byte{0x0b},
		},
		{
			Src:      []byte( "0C"),
			Expected: []byte{0x0c},
		},
		{
			Src:      []byte( "0d"),
			Expected: []byte{0x0d},
		},
		{
			Src:      []byte( "0E"),
			Expected: []byte{0x0e},
		},
		{
			Src:      []byte( "0f"),
			Expected: []byte{0x0f},
		},



		{
			Src:      []byte( "10"),
			Expected: []byte{0x10},
		},
		{
			Src:      []byte( "11"),
			Expected: []byte{0x11},
		},
		{
			Src:      []byte( "12"),
			Expected: []byte{0x12},
		},
		{
			Src:      []byte( "1m"),
			Expected: []byte{0x13},
		},
		{
			Src:      []byte( "14"),
			Expected: []byte{0x14},
		},
		{
			Src:      []byte( "15"),
			Expected: []byte{0x15},
		},
		{
			Src:      []byte( "1x"),
			Expected: []byte{0x16},
		},
		{
			Src:      []byte( "1V"),
			Expected: []byte{0x17},
		},
		{
			Src:      []byte( "1H"),
			Expected: []byte{0x18},
		},
		{
			Src:      []byte( "1k"),
			Expected: []byte{0x19},
		},
		{
			Src:      []byte( "1A"),
			Expected: []byte{0x1a},
		},
		{
			Src:      []byte( "1b"),
			Expected: []byte{0x1b},
		},
		{
			Src:      []byte( "1C"),
			Expected: []byte{0x1c},
		},
		{
			Src:      []byte( "1d"),
			Expected: []byte{0x1d},
		},
		{
			Src:      []byte( "1E"),
			Expected: []byte{0x1e},
		},
		{
			Src:      []byte( "1f"),
			Expected: []byte{0x1f},
		},



		{
			Src:      []byte( "20"),
			Expected: []byte{0x20},
		},
		{
			Src:      []byte( "21"),
			Expected: []byte{0x21},
		},
		{
			Src:      []byte( "22"),
			Expected: []byte{0x22},
		},
		{
			Src:      []byte( "2m"),
			Expected: []byte{0x23},
		},
		{
			Src:      []byte( "24"),
			Expected: []byte{0x24},
		},
		{
			Src:      []byte( "25"),
			Expected: []byte{0x25},
		},
		{
			Src:      []byte( "2x"),
			Expected: []byte{0x26},
		},
		{
			Src:      []byte( "2V"),
			Expected: []byte{0x27},
		},
		{
			Src:      []byte( "2H"),
			Expected: []byte{0x28},
		},
		{
			Src:      []byte( "2k"),
			Expected: []byte{0x29},
		},
		{
			Src:      []byte( "2A"),
			Expected: []byte{0x2a},
		},
		{
			Src:      []byte( "2b"),
			Expected: []byte{0x2b},
		},
		{
			Src:      []byte( "2C"),
			Expected: []byte{0x2c},
		},
		{
			Src:      []byte( "2d"),
			Expected: []byte{0x2d},
		},
		{
			Src:      []byte( "2E"),
			Expected: []byte{0x2e},
		},
		{
			Src:      []byte( "2f"),
			Expected: []byte{0x2f},
		},



		{
			Src:      []byte( "m0"),
			Expected: []byte{0x30},
		},
		{
			Src:      []byte( "m1"),
			Expected: []byte{0x31},
		},
		{
			Src:      []byte( "m2"),
			Expected: []byte{0x32},
		},
		{
			Src:      []byte( "mm"),
			Expected: []byte{0x33},
		},
		{
			Src:      []byte( "m4"),
			Expected: []byte{0x34},
		},
		{
			Src:      []byte( "m5"),
			Expected: []byte{0x35},
		},
		{
			Src:      []byte( "mx"),
			Expected: []byte{0x36},
		},
		{
			Src:      []byte( "mV"),
			Expected: []byte{0x37},
		},
		{
			Src:      []byte( "mH"),
			Expected: []byte{0x38},
		},
		{
			Src:      []byte( "mk"),
			Expected: []byte{0x39},
		},
		{
			Src:      []byte( "mA"),
			Expected: []byte{0x3a},
		},
		{
			Src:      []byte( "mb"),
			Expected: []byte{0x3b},
		},
		{
			Src:      []byte( "mC"),
			Expected: []byte{0x3c},
		},
		{
			Src:      []byte( "md"),
			Expected: []byte{0x3d},
		},
		{
			Src:      []byte( "mE"),
			Expected: []byte{0x3e},
		},
		{
			Src:      []byte( "mf"),
			Expected: []byte{0x3f},
		},



		{
			Src:      []byte( "40"),
			Expected: []byte{0x40},
		},
		{
			Src:      []byte( "41"),
			Expected: []byte{0x41},
		},
		{
			Src:      []byte( "42"),
			Expected: []byte{0x42},
		},
		{
			Src:      []byte( "4m"),
			Expected: []byte{0x43},
		},
		{
			Src:      []byte( "44"),
			Expected: []byte{0x44},
		},
		{
			Src:      []byte( "45"),
			Expected: []byte{0x45},
		},
		{
			Src:      []byte( "4x"),
			Expected: []byte{0x46},
		},
		{
			Src:      []byte( "4V"),
			Expected: []byte{0x47},
		},
		{
			Src:      []byte( "4H"),
			Expected: []byte{0x48},
		},
		{
			Src:      []byte( "4k"),
			Expected: []byte{0x49},
		},
		{
			Src:      []byte( "4A"),
			Expected: []byte{0x4a},
		},
		{
			Src:      []byte( "4b"),
			Expected: []byte{0x4b},
		},
		{
			Src:      []byte( "4C"),
			Expected: []byte{0x4c},
		},
		{
			Src:      []byte( "4d"),
			Expected: []byte{0x4d},
		},
		{
			Src:      []byte( "4E"),
			Expected: []byte{0x4e},
		},
		{
			Src:      []byte( "4f"),
			Expected: []byte{0x4f},
		},



		{
			Src:      []byte( "50"),
			Expected: []byte{0x50},
		},
		{
			Src:      []byte( "51"),
			Expected: []byte{0x51},
		},
		{
			Src:      []byte( "52"),
			Expected: []byte{0x52},
		},
		{
			Src:      []byte( "5m"),
			Expected: []byte{0x53},
		},
		{
			Src:      []byte( "54"),
			Expected: []byte{0x54},
		},
		{
			Src:      []byte( "55"),
			Expected: []byte{0x55},
		},
		{
			Src:      []byte( "5x"),
			Expected: []byte{0x56},
		},
		{
			Src:      []byte( "5V"),
			Expected: []byte{0x57},
		},
		{
			Src:      []byte( "5H"),
			Expected: []byte{0x58},
		},
		{
			Src:      []byte( "5k"),
			Expected: []byte{0x59},
		},
		{
			Src:      []byte( "5A"),
			Expected: []byte{0x5a},
		},
		{
			Src:      []byte( "5b"),
			Expected: []byte{0x5b},
		},
		{
			Src:      []byte( "5C"),
			Expected: []byte{0x5c},
		},
		{
			Src:      []byte( "5d"),
			Expected: []byte{0x5d},
		},
		{
			Src:      []byte( "5E"),
			Expected: []byte{0x5e},
		},
		{
			Src:      []byte( "5f"),
			Expected: []byte{0x5f},
		},



		{
			Src:      []byte( "x0"),
			Expected: []byte{0x60},
		},
		{
			Src:      []byte( "x1"),
			Expected: []byte{0x61},
		},
		{
			Src:      []byte( "x2"),
			Expected: []byte{0x62},
		},
		{
			Src:      []byte( "xm"),
			Expected: []byte{0x63},
		},
		{
			Src:      []byte( "x4"),
			Expected: []byte{0x64},
		},
		{
			Src:      []byte( "x5"),
			Expected: []byte{0x65},
		},
		{
			Src:      []byte( "xx"),
			Expected: []byte{0x66},
		},
		{
			Src:      []byte( "xV"),
			Expected: []byte{0x67},
		},
		{
			Src:      []byte( "xH"),
			Expected: []byte{0x68},
		},
		{
			Src:      []byte( "xk"),
			Expected: []byte{0x69},
		},
		{
			Src:      []byte( "xA"),
			Expected: []byte{0x6a},
		},
		{
			Src:      []byte( "xb"),
			Expected: []byte{0x6b},
		},
		{
			Src:      []byte( "xC"),
			Expected: []byte{0x6c},
		},
		{
			Src:      []byte( "xd"),
			Expected: []byte{0x6d},
		},
		{
			Src:      []byte( "xE"),
			Expected: []byte{0x6e},
		},
		{
			Src:      []byte( "xf"),
			Expected: []byte{0x6f},
		},



		{
			Src:      []byte( "V0"),
			Expected: []byte{0x70},
		},
		{
			Src:      []byte( "V1"),
			Expected: []byte{0x71},
		},
		{
			Src:      []byte( "V2"),
			Expected: []byte{0x72},
		},
		{
			Src:      []byte( "Vm"),
			Expected: []byte{0x73},
		},
		{
			Src:      []byte( "V4"),
			Expected: []byte{0x74},
		},
		{
			Src:      []byte( "V5"),
			Expected: []byte{0x75},
		},
		{
			Src:      []byte( "Vx"),
			Expected: []byte{0x76},
		},
		{
			Src:      []byte( "VV"),
			Expected: []byte{0x77},
		},
		{
			Src:      []byte( "VH"),
			Expected: []byte{0x78},
		},
		{
			Src:      []byte( "Vk"),
			Expected: []byte{0x79},
		},
		{
			Src:      []byte( "VA"),
			Expected: []byte{0x7a},
		},
		{
			Src:      []byte( "Vb"),
			Expected: []byte{0x7b},
		},
		{
			Src:      []byte( "VC"),
			Expected: []byte{0x7c},
		},
		{
			Src:      []byte( "Vd"),
			Expected: []byte{0x7d},
		},
		{
			Src:      []byte( "VE"),
			Expected: []byte{0x7e},
		},
		{
			Src:      []byte( "Vf"),
			Expected: []byte{0x7f},
		},



		{
			Src:      []byte( "H0"),
			Expected: []byte{0x80},
		},
		{
			Src:      []byte( "H1"),
			Expected: []byte{0x81},
		},
		{
			Src:      []byte( "H2"),
			Expected: []byte{0x82},
		},
		{
			Src:      []byte( "Hm"),
			Expected: []byte{0x83},
		},
		{
			Src:      []byte( "H4"),
			Expected: []byte{0x84},
		},
		{
			Src:      []byte( "H5"),
			Expected: []byte{0x85},
		},
		{
			Src:      []byte( "Hx"),
			Expected: []byte{0x86},
		},
		{
			Src:      []byte( "HV"),
			Expected: []byte{0x87},
		},
		{
			Src:      []byte( "HH"),
			Expected: []byte{0x88},
		},
		{
			Src:      []byte( "Hk"),
			Expected: []byte{0x89},
		},
		{
			Src:      []byte( "HA"),
			Expected: []byte{0x8a},
		},
		{
			Src:      []byte( "Hb"),
			Expected: []byte{0x8b},
		},
		{
			Src:      []byte( "HC"),
			Expected: []byte{0x8c},
		},
		{
			Src:      []byte( "Hd"),
			Expected: []byte{0x8d},
		},
		{
			Src:      []byte( "HE"),
			Expected: []byte{0x8e},
		},
		{
			Src:      []byte( "Hf"),
			Expected: []byte{0x8f},
		},



		{
			Src:      []byte( "k0"),
			Expected: []byte{0x90},
		},
		{
			Src:      []byte( "k1"),
			Expected: []byte{0x91},
		},
		{
			Src:      []byte( "k2"),
			Expected: []byte{0x92},
		},
		{
			Src:      []byte( "km"),
			Expected: []byte{0x93},
		},
		{
			Src:      []byte( "k4"),
			Expected: []byte{0x94},
		},
		{
			Src:      []byte( "k5"),
			Expected: []byte{0x95},
		},
		{
			Src:      []byte( "kx"),
			Expected: []byte{0x96},
		},
		{
			Src:      []byte( "kV"),
			Expected: []byte{0x97},
		},
		{
			Src:      []byte( "kH"),
			Expected: []byte{0x98},
		},
		{
			Src:      []byte( "kk"),
			Expected: []byte{0x99},
		},
		{
			Src:      []byte( "kA"),
			Expected: []byte{0x9a},
		},
		{
			Src:      []byte( "kb"),
			Expected: []byte{0x9b},
		},
		{
			Src:      []byte( "kC"),
			Expected: []byte{0x9c},
		},
		{
			Src:      []byte( "kd"),
			Expected: []byte{0x9d},
		},
		{
			Src:      []byte( "kE"),
			Expected: []byte{0x9e},
		},
		{
			Src:      []byte( "kf"),
			Expected: []byte{0x9f},
		},



		{
			Src:      []byte( "A0"),
			Expected: []byte{0xa0},
		},
		{
			Src:      []byte( "A1"),
			Expected: []byte{0xa1},
		},
		{
			Src:      []byte( "A2"),
			Expected: []byte{0xa2},
		},
		{
			Src:      []byte( "Am"),
			Expected: []byte{0xa3},
		},
		{
			Src:      []byte( "A4"),
			Expected: []byte{0xa4},
		},
		{
			Src:      []byte( "A5"),
			Expected: []byte{0xa5},
		},
		{
			Src:      []byte( "Ax"),
			Expected: []byte{0xa6},
		},
		{
			Src:      []byte( "AV"),
			Expected: []byte{0xa7},
		},
		{
			Src:      []byte( "AH"),
			Expected: []byte{0xa8},
		},
		{
			Src:      []byte( "Ak"),
			Expected: []byte{0xa9},
		},
		{
			Src:      []byte( "AA"),
			Expected: []byte{0xaa},
		},
		{
			Src:      []byte( "Ab"),
			Expected: []byte{0xab},
		},
		{
			Src:      []byte( "AC"),
			Expected: []byte{0xac},
		},
		{
			Src:      []byte( "Ad"),
			Expected: []byte{0xad},
		},
		{
			Src:      []byte( "AE"),
			Expected: []byte{0xae},
		},
		{
			Src:      []byte( "Af"),
			Expected: []byte{0xaf},
		},



		{
			Src:      []byte( "b0"),
			Expected: []byte{0xb0},
		},
		{
			Src:      []byte( "b1"),
			Expected: []byte{0xb1},
		},
		{
			Src:      []byte( "b2"),
			Expected: []byte{0xb2},
		},
		{
			Src:      []byte( "bm"),
			Expected: []byte{0xb3},
		},
		{
			Src:      []byte( "b4"),
			Expected: []byte{0xb4},
		},
		{
			Src:      []byte( "b5"),
			Expected: []byte{0xb5},
		},
		{
			Src:      []byte( "bx"),
			Expected: []byte{0xb6},
		},
		{
			Src:      []byte( "bV"),
			Expected: []byte{0xb7},
		},
		{
			Src:      []byte( "bH"),
			Expected: []byte{0xb8},
		},
		{
			Src:      []byte( "bk"),
			Expected: []byte{0xb9},
		},
		{
			Src:      []byte( "bA"),
			Expected: []byte{0xba},
		},
		{
			Src:      []byte( "bb"),
			Expected: []byte{0xbb},
		},
		{
			Src:      []byte( "bC"),
			Expected: []byte{0xbc},
		},
		{
			Src:      []byte( "bd"),
			Expected: []byte{0xbd},
		},
		{
			Src:      []byte( "bE"),
			Expected: []byte{0xbe},
		},
		{
			Src:      []byte( "bf"),
			Expected: []byte{0xbf},
		},



		{
			Src:      []byte( "C0"),
			Expected: []byte{0xc0},
		},
		{
			Src:      []byte( "C1"),
			Expected: []byte{0xc1},
		},
		{
			Src:      []byte( "C2"),
			Expected: []byte{0xc2},
		},
		{
			Src:      []byte( "Cm"),
			Expected: []byte{0xc3},
		},
		{
			Src:      []byte( "C4"),
			Expected: []byte{0xc4},
		},
		{
			Src:      []byte( "C5"),
			Expected: []byte{0xc5},
		},
		{
			Src:      []byte( "Cx"),
			Expected: []byte{0xc6},
		},
		{
			Src:      []byte( "CV"),
			Expected: []byte{0xc7},
		},
		{
			Src:      []byte( "CH"),
			Expected: []byte{0xc8},
		},
		{
			Src:      []byte( "Ck"),
			Expected: []byte{0xc9},
		},
		{
			Src:      []byte( "CA"),
			Expected: []byte{0xca},
		},
		{
			Src:      []byte( "Cb"),
			Expected: []byte{0xcb},
		},
		{
			Src:      []byte( "CC"),
			Expected: []byte{0xcc},
		},
		{
			Src:      []byte( "Cd"),
			Expected: []byte{0xcd},
		},
		{
			Src:      []byte( "CE"),
			Expected: []byte{0xce},
		},
		{
			Src:      []byte( "Cf"),
			Expected: []byte{0xcf},
		},



		{
			Src:      []byte( "d0"),
			Expected: []byte{0xd0},
		},
		{
			Src:      []byte( "d1"),
			Expected: []byte{0xd1},
		},
		{
			Src:      []byte( "d2"),
			Expected: []byte{0xd2},
		},
		{
			Src:      []byte( "dm"),
			Expected: []byte{0xd3},
		},
		{
			Src:      []byte( "d4"),
			Expected: []byte{0xd4},
		},
		{
			Src:      []byte( "d5"),
			Expected: []byte{0xd5},
		},
		{
			Src:      []byte( "dx"),
			Expected: []byte{0xd6},
		},
		{
			Src:      []byte( "dV"),
			Expected: []byte{0xd7},
		},
		{
			Src:      []byte( "dH"),
			Expected: []byte{0xd8},
		},
		{
			Src:      []byte( "dk"),
			Expected: []byte{0xd9},
		},
		{
			Src:      []byte( "dA"),
			Expected: []byte{0xda},
		},
		{
			Src:      []byte( "db"),
			Expected: []byte{0xdb},
		},
		{
			Src:      []byte( "dC"),
			Expected: []byte{0xdc},
		},
		{
			Src:      []byte( "dd"),
			Expected: []byte{0xdd},
		},
		{
			Src:      []byte( "dE"),
			Expected: []byte{0xde},
		},
		{
			Src:      []byte( "df"),
			Expected: []byte{0xdf},
		},



		{
			Src:      []byte( "E0"),
			Expected: []byte{0xe0},
		},
		{
			Src:      []byte( "E1"),
			Expected: []byte{0xe1},
		},
		{
			Src:      []byte( "E2"),
			Expected: []byte{0xe2},
		},
		{
			Src:      []byte( "Em"),
			Expected: []byte{0xe3},
		},
		{
			Src:      []byte( "E4"),
			Expected: []byte{0xe4},
		},
		{
			Src:      []byte( "E5"),
			Expected: []byte{0xe5},
		},
		{
			Src:      []byte( "Ex"),
			Expected: []byte{0xe6},
		},
		{
			Src:      []byte( "EV"),
			Expected: []byte{0xe7},
		},
		{
			Src:      []byte( "EH"),
			Expected: []byte{0xe8},
		},
		{
			Src:      []byte( "Ek"),
			Expected: []byte{0xe9},
		},
		{
			Src:      []byte( "EA"),
			Expected: []byte{0xea},
		},
		{
			Src:      []byte( "Eb"),
			Expected: []byte{0xeb},
		},
		{
			Src:      []byte( "EC"),
			Expected: []byte{0xec},
		},
		{
			Src:      []byte( "Ed"),
			Expected: []byte{0xed},
		},
		{
			Src:      []byte( "EE"),
			Expected: []byte{0xee},
		},
		{
			Src:      []byte( "Ef"),
			Expected: []byte{0xef},
		},



		{
			Src:      []byte( "f0"),
			Expected: []byte{0xf0},
		},
		{
			Src:      []byte( "f1"),
			Expected: []byte{0xf1},
		},
		{
			Src:      []byte( "f2"),
			Expected: []byte{0xf2},
		},
		{
			Src:      []byte( "fm"),
			Expected: []byte{0xf3},
		},
		{
			Src:      []byte( "f4"),
			Expected: []byte{0xf4},
		},
		{
			Src:      []byte( "f5"),
			Expected: []byte{0xf5},
		},
		{
			Src:      []byte( "fx"),
			Expected: []byte{0xf6},
		},
		{
			Src:      []byte( "fV"),
			Expected: []byte{0xf7},
		},
		{
			Src:      []byte( "fH"),
			Expected: []byte{0xf8},
		},
		{
			Src:      []byte( "fk"),
			Expected: []byte{0xf9},
		},
		{
			Src:      []byte( "fA"),
			Expected: []byte{0xfa},
		},
		{
			Src:      []byte( "fb"),
			Expected: []byte{0xfb},
		},
		{
			Src:      []byte( "fC"),
			Expected: []byte{0xfc},
		},
		{
			Src:      []byte( "fd"),
			Expected: []byte{0xfd},
		},
		{
			Src:      []byte( "fE"),
			Expected: []byte{0xfe},
		},
		{
			Src:      []byte( "ff"),
			Expected: []byte{0xff},
		},
	}

	for testNumber, test := range tests {
		var src []byte = test.Src

		dstLen := bravo16.DecodeLen(len(src))
		var dst []byte = make([]byte, dstLen)

		n, err := bravo16.Decode(dst, src)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Logf("Src (string): %q", src)
			t.Logf("Src (bytes): %#v", src)
			continue
		}
		if expected, actual := int64(len(dst)), n; expected != actual {
			t.Errorf("For test #%d, the actual number of bytes written at destination is not what was expected.", testNumber)
			t.Logf("Src (string): %q", src)
			t.Logf("Src (bytes): %#v", src)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}

		if expected, actual := test.Expected, dst; !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual bytes written at destination is not what was expected.", testNumber)
			t.Logf("Src (string): %q", src)
			t.Logf("Src (bytes): %#v", src)
			t.Logf("EXPECTED (string): %q", expected)
			t.Logf("ACTUAL   (string): %q", actual)
			t.Logf("EXPECTED (bytes): %#v", expected)
			t.Logf("ACTUAL   (bytes): %#v", actual)
			continue
		}
	}
}
