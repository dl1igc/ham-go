package wsjtx_test

import (
	"testing"

	"github.com/tzneal/ham-go/adif"
	"github.com/tzneal/ham-go/wsjtx"
)

/*
[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xf3,0x3f,0xc9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x03,0x45,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0f,0x4e,0x38,0x4e,0x57,0x20,0x41,0x42,0x31,0x41,0x58,0x20,0x46,0x4d,0x31,0x37,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xfc,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x03,0x90,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x10,0x4c,0x5a,0x31,0x57,0x56,0x20,0x4e,0x35,0x4c,0x44,0x4f,0x20,0x45,0x4d,0x31,0x33,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xf2,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x04,0x01,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0c,0x43,0x51,0x20,0x57,0x42,0x34,0x58,0x20,0x46,0x4d,0x30,0x36,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xf9,0x3f,0xf4,0xcc,0xcc,0xc0,0x00,0x00,0x00,0x00,0x00,0x04,0x45,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0d,0x43,0x51,0x20,0x4b,0x57,0x34,0x53,0x50,0x20,0x45,0x4d,0x36,0x34,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xe8,0xbf,0xe6,0x66,0x66,0x60,0x00,0x00,0x00,0x00,0x00,0x04,0x91,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0e,0x43,0x51,0x20,0x4b,0x43,0x38,0x4f,0x45,0x4d,0x20,0x45,0x4e,0x39,0x31,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xed,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x04,0xcf,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0f,0x4c,0x5a,0x31,0x50,0x52,0x58,0x20,0x57,0x31,0x4f,0x50,0x20,0x2d,0x31,0x33,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xfe,0x3f,0xd9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x05,0xcd,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0d,0x43,0x51,0x20,0x41,0x43,0x31,0x4d,0x58,0x20,0x46,0x4e,0x33,0x34,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xf6,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x06,0xa0,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0c,0x43,0x51,0x20,0x4e,0x31,0x55,0x4c,0x20,0x45,0x4c,0x39,0x35,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xf6,0x3f,0xd3,0x33,0x33,0x40,0x00,0x00,0x00,0x00,0x00,0x07,0x23,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x11,0x4b,0x56,0x34,0x54,0x54,0x20,0x57,0x42,0x33,0x4a,0x55,0x56,0x20,0x46,0x4d,0x32,0x39,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0x00,0x00,0x00,0x00,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x07,0xe1,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x10,0x4f,0x44,0x35,0x5a,0x5a,0x20,0x41,0x42,0x35,0x56,0x4a,0x20,0x46,0x4d,0x31,0x39,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xee,0x3f,0xc9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x08,0x34,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0d,0x43,0x51,0x20,0x4e,0x32,0x53,0x49,0x54,0x20,0x46,0x4e,0x30,0x32,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xed,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x08,0x7a,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x11,0x56,0x50,0x32,0x45,0x54,0x45,0x20,0x4b,0x31,0x54,0x53,0x54,0x20,0x46,0x4e,0x34,0x32,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0x00,0x00,0x00,0x07,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x08,0xbe,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x10,0x4e,0x36,0x4f,0x50,0x52,0x20,0x4e,0x35,0x41,0x41,0x50,0x20,0x45,0x4d,0x31,0x35,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0x00,0x00,0x00,0x03,0x3f,0xf6,0x66,0x66,0x60,0x00,0x00,0x00,0x00,0x00,0x01,0x51,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0d,0x43,0x51,0x20,0x4e,0x34,0x50,0x41,0x5a,0x20,0x45,0x4c,0x39,0x35,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xf7,0x3f,0xd3,0x33,0x33,0x40,0x00,0x00,0x00,0x00,0x00,0x00,0xf3,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0d,0x43,0x51,0x20,0x4b,0x39,0x5a,0x49,0x45,0x20,0x45,0x4e,0x35,0x34,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x00,0x00,0x00,0x03,0x00,0x00,0x00,0x05,0x32,0x2e,0x31,0x2e,0x32,0x00,0x00,0x00,0x06,0x30,0x30,0x36,0x38,0x66,0x39}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0x00,0x00,0x00,0x07,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x01,0x97,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x10,0x55,0x58,0x36,0x49,0x52,0x20,0x57,0x33,0x4b,0x42,0x47,0x20,0x46,0x4e,0x31,0x30,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xf9,0x3f,0xc9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x03,0xac,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x11,0x3c,0x34,0x41,0x36,0x30,0x43,0x3e,0x20,0x57,0x38,0x55,0x56,0x20,0x46,0x4d,0x32,0x38,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xf9,0x3f,0xd3,0x33,0x33,0x40,0x00,0x00,0x00,0x00,0x00,0x04,0x41,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0e,0x58,0x45,0x33,0x52,0x20,0x4b,0x36,0x4a,0x57,0x20,0x44,0x4d,0x30,0x33,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xe9,0xbf,0xd3,0x33,0x33,0x40,0x00,0x00,0x00,0x00,0x00,0x06,0x22,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x11,0x55,0x57,0x35,0x5a,0x4d,0x20,0x4b,0x4d,0x34,0x49,0x59,0x4f,0x20,0x52,0x2d,0x31,0x38,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xe8,0x3f,0xb9,0x99,0x99,0xa0,0x00,0x00,0x00,0x00,0x00,0x06,0xb1,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x10,0x4b,0x4e,0x36,0x45,0x57,0x4c,0x20,0x4b,0x43,0x31,0x45,0x4c,0x46,0x20,0x37,0x33,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0xff,0xff,0xff,0xeb,0x3f,0xf0,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x06,0xfe,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0e,0x43,0x51,0x20,0x43,0x4d,0x32,0x4a,0x47,0x4d,0x20,0x45,0x4c,0x38,0x33,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x01,0x00,0xa2,0xf6,0xc0,0x00,0x00,0x00,0x07,0x3f,0xd3,0x33,0x33,0x40,0x00,0x00,0x00,0x00,0x00,0x01,0x8a,0x00,0x00,0x00,0x01,0x7e,0x00,0x00,0x00,0x0f,0x55,0x58,0x31,0x56,0x54,0x20,0x4b,0x33,0x45,0x41,0x20,0x46,0x4e,0x32,0x30,0x00,0x00}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58,0x00,0x00,0x00,0x00,0x00,0x6b,0xf0,0xd0,0x00,0x00,0x00,0x03,0x46,0x54,0x38,0x00,0x00,0x00,0x05,0x4e,0x34,0x50,0x41,0x5a,0x00,0x00,0x00,0x02,0x2d,0x34,0x00,0x00,0x00,0x03,0x46,0x54,0x38,0x00,0x00,0x00,0x00,0x00,0x01,0x51,0x00,0x00,0x01,0x51,0x00,0x00,0x00,0x05,0x57,0x34,0x54,0x4e,0x4c,0x00,0x00,0x00,0x06,0x45,0x4d,0x36,0x34,0x4f,0x52,0x00,0x00,0x00,0x04,0x45,0x4c,0x39,0x35,0x00,0xff,0xff,0xff,0xff,0x00,0x00,0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff,0x00,0x00,0x00,0x07,0x44,0x65,0x66,0x61,0x75,0x6c,0x74}

[]byte{0xad,0xbc,0xcb,0xda,0x00,0x00,0x00,0x02,0x00,0x00,0x00,0x06,0x00,0x00,0x00,0x06,0x57,0x53,0x4a,0x54,0x2d,0x58}

*/
func TestDecodeLoggedADIF(t *testing.T) {
	msg, err := wsjtx.Decode([]byte{0xad, 0xbc, 0xcb, 0xda, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x06, 0x57, 0x53, 0x4a, 0x54, 0x2d, 0x58, 0x00, 0x00, 0x01, 0x31, 0x0a, 0x3c, 0x61, 0x64, 0x69, 0x66, 0x5f, 0x76, 0x65, 0x72, 0x3a, 0x35, 0x3e, 0x33, 0x2e, 0x31, 0x2e, 0x30, 0x0a, 0x3c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x69, 0x64, 0x3a, 0x36, 0x3e, 0x57, 0x53, 0x4a, 0x54, 0x2d, 0x58, 0x0a, 0x3c, 0x45, 0x4f, 0x48, 0x3e, 0x0a, 0x3c, 0x63, 0x61, 0x6c, 0x6c, 0x3a, 0x36, 0x3e, 0x57, 0x42, 0x36, 0x46, 0x57, 0x53, 0x20, 0x3c, 0x67, 0x72, 0x69, 0x64, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x3a, 0x34, 0x3e, 0x44, 0x4d, 0x31, 0x32, 0x20, 0x3c, 0x6d, 0x6f, 0x64, 0x65, 0x3a, 0x33, 0x3e, 0x46, 0x54, 0x38, 0x20, 0x3c, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x3a, 0x33, 0x3e, 0x2b, 0x30, 0x37, 0x20, 0x3c, 0x72, 0x73, 0x74, 0x5f, 0x72, 0x63, 0x76, 0x64, 0x3a, 0x33, 0x3e, 0x2d, 0x30, 0x32, 0x20, 0x3c, 0x71, 0x73, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x38, 0x3e, 0x32, 0x30, 0x32, 0x30, 0x30, 0x34, 0x32, 0x33, 0x20, 0x3c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x6e, 0x3a, 0x36, 0x3e, 0x30, 0x32, 0x31, 0x35, 0x34, 0x35, 0x20, 0x3c, 0x71, 0x73, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x3a, 0x38, 0x3e, 0x32, 0x30, 0x32, 0x30, 0x30, 0x34, 0x32, 0x33, 0x20, 0x3c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x3a, 0x36, 0x3e, 0x30, 0x32, 0x31, 0x36, 0x34, 0x35, 0x20, 0x3c, 0x62, 0x61, 0x6e, 0x64, 0x3a, 0x33, 0x3e, 0x32, 0x30, 0x6d, 0x20, 0x3c, 0x66, 0x72, 0x65, 0x71, 0x3a, 0x39, 0x3e, 0x31, 0x34, 0x2e, 0x30, 0x37, 0x36, 0x39, 0x33, 0x37, 0x20, 0x3c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x3a, 0x35, 0x3e, 0x57, 0x34, 0x54, 0x4e, 0x4c, 0x20, 0x3c, 0x6d, 0x79, 0x5f, 0x67, 0x72, 0x69, 0x64, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x3a, 0x36, 0x3e, 0x45, 0x4d, 0x36, 0x34, 0x4f, 0x52, 0x20, 0x3c, 0x74, 0x78, 0x5f, 0x70, 0x77, 0x72, 0x3a, 0x33, 0x3e, 0x32, 0x35, 0x57, 0x20, 0x3c, 0x45, 0x4f, 0x52, 0x3e})
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}
	_, ok := msg.(*wsjtx.LoggedADIF)
	if !ok {
		t.Errorf("expected an ADIF message")
	}
}
func TestDecodeQSO(t *testing.T) {
	msg, err := wsjtx.Decode([]byte{0xad, 0xbc, 0xcb, 0xda, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x06, 0x57, 0x53, 0x4a, 0x54, 0x2d, 0x58, 0x00, 0x00, 0x00, 0x00, 0x00, 0x25, 0x85, 0x53, 0x00, 0x7d, 0x33, 0x38, 0x01, 0x00, 0x00, 0x00, 0x06, 0x57, 0x42, 0x36, 0x46, 0x57, 0x53, 0x00, 0x00, 0x00, 0x04, 0x44, 0x4d, 0x31, 0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd6, 0xcc, 0x09, 0x00, 0x00, 0x00, 0x03, 0x46, 0x54, 0x38, 0x00, 0x00, 0x00, 0x03, 0x2b, 0x30, 0x37, 0x00, 0x00, 0x00, 0x03, 0x2d, 0x30, 0x32, 0x00, 0x00, 0x00, 0x03, 0x32, 0x35, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x25, 0x85, 0x53, 0x00, 0x7c, 0x4a, 0x5a, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x57, 0x34, 0x54, 0x4e, 0x4c, 0x00, 0x00, 0x00, 0x06, 0x45, 0x4d, 0x36, 0x34, 0x4f, 0x52, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}
	qlog := msg.(*wsjtx.QSOLogged)
	exp := "20200423"
	got := adif.UTCDate(qlog.QSOOff)
	if got != exp {
		t.Errorf("expected date = %s, got %s", exp, got)
	}
	exp = "0216"
	got = adif.UTCTime(qlog.QSOOff)
	if got != exp {
		t.Errorf("expected time = %s, got %s", exp, got)
	}
}
