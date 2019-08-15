package bravo16_test

import (
	"github.com/reiver/go-bravo16"

	"bytes"

	"testing"
)

func TestDecodeLiteral(t *testing.T) {

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
			Src:      []byte("0rx1V0V0xCx5"),
			Expected: []byte("apple"),
		},
		{
			Src:      []byte("0r42414E414E41"),
			Expected: []byte("BANANA"),
		},
		{
			Src:      []byte("0r4mxHx5V2V2Vk"),
			Expected: []byte("Cherry"),
		},
		{
			Src:      []byte("0rx4415445"),
			Expected: []byte("dATE"),
		},



		{
			Src:      []byte("0r4Hx5xCxCxf20VVxfV2xCx421"),
			Expected: []byte("Hello world!"),
		},




		{
			Src:      []byte("0r"+ "C0"+ "5m"+ "5E"+ "4b"+ "E2"+ "bV"+ "kf"+ "fd"+ "km"+ "2k"+ "1m"+ "05"+ "4m"+ "xb"+ "fH"+ "Hk"+ "m1"+ "4E"+ "4A"+ "mf"+ "AE"+ "C0"+ "5E"+ "Cf"+ "fC"+ "bb"+ "Vd"+ "fm"+ "1A"+ "dk"+ "E5"+ "1A"),
			Expected: []byte{     0xc0, 0x53, 0x5e, 0x4b, 0xe2, 0xb7, 0x9f, 0xfd, 0x93, 0x29, 0x13, 0x05, 0x43, 0x6b, 0xf8, 0x89, 0x31, 0x4e, 0x4a, 0x3f, 0xae, 0xc0, 0x5e, 0xcf, 0xfc, 0xbb, 0x7d, 0xf3, 0x1a, 0xd9, 0xe5, 0x1a},
		},



		{
			Src:      []byte("0r"+ "00"+ "00"+ "00"+ "00"+ "00"+ "1k"+ "dx"+ "xH"+ "kC"+ "0H"+ "5A"+ "E1"+ "x5"+ "Hm"+ "1E"+ "km"+ "4f"+ "fV"+ "xm"+ "AE"+ "4x"+ "A2"+ "Ax"+ "C1"+ "V2"+ "bm"+ "f1"+ "bx"+ "0A"+ "HC"+ "E2"+ "xf"),
			Expected: []byte{     0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, 0x9c, 0x08, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1, 0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f},
		},



		{
			Src:      []byte("0r00"),
			Expected: []byte{ 0x00},
		},
		{
			Src:      []byte("0r01"),
			Expected: []byte{ 0x01},
		},
		{
			Src:      []byte("0r02"),
			Expected: []byte{ 0x02},
		},
		{
			Src:      []byte("0r0m"),
			Expected: []byte{ 0x03},
		},
		{
			Src:      []byte("0r04"),
			Expected: []byte{ 0x04},
		},
		{
			Src:      []byte("0r05"),
			Expected: []byte{ 0x05},
		},
		{
			Src:      []byte("0r0x"),
			Expected: []byte{ 0x06},
		},
		{
			Src:      []byte("0r0V"),
			Expected: []byte{ 0x07},
		},
		{
			Src:      []byte("0r0H"),
			Expected: []byte{ 0x08},
		},
		{
			Src:      []byte("0r0k"),
			Expected: []byte{ 0x09},
		},
		{
			Src:      []byte("0r0A"),
			Expected: []byte{ 0x0a},
		},
		{
			Src:      []byte("0r0b"),
			Expected: []byte{ 0x0b},
		},
		{
			Src:      []byte("0r0C"),
			Expected: []byte{ 0x0c},
		},
		{
			Src:      []byte("0r0d"),
			Expected: []byte{ 0x0d},
		},
		{
			Src:      []byte("0r0E"),
			Expected: []byte{ 0x0e},
		},
		{
			Src:      []byte("0r0f"),
			Expected: []byte{ 0x0f},
		},



		{
			Src:      []byte("0r10"),
			Expected: []byte{ 0x10},
		},
		{
			Src:      []byte("0r11"),
			Expected: []byte{ 0x11},
		},
		{
			Src:      []byte("0r12"),
			Expected: []byte{ 0x12},
		},
		{
			Src:      []byte("0r1m"),
			Expected: []byte{ 0x13},
		},
		{
			Src:      []byte("0r14"),
			Expected: []byte{ 0x14},
		},
		{
			Src:      []byte("0r15"),
			Expected: []byte{ 0x15},
		},
		{
			Src:      []byte("0r1x"),
			Expected: []byte{ 0x16},
		},
		{
			Src:      []byte("0r1V"),
			Expected: []byte{ 0x17},
		},
		{
			Src:      []byte("0r1H"),
			Expected: []byte{ 0x18},
		},
		{
			Src:      []byte("0r1k"),
			Expected: []byte{ 0x19},
		},
		{
			Src:      []byte("0r1A"),
			Expected: []byte{ 0x1a},
		},
		{
			Src:      []byte("0r1b"),
			Expected: []byte{ 0x1b},
		},
		{
			Src:      []byte("0r1C"),
			Expected: []byte{ 0x1c},
		},
		{
			Src:      []byte("0r1d"),
			Expected: []byte{ 0x1d},
		},
		{
			Src:      []byte("0r1E"),
			Expected: []byte{ 0x1e},
		},
		{
			Src:      []byte("0r1f"),
			Expected: []byte{ 0x1f},
		},



		{
			Src:      []byte("0r20"),
			Expected: []byte{ 0x20},
		},
		{
			Src:      []byte("0r21"),
			Expected: []byte{ 0x21},
		},
		{
			Src:      []byte("0r22"),
			Expected: []byte{ 0x22},
		},
		{
			Src:      []byte("0r2m"),
			Expected: []byte{ 0x23},
		},
		{
			Src:      []byte("0r24"),
			Expected: []byte{ 0x24},
		},
		{
			Src:      []byte("0r25"),
			Expected: []byte{ 0x25},
		},
		{
			Src:      []byte("0r2x"),
			Expected: []byte{ 0x26},
		},
		{
			Src:      []byte("0r2V"),
			Expected: []byte{ 0x27},
		},
		{
			Src:      []byte("0r2H"),
			Expected: []byte{ 0x28},
		},
		{
			Src:      []byte("0r2k"),
			Expected: []byte{ 0x29},
		},
		{
			Src:      []byte("0r2A"),
			Expected: []byte{ 0x2a},
		},
		{
			Src:      []byte("0r2b"),
			Expected: []byte{ 0x2b},
		},
		{
			Src:      []byte("0r2C"),
			Expected: []byte{ 0x2c},
		},
		{
			Src:      []byte("0r2d"),
			Expected: []byte{ 0x2d},
		},
		{
			Src:      []byte("0r2E"),
			Expected: []byte{ 0x2e},
		},
		{
			Src:      []byte("0r2f"),
			Expected: []byte{ 0x2f},
		},



		{
			Src:      []byte("0rm0"),
			Expected: []byte{ 0x30},
		},
		{
			Src:      []byte("0rm1"),
			Expected: []byte{ 0x31},
		},
		{
			Src:      []byte("0rm2"),
			Expected: []byte{ 0x32},
		},
		{
			Src:      []byte("0rmm"),
			Expected: []byte{ 0x33},
		},
		{
			Src:      []byte("0rm4"),
			Expected: []byte{ 0x34},
		},
		{
			Src:      []byte("0rm5"),
			Expected: []byte{ 0x35},
		},
		{
			Src:      []byte("0rmx"),
			Expected: []byte{ 0x36},
		},
		{
			Src:      []byte("0rmV"),
			Expected: []byte{ 0x37},
		},
		{
			Src:      []byte("0rmH"),
			Expected: []byte{ 0x38},
		},
		{
			Src:      []byte("0rmk"),
			Expected: []byte{ 0x39},
		},
		{
			Src:      []byte("0rmA"),
			Expected: []byte{ 0x3a},
		},
		{
			Src:      []byte("0rmb"),
			Expected: []byte{ 0x3b},
		},
		{
			Src:      []byte("0rmC"),
			Expected: []byte{ 0x3c},
		},
		{
			Src:      []byte("0rmd"),
			Expected: []byte{ 0x3d},
		},
		{
			Src:      []byte("0rmE"),
			Expected: []byte{ 0x3e},
		},
		{
			Src:      []byte("0rmf"),
			Expected: []byte{ 0x3f},
		},



		{
			Src:      []byte("0r40"),
			Expected: []byte{ 0x40},
		},
		{
			Src:      []byte("0r41"),
			Expected: []byte{ 0x41},
		},
		{
			Src:      []byte("0r42"),
			Expected: []byte{ 0x42},
		},
		{
			Src:      []byte("0r4m"),
			Expected: []byte{ 0x43},
		},
		{
			Src:      []byte("0r44"),
			Expected: []byte{ 0x44},
		},
		{
			Src:      []byte("0r45"),
			Expected: []byte{ 0x45},
		},
		{
			Src:      []byte("0r4x"),
			Expected: []byte{ 0x46},
		},
		{
			Src:      []byte("0r4V"),
			Expected: []byte{ 0x47},
		},
		{
			Src:      []byte("0r4H"),
			Expected: []byte{ 0x48},
		},
		{
			Src:      []byte("0r4k"),
			Expected: []byte{ 0x49},
		},
		{
			Src:      []byte("0r4A"),
			Expected: []byte{ 0x4a},
		},
		{
			Src:      []byte("0r4b"),
			Expected: []byte{ 0x4b},
		},
		{
			Src:      []byte("0r4C"),
			Expected: []byte{ 0x4c},
		},
		{
			Src:      []byte("0r4d"),
			Expected: []byte{ 0x4d},
		},
		{
			Src:      []byte("0r4E"),
			Expected: []byte{ 0x4e},
		},
		{
			Src:      []byte("0r4f"),
			Expected: []byte{ 0x4f},
		},



		{
			Src:      []byte("0r50"),
			Expected: []byte{ 0x50},
		},
		{
			Src:      []byte("0r51"),
			Expected: []byte{ 0x51},
		},
		{
			Src:      []byte("0r52"),
			Expected: []byte{ 0x52},
		},
		{
			Src:      []byte("0r5m"),
			Expected: []byte{ 0x53},
		},
		{
			Src:      []byte("0r54"),
			Expected: []byte{ 0x54},
		},
		{
			Src:      []byte("0r55"),
			Expected: []byte{ 0x55},
		},
		{
			Src:      []byte("0r5x"),
			Expected: []byte{ 0x56},
		},
		{
			Src:      []byte("0r5V"),
			Expected: []byte{ 0x57},
		},
		{
			Src:      []byte("0r5H"),
			Expected: []byte{ 0x58},
		},
		{
			Src:      []byte("0r5k"),
			Expected: []byte{ 0x59},
		},
		{
			Src:      []byte("0r5A"),
			Expected: []byte{ 0x5a},
		},
		{
			Src:      []byte("0r5b"),
			Expected: []byte{ 0x5b},
		},
		{
			Src:      []byte("0r5C"),
			Expected: []byte{ 0x5c},
		},
		{
			Src:      []byte("0r5d"),
			Expected: []byte{ 0x5d},
		},
		{
			Src:      []byte("0r5E"),
			Expected: []byte{ 0x5e},
		},
		{
			Src:      []byte("0r5f"),
			Expected: []byte{ 0x5f},
		},



		{
			Src:      []byte("0rx0"),
			Expected: []byte{ 0x60},
		},
		{
			Src:      []byte("0rx1"),
			Expected: []byte{ 0x61},
		},
		{
			Src:      []byte("0rx2"),
			Expected: []byte{ 0x62},
		},
		{
			Src:      []byte("0rxm"),
			Expected: []byte{ 0x63},
		},
		{
			Src:      []byte("0rx4"),
			Expected: []byte{ 0x64},
		},
		{
			Src:      []byte("0rx5"),
			Expected: []byte{ 0x65},
		},
		{
			Src:      []byte("0rxx"),
			Expected: []byte{ 0x66},
		},
		{
			Src:      []byte("0rxV"),
			Expected: []byte{ 0x67},
		},
		{
			Src:      []byte("0rxH"),
			Expected: []byte{ 0x68},
		},
		{
			Src:      []byte("0rxk"),
			Expected: []byte{ 0x69},
		},
		{
			Src:      []byte("0rxA"),
			Expected: []byte{ 0x6a},
		},
		{
			Src:      []byte("0rxb"),
			Expected: []byte{ 0x6b},
		},
		{
			Src:      []byte("0rxC"),
			Expected: []byte{ 0x6c},
		},
		{
			Src:      []byte("0rxd"),
			Expected: []byte{ 0x6d},
		},
		{
			Src:      []byte("0rxE"),
			Expected: []byte{ 0x6e},
		},
		{
			Src:      []byte("0rxf"),
			Expected: []byte{ 0x6f},
		},



		{
			Src:      []byte("0rV0"),
			Expected: []byte{ 0x70},
		},
		{
			Src:      []byte("0rV1"),
			Expected: []byte{ 0x71},
		},
		{
			Src:      []byte("0rV2"),
			Expected: []byte{ 0x72},
		},
		{
			Src:      []byte("0rVm"),
			Expected: []byte{ 0x73},
		},
		{
			Src:      []byte("0rV4"),
			Expected: []byte{ 0x74},
		},
		{
			Src:      []byte("0rV5"),
			Expected: []byte{ 0x75},
		},
		{
			Src:      []byte("0rVx"),
			Expected: []byte{ 0x76},
		},
		{
			Src:      []byte("0rVV"),
			Expected: []byte{ 0x77},
		},
		{
			Src:      []byte("0rVH"),
			Expected: []byte{ 0x78},
		},
		{
			Src:      []byte("0rVk"),
			Expected: []byte{ 0x79},
		},
		{
			Src:      []byte("0rVA"),
			Expected: []byte{ 0x7a},
		},
		{
			Src:      []byte("0rVb"),
			Expected: []byte{ 0x7b},
		},
		{
			Src:      []byte("0rVC"),
			Expected: []byte{ 0x7c},
		},
		{
			Src:      []byte("0rVd"),
			Expected: []byte{ 0x7d},
		},
		{
			Src:      []byte("0rVE"),
			Expected: []byte{ 0x7e},
		},
		{
			Src:      []byte("0rVf"),
			Expected: []byte{ 0x7f},
		},



		{
			Src:      []byte("0rH0"),
			Expected: []byte{ 0x80},
		},
		{
			Src:      []byte("0rH1"),
			Expected: []byte{ 0x81},
		},
		{
			Src:      []byte("0rH2"),
			Expected: []byte{ 0x82},
		},
		{
			Src:      []byte("0rHm"),
			Expected: []byte{ 0x83},
		},
		{
			Src:      []byte("0rH4"),
			Expected: []byte{ 0x84},
		},
		{
			Src:      []byte("0rH5"),
			Expected: []byte{ 0x85},
		},
		{
			Src:      []byte("0rHx"),
			Expected: []byte{ 0x86},
		},
		{
			Src:      []byte("0rHV"),
			Expected: []byte{ 0x87},
		},
		{
			Src:      []byte("0rHH"),
			Expected: []byte{ 0x88},
		},
		{
			Src:      []byte("0rHk"),
			Expected: []byte{ 0x89},
		},
		{
			Src:      []byte("0rHA"),
			Expected: []byte{ 0x8a},
		},
		{
			Src:      []byte("0rHb"),
			Expected: []byte{ 0x8b},
		},
		{
			Src:      []byte("0rHC"),
			Expected: []byte{ 0x8c},
		},
		{
			Src:      []byte("0rHd"),
			Expected: []byte{ 0x8d},
		},
		{
			Src:      []byte("0rHE"),
			Expected: []byte{ 0x8e},
		},
		{
			Src:      []byte("0rHf"),
			Expected: []byte{ 0x8f},
		},



		{
			Src:      []byte("0rk0"),
			Expected: []byte{ 0x90},
		},
		{
			Src:      []byte("0rk1"),
			Expected: []byte{ 0x91},
		},
		{
			Src:      []byte("0rk2"),
			Expected: []byte{ 0x92},
		},
		{
			Src:      []byte("0rkm"),
			Expected: []byte{ 0x93},
		},
		{
			Src:      []byte("0rk4"),
			Expected: []byte{ 0x94},
		},
		{
			Src:      []byte("0rk5"),
			Expected: []byte{ 0x95},
		},
		{
			Src:      []byte("0rkx"),
			Expected: []byte{ 0x96},
		},
		{
			Src:      []byte("0rkV"),
			Expected: []byte{ 0x97},
		},
		{
			Src:      []byte("0rkH"),
			Expected: []byte{ 0x98},
		},
		{
			Src:      []byte("0rkk"),
			Expected: []byte{ 0x99},
		},
		{
			Src:      []byte("0rkA"),
			Expected: []byte{ 0x9a},
		},
		{
			Src:      []byte("0rkb"),
			Expected: []byte{ 0x9b},
		},
		{
			Src:      []byte("0rkC"),
			Expected: []byte{ 0x9c},
		},
		{
			Src:      []byte("0rkd"),
			Expected: []byte{ 0x9d},
		},
		{
			Src:      []byte("0rkE"),
			Expected: []byte{ 0x9e},
		},
		{
			Src:      []byte("0rkf"),
			Expected: []byte{ 0x9f},
		},



		{
			Src:      []byte("0rA0"),
			Expected: []byte{ 0xa0},
		},
		{
			Src:      []byte("0rA1"),
			Expected: []byte{ 0xa1},
		},
		{
			Src:      []byte("0rA2"),
			Expected: []byte{ 0xa2},
		},
		{
			Src:      []byte("0rAm"),
			Expected: []byte{ 0xa3},
		},
		{
			Src:      []byte("0rA4"),
			Expected: []byte{ 0xa4},
		},
		{
			Src:      []byte("0rA5"),
			Expected: []byte{ 0xa5},
		},
		{
			Src:      []byte("0rAx"),
			Expected: []byte{ 0xa6},
		},
		{
			Src:      []byte("0rAV"),
			Expected: []byte{ 0xa7},
		},
		{
			Src:      []byte("0rAH"),
			Expected: []byte{ 0xa8},
		},
		{
			Src:      []byte("0rAk"),
			Expected: []byte{ 0xa9},
		},
		{
			Src:      []byte("0rAA"),
			Expected: []byte{ 0xaa},
		},
		{
			Src:      []byte("0rAb"),
			Expected: []byte{ 0xab},
		},
		{
			Src:      []byte("0rAC"),
			Expected: []byte{ 0xac},
		},
		{
			Src:      []byte("0rAd"),
			Expected: []byte{ 0xad},
		},
		{
			Src:      []byte("0rAE"),
			Expected: []byte{ 0xae},
		},
		{
			Src:      []byte("0rAf"),
			Expected: []byte{ 0xaf},
		},



		{
			Src:      []byte("0rb0"),
			Expected: []byte{ 0xb0},
		},
		{
			Src:      []byte("0rb1"),
			Expected: []byte{ 0xb1},
		},
		{
			Src:      []byte("0rb2"),
			Expected: []byte{ 0xb2},
		},
		{
			Src:      []byte("0rbm"),
			Expected: []byte{ 0xb3},
		},
		{
			Src:      []byte("0rb4"),
			Expected: []byte{ 0xb4},
		},
		{
			Src:      []byte("0rb5"),
			Expected: []byte{ 0xb5},
		},
		{
			Src:      []byte("0rbx"),
			Expected: []byte{ 0xb6},
		},
		{
			Src:      []byte("0rbV"),
			Expected: []byte{ 0xb7},
		},
		{
			Src:      []byte("0rbH"),
			Expected: []byte{ 0xb8},
		},
		{
			Src:      []byte("0rbk"),
			Expected: []byte{ 0xb9},
		},
		{
			Src:      []byte("0rbA"),
			Expected: []byte{ 0xba},
		},
		{
			Src:      []byte("0rbb"),
			Expected: []byte{ 0xbb},
		},
		{
			Src:      []byte("0rbC"),
			Expected: []byte{ 0xbc},
		},
		{
			Src:      []byte("0rbd"),
			Expected: []byte{ 0xbd},
		},
		{
			Src:      []byte("0rbE"),
			Expected: []byte{ 0xbe},
		},
		{
			Src:      []byte("0rbf"),
			Expected: []byte{ 0xbf},
		},



		{
			Src:      []byte("0rC0"),
			Expected: []byte{ 0xc0},
		},
		{
			Src:      []byte("0rC1"),
			Expected: []byte{ 0xc1},
		},
		{
			Src:      []byte("0rC2"),
			Expected: []byte{ 0xc2},
		},
		{
			Src:      []byte("0rCm"),
			Expected: []byte{ 0xc3},
		},
		{
			Src:      []byte("0rC4"),
			Expected: []byte{ 0xc4},
		},
		{
			Src:      []byte("0rC5"),
			Expected: []byte{ 0xc5},
		},
		{
			Src:      []byte("0rCx"),
			Expected: []byte{ 0xc6},
		},
		{
			Src:      []byte("0rCV"),
			Expected: []byte{ 0xc7},
		},
		{
			Src:      []byte("0rCH"),
			Expected: []byte{ 0xc8},
		},
		{
			Src:      []byte("0rCk"),
			Expected: []byte{ 0xc9},
		},
		{
			Src:      []byte("0rCA"),
			Expected: []byte{ 0xca},
		},
		{
			Src:      []byte("0rCb"),
			Expected: []byte{ 0xcb},
		},
		{
			Src:      []byte("0rCC"),
			Expected: []byte{ 0xcc},
		},
		{
			Src:      []byte("0rCd"),
			Expected: []byte{ 0xcd},
		},
		{
			Src:      []byte("0rCE"),
			Expected: []byte{ 0xce},
		},
		{
			Src:      []byte("0rCf"),
			Expected: []byte{ 0xcf},
		},



		{
			Src:      []byte("0rd0"),
			Expected: []byte{ 0xd0},
		},
		{
			Src:      []byte("0rd1"),
			Expected: []byte{ 0xd1},
		},
		{
			Src:      []byte("0rd2"),
			Expected: []byte{ 0xd2},
		},
		{
			Src:      []byte("0rdm"),
			Expected: []byte{ 0xd3},
		},
		{
			Src:      []byte("0rd4"),
			Expected: []byte{ 0xd4},
		},
		{
			Src:      []byte("0rd5"),
			Expected: []byte{ 0xd5},
		},
		{
			Src:      []byte("0rdx"),
			Expected: []byte{ 0xd6},
		},
		{
			Src:      []byte("0rdV"),
			Expected: []byte{ 0xd7},
		},
		{
			Src:      []byte("0rdH"),
			Expected: []byte{ 0xd8},
		},
		{
			Src:      []byte("0rdk"),
			Expected: []byte{ 0xd9},
		},
		{
			Src:      []byte("0rdA"),
			Expected: []byte{ 0xda},
		},
		{
			Src:      []byte("0rdb"),
			Expected: []byte{ 0xdb},
		},
		{
			Src:      []byte("0rdC"),
			Expected: []byte{ 0xdc},
		},
		{
			Src:      []byte("0rdd"),
			Expected: []byte{ 0xdd},
		},
		{
			Src:      []byte("0rdE"),
			Expected: []byte{ 0xde},
		},
		{
			Src:      []byte("0rdf"),
			Expected: []byte{ 0xdf},
		},



		{
			Src:      []byte("0rE0"),
			Expected: []byte{ 0xe0},
		},
		{
			Src:      []byte("0rE1"),
			Expected: []byte{ 0xe1},
		},
		{
			Src:      []byte("0rE2"),
			Expected: []byte{ 0xe2},
		},
		{
			Src:      []byte("0rEm"),
			Expected: []byte{ 0xe3},
		},
		{
			Src:      []byte("0rE4"),
			Expected: []byte{ 0xe4},
		},
		{
			Src:      []byte("0rE5"),
			Expected: []byte{ 0xe5},
		},
		{
			Src:      []byte("0rEx"),
			Expected: []byte{ 0xe6},
		},
		{
			Src:      []byte("0rEV"),
			Expected: []byte{ 0xe7},
		},
		{
			Src:      []byte("0rEH"),
			Expected: []byte{ 0xe8},
		},
		{
			Src:      []byte("0rEk"),
			Expected: []byte{ 0xe9},
		},
		{
			Src:      []byte("0rEA"),
			Expected: []byte{ 0xea},
		},
		{
			Src:      []byte("0rEb"),
			Expected: []byte{ 0xeb},
		},
		{
			Src:      []byte("0rEC"),
			Expected: []byte{ 0xec},
		},
		{
			Src:      []byte("0rEd"),
			Expected: []byte{ 0xed},
		},
		{
			Src:      []byte("0rEE"),
			Expected: []byte{ 0xee},
		},
		{
			Src:      []byte("0rEf"),
			Expected: []byte{ 0xef},
		},



		{
			Src:      []byte("0rf0"),
			Expected: []byte{ 0xf0},
		},
		{
			Src:      []byte("0rf1"),
			Expected: []byte{ 0xf1},
		},
		{
			Src:      []byte("0rf2"),
			Expected: []byte{ 0xf2},
		},
		{
			Src:      []byte("0rfm"),
			Expected: []byte{ 0xf3},
		},
		{
			Src:      []byte("0rf4"),
			Expected: []byte{ 0xf4},
		},
		{
			Src:      []byte("0rf5"),
			Expected: []byte{ 0xf5},
		},
		{
			Src:      []byte("0rfx"),
			Expected: []byte{ 0xf6},
		},
		{
			Src:      []byte("0rfV"),
			Expected: []byte{ 0xf7},
		},
		{
			Src:      []byte("0rfH"),
			Expected: []byte{ 0xf8},
		},
		{
			Src:      []byte("0rfk"),
			Expected: []byte{ 0xf9},
		},
		{
			Src:      []byte("0rfA"),
			Expected: []byte{ 0xfa},
		},
		{
			Src:      []byte("0rfb"),
			Expected: []byte{ 0xfb},
		},
		{
			Src:      []byte("0rfC"),
			Expected: []byte{ 0xfc},
		},
		{
			Src:      []byte("0rfd"),
			Expected: []byte{ 0xfd},
		},
		{
			Src:      []byte("0rfE"),
			Expected: []byte{ 0xfe},
		},
		{
			Src:      []byte("0rff"),
			Expected: []byte{ 0xff},
		},
	}

	for testNumber, test := range tests {
		var src []byte = test.Src

		dstLen := bravo16.DecodeLiteralLen(len(src))
		var dst []byte = make([]byte, dstLen)

		n, err := bravo16.DecodeLiteral(dst, src)
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
