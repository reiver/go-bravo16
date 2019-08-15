package bravo16_test

import (
	"github.com/reiver/go-bravo16"

	"bytes"

	"testing"
)

func TestEncode(t *testing.T) {

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
			Src:      []byte("apple"),
			Expected: []byte("x1V0V0xCx5"),
		},
		{
			Src:      []byte("BANANA"),
			Expected: []byte("42414E414E41"),
		},
		{
			Src:      []byte("Cherry"),
			Expected: []byte("4mxHx5V2V2Vk"),
		},
		{
			Src:      []byte("dATE"),
			Expected: []byte("x4415445"),
		},



		{
			Src:      []byte("Hello world!"),
			Expected: []byte("4Hx5xCxCxf20VVxfV2xCx421"),
		},



		{
			Src:      []byte{0xc0, 0x53, 0x5e, 0x4b, 0xe2, 0xb7, 0x9f, 0xfd, 0x93, 0x29, 0x13, 0x05, 0x43, 0x6b, 0xf8, 0x89, 0x31, 0x4e, 0x4a, 0x3f, 0xae, 0xc0, 0x5e, 0xcf, 0xfc, 0xbb, 0x7d, 0xf3, 0x1a, 0xd9, 0xe5, 0x1a},
			Expected: []byte("C05m5E4bE2bVkffdkm2k1m054mxbfHHkm14E4AmfAEC05ECffCbbVdfm1AdkE51A"),
		},



		{
			Src:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, 0x9c, 0x08, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1, 0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f},
			Expected: []byte("00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf"),
		},



		{
			Src:      []byte{0x00},
			Expected: []byte( "00"),
		},
		{
			Src:      []byte{0x01},
			Expected: []byte( "01"),
		},
		{
			Src:      []byte{0x02},
			Expected: []byte( "02"),
		},
		{
			Src:      []byte{0x03},
			Expected: []byte( "0m"),
		},
		{
			Src:      []byte{0x04},
			Expected: []byte( "04"),
		},
		{
			Src:      []byte{0x05},
			Expected: []byte( "05"),
		},
		{
			Src:      []byte{0x06},
			Expected: []byte( "0x"),
		},
		{
			Src:      []byte{0x07},
			Expected: []byte( "0V"),
		},
		{
			Src:      []byte{0x08},
			Expected: []byte( "0H"),
		},
		{
			Src:      []byte{0x09},
			Expected: []byte( "0k"),
		},
		{
			Src:      []byte{0x0a},
			Expected: []byte( "0A"),
		},
		{
			Src:      []byte{0x0b},
			Expected: []byte( "0b"),
		},
		{
			Src:      []byte{0x0c},
			Expected: []byte( "0C"),
		},
		{
			Src:      []byte{0x0d},
			Expected: []byte( "0d"),
		},
		{
			Src:      []byte{0x0e},
			Expected: []byte( "0E"),
		},
		{
			Src:      []byte{0x0f},
			Expected: []byte( "0f"),
		},



		{
			Src:      []byte{0x10},
			Expected: []byte( "10"),
		},
		{
			Src:      []byte{0x11},
			Expected: []byte( "11"),
		},
		{
			Src:      []byte{0x12},
			Expected: []byte( "12"),
		},
		{
			Src:      []byte{0x13},
			Expected: []byte( "1m"),
		},
		{
			Src:      []byte{0x14},
			Expected: []byte( "14"),
		},
		{
			Src:      []byte{0x15},
			Expected: []byte( "15"),
		},
		{
			Src:      []byte{0x16},
			Expected: []byte( "1x"),
		},
		{
			Src:      []byte{0x17},
			Expected: []byte( "1V"),
		},
		{
			Src:      []byte{0x18},
			Expected: []byte( "1H"),
		},
		{
			Src:      []byte{0x19},
			Expected: []byte( "1k"),
		},
		{
			Src:      []byte{0x1a},
			Expected: []byte( "1A"),
		},
		{
			Src:      []byte{0x1b},
			Expected: []byte( "1b"),
		},
		{
			Src:      []byte{0x1c},
			Expected: []byte( "1C"),
		},
		{
			Src:      []byte{0x1d},
			Expected: []byte( "1d"),
		},
		{
			Src:      []byte{0x1e},
			Expected: []byte( "1E"),
		},
		{
			Src:      []byte{0x1f},
			Expected: []byte( "1f"),
		},



		{
			Src:      []byte{0x20},
			Expected: []byte( "20"),
		},
		{
			Src:      []byte{0x21},
			Expected: []byte( "21"),
		},
		{
			Src:      []byte{0x22},
			Expected: []byte( "22"),
		},
		{
			Src:      []byte{0x23},
			Expected: []byte( "2m"),
		},
		{
			Src:      []byte{0x24},
			Expected: []byte( "24"),
		},
		{
			Src:      []byte{0x25},
			Expected: []byte( "25"),
		},
		{
			Src:      []byte{0x26},
			Expected: []byte( "2x"),
		},
		{
			Src:      []byte{0x27},
			Expected: []byte( "2V"),
		},
		{
			Src:      []byte{0x28},
			Expected: []byte( "2H"),
		},
		{
			Src:      []byte{0x29},
			Expected: []byte( "2k"),
		},
		{
			Src:      []byte{0x2a},
			Expected: []byte( "2A"),
		},
		{
			Src:      []byte{0x2b},
			Expected: []byte( "2b"),
		},
		{
			Src:      []byte{0x2c},
			Expected: []byte( "2C"),
		},
		{
			Src:      []byte{0x2d},
			Expected: []byte( "2d"),
		},
		{
			Src:      []byte{0x2e},
			Expected: []byte( "2E"),
		},
		{
			Src:      []byte{0x2f},
			Expected: []byte( "2f"),
		},



		{
			Src:      []byte{0x30},
			Expected: []byte( "m0"),
		},
		{
			Src:      []byte{0x31},
			Expected: []byte( "m1"),
		},
		{
			Src:      []byte{0x32},
			Expected: []byte( "m2"),
		},
		{
			Src:      []byte{0x33},
			Expected: []byte( "mm"),
		},
		{
			Src:      []byte{0x34},
			Expected: []byte( "m4"),
		},
		{
			Src:      []byte{0x35},
			Expected: []byte( "m5"),
		},
		{
			Src:      []byte{0x36},
			Expected: []byte( "mx"),
		},
		{
			Src:      []byte{0x37},
			Expected: []byte( "mV"),
		},
		{
			Src:      []byte{0x38},
			Expected: []byte( "mH"),
		},
		{
			Src:      []byte{0x39},
			Expected: []byte( "mk"),
		},
		{
			Src:      []byte{0x3a},
			Expected: []byte( "mA"),
		},
		{
			Src:      []byte{0x3b},
			Expected: []byte( "mb"),
		},
		{
			Src:      []byte{0x3c},
			Expected: []byte( "mC"),
		},
		{
			Src:      []byte{0x3d},
			Expected: []byte( "md"),
		},
		{
			Src:      []byte{0x3e},
			Expected: []byte( "mE"),
		},
		{
			Src:      []byte{0x3f},
			Expected: []byte( "mf"),
		},



		{
			Src:      []byte{0x40},
			Expected: []byte( "40"),
		},
		{
			Src:      []byte{0x41},
			Expected: []byte( "41"),
		},
		{
			Src:      []byte{0x42},
			Expected: []byte( "42"),
		},
		{
			Src:      []byte{0x43},
			Expected: []byte( "4m"),
		},
		{
			Src:      []byte{0x44},
			Expected: []byte( "44"),
		},
		{
			Src:      []byte{0x45},
			Expected: []byte( "45"),
		},
		{
			Src:      []byte{0x46},
			Expected: []byte( "4x"),
		},
		{
			Src:      []byte{0x47},
			Expected: []byte( "4V"),
		},
		{
			Src:      []byte{0x48},
			Expected: []byte( "4H"),
		},
		{
			Src:      []byte{0x49},
			Expected: []byte( "4k"),
		},
		{
			Src:      []byte{0x4a},
			Expected: []byte( "4A"),
		},
		{
			Src:      []byte{0x4b},
			Expected: []byte( "4b"),
		},
		{
			Src:      []byte{0x4c},
			Expected: []byte( "4C"),
		},
		{
			Src:      []byte{0x4d},
			Expected: []byte( "4d"),
		},
		{
			Src:      []byte{0x4e},
			Expected: []byte( "4E"),
		},
		{
			Src:      []byte{0x4f},
			Expected: []byte( "4f"),
		},



		{
			Src:      []byte{0x50},
			Expected: []byte( "50"),
		},
		{
			Src:      []byte{0x51},
			Expected: []byte( "51"),
		},
		{
			Src:      []byte{0x52},
			Expected: []byte( "52"),
		},
		{
			Src:      []byte{0x53},
			Expected: []byte( "5m"),
		},
		{
			Src:      []byte{0x54},
			Expected: []byte( "54"),
		},
		{
			Src:      []byte{0x55},
			Expected: []byte( "55"),
		},
		{
			Src:      []byte{0x56},
			Expected: []byte( "5x"),
		},
		{
			Src:      []byte{0x57},
			Expected: []byte( "5V"),
		},
		{
			Src:      []byte{0x58},
			Expected: []byte( "5H"),
		},
		{
			Src:      []byte{0x59},
			Expected: []byte( "5k"),
		},
		{
			Src:      []byte{0x5a},
			Expected: []byte( "5A"),
		},
		{
			Src:      []byte{0x5b},
			Expected: []byte( "5b"),
		},
		{
			Src:      []byte{0x5c},
			Expected: []byte( "5C"),
		},
		{
			Src:      []byte{0x5d},
			Expected: []byte( "5d"),
		},
		{
			Src:      []byte{0x5e},
			Expected: []byte( "5E"),
		},
		{
			Src:      []byte{0x5f},
			Expected: []byte( "5f"),
		},



		{
			Src:      []byte{0x60},
			Expected: []byte( "x0"),
		},
		{
			Src:      []byte{0x61},
			Expected: []byte( "x1"),
		},
		{
			Src:      []byte{0x62},
			Expected: []byte( "x2"),
		},
		{
			Src:      []byte{0x63},
			Expected: []byte( "xm"),
		},
		{
			Src:      []byte{0x64},
			Expected: []byte( "x4"),
		},
		{
			Src:      []byte{0x65},
			Expected: []byte( "x5"),
		},
		{
			Src:      []byte{0x66},
			Expected: []byte( "xx"),
		},
		{
			Src:      []byte{0x67},
			Expected: []byte( "xV"),
		},
		{
			Src:      []byte{0x68},
			Expected: []byte( "xH"),
		},
		{
			Src:      []byte{0x69},
			Expected: []byte( "xk"),
		},
		{
			Src:      []byte{0x6a},
			Expected: []byte( "xA"),
		},
		{
			Src:      []byte{0x6b},
			Expected: []byte( "xb"),
		},
		{
			Src:      []byte{0x6c},
			Expected: []byte( "xC"),
		},
		{
			Src:      []byte{0x6d},
			Expected: []byte( "xd"),
		},
		{
			Src:      []byte{0x6e},
			Expected: []byte( "xE"),
		},
		{
			Src:      []byte{0x6f},
			Expected: []byte( "xf"),
		},



		{
			Src:      []byte{0x70},
			Expected: []byte( "V0"),
		},
		{
			Src:      []byte{0x71},
			Expected: []byte( "V1"),
		},
		{
			Src:      []byte{0x72},
			Expected: []byte( "V2"),
		},
		{
			Src:      []byte{0x73},
			Expected: []byte( "Vm"),
		},
		{
			Src:      []byte{0x74},
			Expected: []byte( "V4"),
		},
		{
			Src:      []byte{0x75},
			Expected: []byte( "V5"),
		},
		{
			Src:      []byte{0x76},
			Expected: []byte( "Vx"),
		},
		{
			Src:      []byte{0x77},
			Expected: []byte( "VV"),
		},
		{
			Src:      []byte{0x78},
			Expected: []byte( "VH"),
		},
		{
			Src:      []byte{0x79},
			Expected: []byte( "Vk"),
		},
		{
			Src:      []byte{0x7a},
			Expected: []byte( "VA"),
		},
		{
			Src:      []byte{0x7b},
			Expected: []byte( "Vb"),
		},
		{
			Src:      []byte{0x7c},
			Expected: []byte( "VC"),
		},
		{
			Src:      []byte{0x7d},
			Expected: []byte( "Vd"),
		},
		{
			Src:      []byte{0x7e},
			Expected: []byte( "VE"),
		},
		{
			Src:      []byte{0x7f},
			Expected: []byte( "Vf"),
		},



		{
			Src:      []byte{0x80},
			Expected: []byte( "H0"),
		},
		{
			Src:      []byte{0x81},
			Expected: []byte( "H1"),
		},
		{
			Src:      []byte{0x82},
			Expected: []byte( "H2"),
		},
		{
			Src:      []byte{0x83},
			Expected: []byte( "Hm"),
		},
		{
			Src:      []byte{0x84},
			Expected: []byte( "H4"),
		},
		{
			Src:      []byte{0x85},
			Expected: []byte( "H5"),
		},
		{
			Src:      []byte{0x86},
			Expected: []byte( "Hx"),
		},
		{
			Src:      []byte{0x87},
			Expected: []byte( "HV"),
		},
		{
			Src:      []byte{0x88},
			Expected: []byte( "HH"),
		},
		{
			Src:      []byte{0x89},
			Expected: []byte( "Hk"),
		},
		{
			Src:      []byte{0x8a},
			Expected: []byte( "HA"),
		},
		{
			Src:      []byte{0x8b},
			Expected: []byte( "Hb"),
		},
		{
			Src:      []byte{0x8c},
			Expected: []byte( "HC"),
		},
		{
			Src:      []byte{0x8d},
			Expected: []byte( "Hd"),
		},
		{
			Src:      []byte{0x8e},
			Expected: []byte( "HE"),
		},
		{
			Src:      []byte{0x8f},
			Expected: []byte( "Hf"),
		},



		{
			Src:      []byte{0x90},
			Expected: []byte( "k0"),
		},
		{
			Src:      []byte{0x91},
			Expected: []byte( "k1"),
		},
		{
			Src:      []byte{0x92},
			Expected: []byte( "k2"),
		},
		{
			Src:      []byte{0x93},
			Expected: []byte( "km"),
		},
		{
			Src:      []byte{0x94},
			Expected: []byte( "k4"),
		},
		{
			Src:      []byte{0x95},
			Expected: []byte( "k5"),
		},
		{
			Src:      []byte{0x96},
			Expected: []byte( "kx"),
		},
		{
			Src:      []byte{0x97},
			Expected: []byte( "kV"),
		},
		{
			Src:      []byte{0x98},
			Expected: []byte( "kH"),
		},
		{
			Src:      []byte{0x99},
			Expected: []byte( "kk"),
		},
		{
			Src:      []byte{0x9a},
			Expected: []byte( "kA"),
		},
		{
			Src:      []byte{0x9b},
			Expected: []byte( "kb"),
		},
		{
			Src:      []byte{0x9c},
			Expected: []byte( "kC"),
		},
		{
			Src:      []byte{0x9d},
			Expected: []byte( "kd"),
		},
		{
			Src:      []byte{0x9e},
			Expected: []byte( "kE"),
		},
		{
			Src:      []byte{0x9f},
			Expected: []byte( "kf"),
		},



		{
			Src:      []byte{0xa0},
			Expected: []byte( "A0"),
		},
		{
			Src:      []byte{0xa1},
			Expected: []byte( "A1"),
		},
		{
			Src:      []byte{0xa2},
			Expected: []byte( "A2"),
		},
		{
			Src:      []byte{0xa3},
			Expected: []byte( "Am"),
		},
		{
			Src:      []byte{0xa4},
			Expected: []byte( "A4"),
		},
		{
			Src:      []byte{0xa5},
			Expected: []byte( "A5"),
		},
		{
			Src:      []byte{0xa6},
			Expected: []byte( "Ax"),
		},
		{
			Src:      []byte{0xa7},
			Expected: []byte( "AV"),
		},
		{
			Src:      []byte{0xa8},
			Expected: []byte( "AH"),
		},
		{
			Src:      []byte{0xa9},
			Expected: []byte( "Ak"),
		},
		{
			Src:      []byte{0xaa},
			Expected: []byte( "AA"),
		},
		{
			Src:      []byte{0xab},
			Expected: []byte( "Ab"),
		},
		{
			Src:      []byte{0xac},
			Expected: []byte( "AC"),
		},
		{
			Src:      []byte{0xad},
			Expected: []byte( "Ad"),
		},
		{
			Src:      []byte{0xae},
			Expected: []byte( "AE"),
		},
		{
			Src:      []byte{0xaf},
			Expected: []byte( "Af"),
		},




		{
			Src:      []byte{0xb0},
			Expected: []byte( "b0"),
		},
		{
			Src:      []byte{0xb1},
			Expected: []byte( "b1"),
		},
		{
			Src:      []byte{0xb2},
			Expected: []byte( "b2"),
		},
		{
			Src:      []byte{0xb3},
			Expected: []byte( "bm"),
		},
		{
			Src:      []byte{0xb4},
			Expected: []byte( "b4"),
		},
		{
			Src:      []byte{0xb5},
			Expected: []byte( "b5"),
		},
		{
			Src:      []byte{0xb6},
			Expected: []byte( "bx"),
		},
		{
			Src:      []byte{0xb7},
			Expected: []byte( "bV"),
		},
		{
			Src:      []byte{0xb8},
			Expected: []byte( "bH"),
		},
		{
			Src:      []byte{0xb9},
			Expected: []byte( "bk"),
		},
		{
			Src:      []byte{0xba},
			Expected: []byte( "bA"),
		},
		{
			Src:      []byte{0xbb},
			Expected: []byte( "bb"),
		},
		{
			Src:      []byte{0xbc},
			Expected: []byte( "bC"),
		},
		{
			Src:      []byte{0xbd},
			Expected: []byte( "bd"),
		},
		{
			Src:      []byte{0xbe},
			Expected: []byte( "bE"),
		},
		{
			Src:      []byte{0xbf},
			Expected: []byte( "bf"),
		},



		{
			Src:      []byte{0xc0},
			Expected: []byte( "C0"),
		},
		{
			Src:      []byte{0xc1},
			Expected: []byte( "C1"),
		},
		{
			Src:      []byte{0xc2},
			Expected: []byte( "C2"),
		},
		{
			Src:      []byte{0xc3},
			Expected: []byte( "Cm"),
		},
		{
			Src:      []byte{0xc4},
			Expected: []byte( "C4"),
		},
		{
			Src:      []byte{0xc5},
			Expected: []byte( "C5"),
		},
		{
			Src:      []byte{0xc6},
			Expected: []byte( "Cx"),
		},
		{
			Src:      []byte{0xc7},
			Expected: []byte( "CV"),
		},
		{
			Src:      []byte{0xc8},
			Expected: []byte( "CH"),
		},
		{
			Src:      []byte{0xc9},
			Expected: []byte( "Ck"),
		},
		{
			Src:      []byte{0xca},
			Expected: []byte( "CA"),
		},
		{
			Src:      []byte{0xcb},
			Expected: []byte( "Cb"),
		},
		{
			Src:      []byte{0xcc},
			Expected: []byte( "CC"),
		},
		{
			Src:      []byte{0xcd},
			Expected: []byte( "Cd"),
		},
		{
			Src:      []byte{0xce},
			Expected: []byte( "CE"),
		},
		{
			Src:      []byte{0xcf},
			Expected: []byte( "Cf"),
		},



		{
			Src:      []byte{0xd0},
			Expected: []byte( "d0"),
		},
		{
			Src:      []byte{0xd1},
			Expected: []byte( "d1"),
		},
		{
			Src:      []byte{0xd2},
			Expected: []byte( "d2"),
		},
		{
			Src:      []byte{0xd3},
			Expected: []byte( "dm"),
		},
		{
			Src:      []byte{0xd4},
			Expected: []byte( "d4"),
		},
		{
			Src:      []byte{0xd5},
			Expected: []byte( "d5"),
		},
		{
			Src:      []byte{0xd6},
			Expected: []byte( "dx"),
		},
		{
			Src:      []byte{0xd7},
			Expected: []byte( "dV"),
		},
		{
			Src:      []byte{0xd8},
			Expected: []byte( "dH"),
		},
		{
			Src:      []byte{0xd9},
			Expected: []byte( "dk"),
		},
		{
			Src:      []byte{0xda},
			Expected: []byte( "dA"),
		},
		{
			Src:      []byte{0xdb},
			Expected: []byte( "db"),
		},
		{
			Src:      []byte{0xdc},
			Expected: []byte( "dC"),
		},
		{
			Src:      []byte{0xdd},
			Expected: []byte( "dd"),
		},
		{
			Src:      []byte{0xde},
			Expected: []byte( "dE"),
		},
		{
			Src:      []byte{0xdf},
			Expected: []byte( "df"),
		},



		{
			Src:      []byte{0xe0},
			Expected: []byte( "E0"),
		},
		{
			Src:      []byte{0xe1},
			Expected: []byte( "E1"),
		},
		{
			Src:      []byte{0xe2},
			Expected: []byte( "E2"),
		},
		{
			Src:      []byte{0xe3},
			Expected: []byte( "Em"),
		},
		{
			Src:      []byte{0xe4},
			Expected: []byte( "E4"),
		},
		{
			Src:      []byte{0xe5},
			Expected: []byte( "E5"),
		},
		{
			Src:      []byte{0xe6},
			Expected: []byte( "Ex"),
		},
		{
			Src:      []byte{0xe7},
			Expected: []byte( "EV"),
		},
		{
			Src:      []byte{0xe8},
			Expected: []byte( "EH"),
		},
		{
			Src:      []byte{0xe9},
			Expected: []byte( "Ek"),
		},
		{
			Src:      []byte{0xea},
			Expected: []byte( "EA"),
		},
		{
			Src:      []byte{0xeb},
			Expected: []byte( "Eb"),
		},
		{
			Src:      []byte{0xec},
			Expected: []byte( "EC"),
		},
		{
			Src:      []byte{0xed},
			Expected: []byte( "Ed"),
		},
		{
			Src:      []byte{0xee},
			Expected: []byte( "EE"),
		},
		{
			Src:      []byte{0xef},
			Expected: []byte( "Ef"),
		},



		{
			Src:      []byte{0xf0},
			Expected: []byte( "f0"),
		},
		{
			Src:      []byte{0xf1},
			Expected: []byte( "f1"),
		},
		{
			Src:      []byte{0xf2},
			Expected: []byte( "f2"),
		},
		{
			Src:      []byte{0xf3},
			Expected: []byte( "fm"),
		},
		{
			Src:      []byte{0xf4},
			Expected: []byte( "f4"),
		},
		{
			Src:      []byte{0xf5},
			Expected: []byte( "f5"),
		},
		{
			Src:      []byte{0xf6},
			Expected: []byte( "fx"),
		},
		{
			Src:      []byte{0xf7},
			Expected: []byte( "fV"),
		},
		{
			Src:      []byte{0xf8},
			Expected: []byte( "fH"),
		},
		{
			Src:      []byte{0xf9},
			Expected: []byte( "fk"),
		},
		{
			Src:      []byte{0xfa},
			Expected: []byte( "fA"),
		},
		{
			Src:      []byte{0xfb},
			Expected: []byte( "fb"),
		},
		{
			Src:      []byte{0xfc},
			Expected: []byte( "fC"),
		},
		{
			Src:      []byte{0xfd},
			Expected: []byte( "fd"),
		},
		{
			Src:      []byte{0xfe},
			Expected: []byte( "fE"),
		},
		{
			Src:      []byte{0xff},
			Expected: []byte( "ff"),
		},
	}

	for testNumber, test := range tests {

		dstLen := bravo16.EncodeLen(len(test.Src))

		var dst []byte = make([]byte, dstLen)

		n, err := bravo16.Encode(dst, test.Src)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Logf("Src (string): %q", test.Src)
			t.Logf("Src (bytes): %#v", test.Src)
			continue
		}
		if expected, actual := int64(dstLen), n; expected != actual {
			t.Errorf("For test #%d, the actual number of bytes written at destination is not what was expected.", testNumber)
			t.Logf("Src (string): %q", test.Src)
			t.Logf("Src (bytes): %#v", test.Src)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}

		if expected, actual := test.Expected, dst; !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual bytes written at destination is not what was expected.", testNumber)
			t.Logf("Src (string): %q", test.Src)
			t.Logf("Src (bytes): %#v", test.Src)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
