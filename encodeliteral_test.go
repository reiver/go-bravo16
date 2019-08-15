package bravo16_test

import (
	"github.com/reiver/go-bravo16"

	"bytes"

	"testing"
)

func TestEncodeLiteral(t *testing.T) {

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
			Expected: []byte("0rx1V0V0xCx5"),
		},
		{
			Src:      []byte("BANANA"),
			Expected: []byte("0r42414E414E41"),
		},
		{
			Src:      []byte("Cherry"),
			Expected: []byte("0r4mxHx5V2V2Vk"),
		},
		{
			Src:      []byte("dATE"),
			Expected: []byte("0rx4415445"),
		},



		{
			Src:      []byte("Hello world!"),
			Expected: []byte("0r4Hx5xCxCxf20VVxfV2xCx421"),
		},



		{
			Src:      []byte{0xc0, 0x53, 0x5e, 0x4b, 0xe2, 0xb7, 0x9f, 0xfd, 0x93, 0x29, 0x13, 0x05, 0x43, 0x6b, 0xf8, 0x89, 0x31, 0x4e, 0x4a, 0x3f, 0xae, 0xc0, 0x5e, 0xcf, 0xfc, 0xbb, 0x7d, 0xf3, 0x1a, 0xd9, 0xe5, 0x1a},
			Expected: []byte("0rC05m5E4bE2bVkffdkm2k1m054mxbfHHkm14E4AmfAEC05ECffCbbVdfm1AdkE51A"),
		},



		{
			Src:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, 0x9c, 0x08, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1, 0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f},
			Expected: []byte("0r00000000001kdxxHkC0H5AE1x5Hm1Ekm4ffVxmAE4xA2AxC1V2bmf1bx0AHCE2xf"),
		},



		{
			Src:      []byte{ 0x00},
			Expected: []byte("0r00"),
		},
		{
			Src:      []byte{ 0x01},
			Expected: []byte("0r01"),
		},
		{
			Src:      []byte{ 0x02},
			Expected: []byte("0r02"),
		},
		{
			Src:      []byte{ 0x03},
			Expected: []byte("0r0m"),
		},
		{
			Src:      []byte{ 0x04},
			Expected: []byte("0r04"),
		},
		{
			Src:      []byte{ 0x05},
			Expected: []byte("0r05"),
		},
		{
			Src:      []byte{ 0x06},
			Expected: []byte("0r0x"),
		},
		{
			Src:      []byte{ 0x07},
			Expected: []byte("0r0V"),
		},
		{
			Src:      []byte{ 0x08},
			Expected: []byte("0r0H"),
		},
		{
			Src:      []byte{ 0x09},
			Expected: []byte("0r0k"),
		},
		{
			Src:      []byte{ 0x0a},
			Expected: []byte("0r0A"),
		},
		{
			Src:      []byte{ 0x0b},
			Expected: []byte("0r0b"),
		},
		{
			Src:      []byte{ 0x0c},
			Expected: []byte("0r0C"),
		},
		{
			Src:      []byte{ 0x0d},
			Expected: []byte("0r0d"),
		},
		{
			Src:      []byte{ 0x0e},
			Expected: []byte("0r0E"),
		},
		{
			Src:      []byte{ 0x0f},
			Expected: []byte("0r0f"),
		},



		{
			Src:      []byte{ 0x10},
			Expected: []byte("0r10"),
		},
		{
			Src:      []byte{ 0x11},
			Expected: []byte("0r11"),
		},
		{
			Src:      []byte{ 0x12},
			Expected: []byte("0r12"),
		},
		{
			Src:      []byte{ 0x13},
			Expected: []byte("0r1m"),
		},
		{
			Src:      []byte{ 0x14},
			Expected: []byte("0r14"),
		},
		{
			Src:      []byte{ 0x15},
			Expected: []byte("0r15"),
		},
		{
			Src:      []byte{ 0x16},
			Expected: []byte("0r1x"),
		},
		{
			Src:      []byte{ 0x17},
			Expected: []byte("0r1V"),
		},
		{
			Src:      []byte{ 0x18},
			Expected: []byte("0r1H"),
		},
		{
			Src:      []byte{ 0x19},
			Expected: []byte("0r1k"),
		},
		{
			Src:      []byte{ 0x1a},
			Expected: []byte("0r1A"),
		},
		{
			Src:      []byte{ 0x1b},
			Expected: []byte("0r1b"),
		},
		{
			Src:      []byte{ 0x1c},
			Expected: []byte("0r1C"),
		},
		{
			Src:      []byte{ 0x1d},
			Expected: []byte("0r1d"),
		},
		{
			Src:      []byte{ 0x1e},
			Expected: []byte("0r1E"),
		},
		{
			Src:      []byte{ 0x1f},
			Expected: []byte("0r1f"),
		},



		{
			Src:      []byte{ 0x20},
			Expected: []byte("0r20"),
		},
		{
			Src:      []byte{ 0x21},
			Expected: []byte("0r21"),
		},
		{
			Src:      []byte{ 0x22},
			Expected: []byte("0r22"),
		},
		{
			Src:      []byte{ 0x23},
			Expected: []byte("0r2m"),
		},
		{
			Src:      []byte{ 0x24},
			Expected: []byte("0r24"),
		},
		{
			Src:      []byte{ 0x25},
			Expected: []byte("0r25"),
		},
		{
			Src:      []byte{ 0x26},
			Expected: []byte("0r2x"),
		},
		{
			Src:      []byte{ 0x27},
			Expected: []byte("0r2V"),
		},
		{
			Src:      []byte{ 0x28},
			Expected: []byte("0r2H"),
		},
		{
			Src:      []byte{ 0x29},
			Expected: []byte("0r2k"),
		},
		{
			Src:      []byte{ 0x2a},
			Expected: []byte("0r2A"),
		},
		{
			Src:      []byte{ 0x2b},
			Expected: []byte("0r2b"),
		},
		{
			Src:      []byte{ 0x2c},
			Expected: []byte("0r2C"),
		},
		{
			Src:      []byte{ 0x2d},
			Expected: []byte("0r2d"),
		},
		{
			Src:      []byte{ 0x2e},
			Expected: []byte("0r2E"),
		},
		{
			Src:      []byte{ 0x2f},
			Expected: []byte("0r2f"),
		},



		{
			Src:      []byte{ 0x30},
			Expected: []byte("0rm0"),
		},
		{
			Src:      []byte{ 0x31},
			Expected: []byte("0rm1"),
		},
		{
			Src:      []byte{ 0x32},
			Expected: []byte("0rm2"),
		},
		{
			Src:      []byte{ 0x33},
			Expected: []byte("0rmm"),
		},
		{
			Src:      []byte{ 0x34},
			Expected: []byte("0rm4"),
		},
		{
			Src:      []byte{ 0x35},
			Expected: []byte("0rm5"),
		},
		{
			Src:      []byte{ 0x36},
			Expected: []byte("0rmx"),
		},
		{
			Src:      []byte{ 0x37},
			Expected: []byte("0rmV"),
		},
		{
			Src:      []byte{ 0x38},
			Expected: []byte("0rmH"),
		},
		{
			Src:      []byte{ 0x39},
			Expected: []byte("0rmk"),
		},
		{
			Src:      []byte{ 0x3a},
			Expected: []byte("0rmA"),
		},
		{
			Src:      []byte{ 0x3b},
			Expected: []byte("0rmb"),
		},
		{
			Src:      []byte{ 0x3c},
			Expected: []byte("0rmC"),
		},
		{
			Src:      []byte{ 0x3d},
			Expected: []byte("0rmd"),
		},
		{
			Src:      []byte{ 0x3e},
			Expected: []byte("0rmE"),
		},
		{
			Src:      []byte{ 0x3f},
			Expected: []byte("0rmf"),
		},



		{
			Src:      []byte{ 0x40},
			Expected: []byte("0r40"),
		},
		{
			Src:      []byte{ 0x41},
			Expected: []byte("0r41"),
		},
		{
			Src:      []byte{ 0x42},
			Expected: []byte("0r42"),
		},
		{
			Src:      []byte{ 0x43},
			Expected: []byte("0r4m"),
		},
		{
			Src:      []byte{ 0x44},
			Expected: []byte("0r44"),
		},
		{
			Src:      []byte{ 0x45},
			Expected: []byte("0r45"),
		},
		{
			Src:      []byte{ 0x46},
			Expected: []byte("0r4x"),
		},
		{
			Src:      []byte{ 0x47},
			Expected: []byte("0r4V"),
		},
		{
			Src:      []byte{ 0x48},
			Expected: []byte("0r4H"),
		},
		{
			Src:      []byte{ 0x49},
			Expected: []byte("0r4k"),
		},
		{
			Src:      []byte{ 0x4a},
			Expected: []byte("0r4A"),
		},
		{
			Src:      []byte{ 0x4b},
			Expected: []byte("0r4b"),
		},
		{
			Src:      []byte{ 0x4c},
			Expected: []byte("0r4C"),
		},
		{
			Src:      []byte{ 0x4d},
			Expected: []byte("0r4d"),
		},
		{
			Src:      []byte{ 0x4e},
			Expected: []byte("0r4E"),
		},
		{
			Src:      []byte{ 0x4f},
			Expected: []byte("0r4f"),
		},



		{
			Src:      []byte{ 0x50},
			Expected: []byte("0r50"),
		},
		{
			Src:      []byte{ 0x51},
			Expected: []byte("0r51"),
		},
		{
			Src:      []byte{ 0x52},
			Expected: []byte("0r52"),
		},
		{
			Src:      []byte{ 0x53},
			Expected: []byte("0r5m"),
		},
		{
			Src:      []byte{ 0x54},
			Expected: []byte("0r54"),
		},
		{
			Src:      []byte{ 0x55},
			Expected: []byte("0r55"),
		},
		{
			Src:      []byte{ 0x56},
			Expected: []byte("0r5x"),
		},
		{
			Src:      []byte{ 0x57},
			Expected: []byte("0r5V"),
		},
		{
			Src:      []byte{ 0x58},
			Expected: []byte("0r5H"),
		},
		{
			Src:      []byte{ 0x59},
			Expected: []byte("0r5k"),
		},
		{
			Src:      []byte{ 0x5a},
			Expected: []byte("0r5A"),
		},
		{
			Src:      []byte{ 0x5b},
			Expected: []byte("0r5b"),
		},
		{
			Src:      []byte{ 0x5c},
			Expected: []byte("0r5C"),
		},
		{
			Src:      []byte{ 0x5d},
			Expected: []byte("0r5d"),
		},
		{
			Src:      []byte{ 0x5e},
			Expected: []byte("0r5E"),
		},
		{
			Src:      []byte{ 0x5f},
			Expected: []byte("0r5f"),
		},



		{
			Src:      []byte{ 0x60},
			Expected: []byte("0rx0"),
		},
		{
			Src:      []byte{ 0x61},
			Expected: []byte("0rx1"),
		},
		{
			Src:      []byte{ 0x62},
			Expected: []byte("0rx2"),
		},
		{
			Src:      []byte{ 0x63},
			Expected: []byte("0rxm"),
		},
		{
			Src:      []byte{ 0x64},
			Expected: []byte("0rx4"),
		},
		{
			Src:      []byte{ 0x65},
			Expected: []byte("0rx5"),
		},
		{
			Src:      []byte{ 0x66},
			Expected: []byte("0rxx"),
		},
		{
			Src:      []byte{ 0x67},
			Expected: []byte("0rxV"),
		},
		{
			Src:      []byte{ 0x68},
			Expected: []byte("0rxH"),
		},
		{
			Src:      []byte{ 0x69},
			Expected: []byte("0rxk"),
		},
		{
			Src:      []byte{ 0x6a},
			Expected: []byte("0rxA"),
		},
		{
			Src:      []byte{ 0x6b},
			Expected: []byte("0rxb"),
		},
		{
			Src:      []byte{ 0x6c},
			Expected: []byte("0rxC"),
		},
		{
			Src:      []byte{ 0x6d},
			Expected: []byte("0rxd"),
		},
		{
			Src:      []byte{ 0x6e},
			Expected: []byte("0rxE"),
		},
		{
			Src:      []byte{ 0x6f},
			Expected: []byte("0rxf"),
		},



		{
			Src:      []byte{ 0x70},
			Expected: []byte("0rV0"),
		},
		{
			Src:      []byte{ 0x71},
			Expected: []byte("0rV1"),
		},
		{
			Src:      []byte{ 0x72},
			Expected: []byte("0rV2"),
		},
		{
			Src:      []byte{ 0x73},
			Expected: []byte("0rVm"),
		},
		{
			Src:      []byte{ 0x74},
			Expected: []byte("0rV4"),
		},
		{
			Src:      []byte{ 0x75},
			Expected: []byte("0rV5"),
		},
		{
			Src:      []byte{ 0x76},
			Expected: []byte("0rVx"),
		},
		{
			Src:      []byte{ 0x77},
			Expected: []byte("0rVV"),
		},
		{
			Src:      []byte{ 0x78},
			Expected: []byte("0rVH"),
		},
		{
			Src:      []byte{ 0x79},
			Expected: []byte("0rVk"),
		},
		{
			Src:      []byte{ 0x7a},
			Expected: []byte("0rVA"),
		},
		{
			Src:      []byte{ 0x7b},
			Expected: []byte("0rVb"),
		},
		{
			Src:      []byte{ 0x7c},
			Expected: []byte("0rVC"),
		},
		{
			Src:      []byte{ 0x7d},
			Expected: []byte("0rVd"),
		},
		{
			Src:      []byte{ 0x7e},
			Expected: []byte("0rVE"),
		},
		{
			Src:      []byte{ 0x7f},
			Expected: []byte("0rVf"),
		},



		{
			Src:      []byte{ 0x80},
			Expected: []byte("0rH0"),
		},
		{
			Src:      []byte{ 0x81},
			Expected: []byte("0rH1"),
		},
		{
			Src:      []byte{ 0x82},
			Expected: []byte("0rH2"),
		},
		{
			Src:      []byte{ 0x83},
			Expected: []byte("0rHm"),
		},
		{
			Src:      []byte{ 0x84},
			Expected: []byte("0rH4"),
		},
		{
			Src:      []byte{ 0x85},
			Expected: []byte("0rH5"),
		},
		{
			Src:      []byte{ 0x86},
			Expected: []byte("0rHx"),
		},
		{
			Src:      []byte{ 0x87},
			Expected: []byte("0rHV"),
		},
		{
			Src:      []byte{ 0x88},
			Expected: []byte("0rHH"),
		},
		{
			Src:      []byte{ 0x89},
			Expected: []byte("0rHk"),
		},
		{
			Src:      []byte{ 0x8a},
			Expected: []byte("0rHA"),
		},
		{
			Src:      []byte{ 0x8b},
			Expected: []byte("0rHb"),
		},
		{
			Src:      []byte{ 0x8c},
			Expected: []byte("0rHC"),
		},
		{
			Src:      []byte{ 0x8d},
			Expected: []byte("0rHd"),
		},
		{
			Src:      []byte{ 0x8e},
			Expected: []byte("0rHE"),
		},
		{
			Src:      []byte{ 0x8f},
			Expected: []byte("0rHf"),
		},



		{
			Src:      []byte{ 0x90},
			Expected: []byte("0rk0"),
		},
		{
			Src:      []byte{ 0x91},
			Expected: []byte("0rk1"),
		},
		{
			Src:      []byte{ 0x92},
			Expected: []byte("0rk2"),
		},
		{
			Src:      []byte{ 0x93},
			Expected: []byte("0rkm"),
		},
		{
			Src:      []byte{ 0x94},
			Expected: []byte("0rk4"),
		},
		{
			Src:      []byte{ 0x95},
			Expected: []byte("0rk5"),
		},
		{
			Src:      []byte{ 0x96},
			Expected: []byte("0rkx"),
		},
		{
			Src:      []byte{ 0x97},
			Expected: []byte("0rkV"),
		},
		{
			Src:      []byte{ 0x98},
			Expected: []byte("0rkH"),
		},
		{
			Src:      []byte{ 0x99},
			Expected: []byte("0rkk"),
		},
		{
			Src:      []byte{ 0x9a},
			Expected: []byte("0rkA"),
		},
		{
			Src:      []byte{ 0x9b},
			Expected: []byte("0rkb"),
		},
		{
			Src:      []byte{ 0x9c},
			Expected: []byte("0rkC"),
		},
		{
			Src:      []byte{ 0x9d},
			Expected: []byte("0rkd"),
		},
		{
			Src:      []byte{ 0x9e},
			Expected: []byte("0rkE"),
		},
		{
			Src:      []byte{ 0x9f},
			Expected: []byte("0rkf"),
		},



		{
			Src:      []byte{ 0xa0},
			Expected: []byte("0rA0"),
		},
		{
			Src:      []byte{ 0xa1},
			Expected: []byte("0rA1"),
		},
		{
			Src:      []byte{ 0xa2},
			Expected: []byte("0rA2"),
		},
		{
			Src:      []byte{ 0xa3},
			Expected: []byte("0rAm"),
		},
		{
			Src:      []byte{ 0xa4},
			Expected: []byte("0rA4"),
		},
		{
			Src:      []byte{ 0xa5},
			Expected: []byte("0rA5"),
		},
		{
			Src:      []byte{ 0xa6},
			Expected: []byte("0rAx"),
		},
		{
			Src:      []byte{ 0xa7},
			Expected: []byte("0rAV"),
		},
		{
			Src:      []byte{ 0xa8},
			Expected: []byte("0rAH"),
		},
		{
			Src:      []byte{ 0xa9},
			Expected: []byte("0rAk"),
		},
		{
			Src:      []byte{ 0xaa},
			Expected: []byte("0rAA"),
		},
		{
			Src:      []byte{ 0xab},
			Expected: []byte("0rAb"),
		},
		{
			Src:      []byte{ 0xac},
			Expected: []byte("0rAC"),
		},
		{
			Src:      []byte{ 0xad},
			Expected: []byte("0rAd"),
		},
		{
			Src:      []byte{ 0xae},
			Expected: []byte("0rAE"),
		},
		{
			Src:      []byte{ 0xaf},
			Expected: []byte("0rAf"),
		},




		{
			Src:      []byte{ 0xb0},
			Expected: []byte("0rb0"),
		},
		{
			Src:      []byte{ 0xb1},
			Expected: []byte("0rb1"),
		},
		{
			Src:      []byte{ 0xb2},
			Expected: []byte("0rb2"),
		},
		{
			Src:      []byte{ 0xb3},
			Expected: []byte("0rbm"),
		},
		{
			Src:      []byte{ 0xb4},
			Expected: []byte("0rb4"),
		},
		{
			Src:      []byte{ 0xb5},
			Expected: []byte("0rb5"),
		},
		{
			Src:      []byte{ 0xb6},
			Expected: []byte("0rbx"),
		},
		{
			Src:      []byte{ 0xb7},
			Expected: []byte("0rbV"),
		},
		{
			Src:      []byte{ 0xb8},
			Expected: []byte("0rbH"),
		},
		{
			Src:      []byte{ 0xb9},
			Expected: []byte("0rbk"),
		},
		{
			Src:      []byte{ 0xba},
			Expected: []byte("0rbA"),
		},
		{
			Src:      []byte{ 0xbb},
			Expected: []byte("0rbb"),
		},
		{
			Src:      []byte{ 0xbc},
			Expected: []byte("0rbC"),
		},
		{
			Src:      []byte{ 0xbd},
			Expected: []byte("0rbd"),
		},
		{
			Src:      []byte{ 0xbe},
			Expected: []byte("0rbE"),
		},
		{
			Src:      []byte{ 0xbf},
			Expected: []byte("0rbf"),
		},



		{
			Src:      []byte{ 0xc0},
			Expected: []byte("0rC0"),
		},
		{
			Src:      []byte{ 0xc1},
			Expected: []byte("0rC1"),
		},
		{
			Src:      []byte{ 0xc2},
			Expected: []byte("0rC2"),
		},
		{
			Src:      []byte{ 0xc3},
			Expected: []byte("0rCm"),
		},
		{
			Src:      []byte{ 0xc4},
			Expected: []byte("0rC4"),
		},
		{
			Src:      []byte{ 0xc5},
			Expected: []byte("0rC5"),
		},
		{
			Src:      []byte{ 0xc6},
			Expected: []byte("0rCx"),
		},
		{
			Src:      []byte{ 0xc7},
			Expected: []byte("0rCV"),
		},
		{
			Src:      []byte{ 0xc8},
			Expected: []byte("0rCH"),
		},
		{
			Src:      []byte{ 0xc9},
			Expected: []byte("0rCk"),
		},
		{
			Src:      []byte{ 0xca},
			Expected: []byte("0rCA"),
		},
		{
			Src:      []byte{ 0xcb},
			Expected: []byte("0rCb"),
		},
		{
			Src:      []byte{ 0xcc},
			Expected: []byte("0rCC"),
		},
		{
			Src:      []byte{ 0xcd},
			Expected: []byte("0rCd"),
		},
		{
			Src:      []byte{ 0xce},
			Expected: []byte("0rCE"),
		},
		{
			Src:      []byte{ 0xcf},
			Expected: []byte("0rCf"),
		},



		{
			Src:      []byte{ 0xd0},
			Expected: []byte("0rd0"),
		},
		{
			Src:      []byte{ 0xd1},
			Expected: []byte("0rd1"),
		},
		{
			Src:      []byte{ 0xd2},
			Expected: []byte("0rd2"),
		},
		{
			Src:      []byte{ 0xd3},
			Expected: []byte("0rdm"),
		},
		{
			Src:      []byte{ 0xd4},
			Expected: []byte("0rd4"),
		},
		{
			Src:      []byte{ 0xd5},
			Expected: []byte("0rd5"),
		},
		{
			Src:      []byte{ 0xd6},
			Expected: []byte("0rdx"),
		},
		{
			Src:      []byte{ 0xd7},
			Expected: []byte("0rdV"),
		},
		{
			Src:      []byte{ 0xd8},
			Expected: []byte("0rdH"),
		},
		{
			Src:      []byte{ 0xd9},
			Expected: []byte("0rdk"),
		},
		{
			Src:      []byte{ 0xda},
			Expected: []byte("0rdA"),
		},
		{
			Src:      []byte{ 0xdb},
			Expected: []byte("0rdb"),
		},
		{
			Src:      []byte{ 0xdc},
			Expected: []byte("0rdC"),
		},
		{
			Src:      []byte{ 0xdd},
			Expected: []byte("0rdd"),
		},
		{
			Src:      []byte{ 0xde},
			Expected: []byte("0rdE"),
		},
		{
			Src:      []byte{ 0xdf},
			Expected: []byte("0rdf"),
		},



		{
			Src:      []byte{ 0xe0},
			Expected: []byte("0rE0"),
		},
		{
			Src:      []byte{ 0xe1},
			Expected: []byte("0rE1"),
		},
		{
			Src:      []byte{ 0xe2},
			Expected: []byte("0rE2"),
		},
		{
			Src:      []byte{ 0xe3},
			Expected: []byte("0rEm"),
		},
		{
			Src:      []byte{ 0xe4},
			Expected: []byte("0rE4"),
		},
		{
			Src:      []byte{ 0xe5},
			Expected: []byte("0rE5"),
		},
		{
			Src:      []byte{ 0xe6},
			Expected: []byte("0rEx"),
		},
		{
			Src:      []byte{ 0xe7},
			Expected: []byte("0rEV"),
		},
		{
			Src:      []byte{ 0xe8},
			Expected: []byte("0rEH"),
		},
		{
			Src:      []byte{ 0xe9},
			Expected: []byte("0rEk"),
		},
		{
			Src:      []byte{ 0xea},
			Expected: []byte("0rEA"),
		},
		{
			Src:      []byte{ 0xeb},
			Expected: []byte("0rEb"),
		},
		{
			Src:      []byte{ 0xec},
			Expected: []byte("0rEC"),
		},
		{
			Src:      []byte{ 0xed},
			Expected: []byte("0rEd"),
		},
		{
			Src:      []byte{ 0xee},
			Expected: []byte("0rEE"),
		},
		{
			Src:      []byte{ 0xef},
			Expected: []byte("0rEf"),
		},



		{
			Src:      []byte{ 0xf0},
			Expected: []byte("0rf0"),
		},
		{
			Src:      []byte{ 0xf1},
			Expected: []byte("0rf1"),
		},
		{
			Src:      []byte{ 0xf2},
			Expected: []byte("0rf2"),
		},
		{
			Src:      []byte{ 0xf3},
			Expected: []byte("0rfm"),
		},
		{
			Src:      []byte{ 0xf4},
			Expected: []byte("0rf4"),
		},
		{
			Src:      []byte{ 0xf5},
			Expected: []byte("0rf5"),
		},
		{
			Src:      []byte{ 0xf6},
			Expected: []byte("0rfx"),
		},
		{
			Src:      []byte{ 0xf7},
			Expected: []byte("0rfV"),
		},
		{
			Src:      []byte{ 0xf8},
			Expected: []byte("0rfH"),
		},
		{
			Src:      []byte{ 0xf9},
			Expected: []byte("0rfk"),
		},
		{
			Src:      []byte{ 0xfa},
			Expected: []byte("0rfA"),
		},
		{
			Src:      []byte{ 0xfb},
			Expected: []byte("0rfb"),
		},
		{
			Src:      []byte{ 0xfc},
			Expected: []byte("0rfC"),
		},
		{
			Src:      []byte{ 0xfd},
			Expected: []byte("0rfd"),
		},
		{
			Src:      []byte{ 0xfe},
			Expected: []byte("0rfE"),
		},
		{
			Src:      []byte{ 0xff},
			Expected: []byte("0rff"),
		},
	}

	for testNumber, test := range tests {

		dstLen := bravo16.EncodeLiteralLen(len(test.Src))

		var dst []byte = make([]byte, dstLen)

		n, err := bravo16.EncodeLiteral(dst, test.Src)
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
			t.Logf("Dst (string):  %q", dst)
			t.Logf("Dst (bytes):  %#v", dst)
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
