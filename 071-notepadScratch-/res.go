package main

var resources []byte = []byte {
	0x53,0x41,0x72,0x00,0x16,0x00,0x00,0x00,0x6e,0x00,0xff,0xff,0x01,0x00,0x0c,0x00,0x6f,0x00,0xff,0xff,0x02,0x00,0xff,0xff,0x74,0x00,0xff,0xff,0x03,0x00,0xff,0xff,0x65,0x00,0xff,0xff,0x04,0x00,0xff,0xff,0x70,
	0x00,0xff,0xff,0x05,0x00,0xff,0xff,0x61,0x00,0xff,0xff,0x06,0x00,0xff,0xff,0x64,0x00,0xff,0xff,0x07,0x00,0xff,0xff,0x2e,0x00,0xff,0xff,0x08,0x00,0xff,0xff,0x68,0x00,0xff,0xff,0x09,0x00,0xff,0xff,0x74,0x00,
	0xff,0xff,0x0a,0x00,0xff,0xff,0x6d,0x00,0xff,0xff,0x0b,0x00,0xff,0xff,0x00,0x00,0xff,0xff,0x01,0x00,0xff,0xff,0x73,0x00,0xff,0xff,0x0d,0x00,0xff,0xff,0x74,0x00,0xff,0xff,0x0e,0x00,0xff,0xff,0x79,0x00,0xff,
	0xff,0x0f,0x00,0xff,0xff,0x6c,0x00,0xff,0xff,0x10,0x00,0xff,0xff,0x65,0x00,0xff,0xff,0x11,0x00,0xff,0xff,0x2e,0x00,0xff,0xff,0x12,0x00,0xff,0xff,0x63,0x00,0xff,0xff,0x13,0x00,0xff,0xff,0x73,0x00,0xff,0xff,
	0x14,0x00,0xff,0xff,0x73,0x00,0xff,0xff,0x15,0x00,0xff,0xff,0x00,0x00,0xff,0xff,0x02,0x00,0xff,0xff,0x02,0x00,0x00,0x00,0xd4,0x00,0x00,0x00,0x1a,0x03,0x00,0x00,0xe8,0x06,0x00,0x00,0xee,0x03,0x00,0x00,0x81,
	0x02,0x00,0x00,0x6d,0x06,0x00,0x00,0x08,0x3c,0x68,0x74,0x6d,0x6c,0x3e,0x0d,0x0a,0x20,0x20,0x00,0x04,0x3c,0x68,0x65,0x61,0x64,0xa0,0x0b,0x40,0x00,0x0c,0x3c,0x73,0x74,0x79,0x6c,0x65,0x20,0x20,0x73,0x72,0x63,
	0x3d,0x22,0x60,0x0b,0x06,0x2e,0x63,0x73,0x73,0x22,0x20,0x2f,0xa0,0x23,0x01,0x3c,0x2f,0xe0,0x02,0x30,0x05,0x3c,0x62,0x6f,0x64,0x79,0x3e,0x20,0x08,0xe0,0x03,0x00,0x80,0x1a,0x40,0x00,0x1f,0x3c,0x21,0x2d,0x2d,
	0x20,0x54,0x65,0x78,0x74,0x61,0x72,0x65,0x61,0x20,0x74,0x6f,0x20,0x67,0x65,0x74,0x20,0x69,0x6e,0x70,0x75,0x74,0x20,0x66,0x72,0x6f,0x6d,0x20,0x05,0x75,0x73,0x65,0x72,0x2d,0x2d,0xa0,0x4b,0x40,0x00,0x01,0x3c,
	0x74,0xc0,0x2c,0x0c,0x63,0x6f,0x6c,0x73,0x3d,0x22,0x34,0x38,0x22,0x20,0x72,0x6f,0x77,0x20,0x09,0x0b,0x32,0x34,0x22,0x20,0x23,0x64,0x61,0x74,0x61,0x3e,0x3c,0x2f,0xc0,0x24,0xe0,0x00,0x78,0x01,0x0d,0x0a,0xe0,
	0x02,0x42,0x40,0x73,0x1e,0x53,0x69,0x6d,0x75,0x6c,0x61,0x74,0x69,0x6e,0x67,0x20,0x4d,0x65,0x6e,0x75,0x62,0x61,0x72,0x20,0x66,0x6f,0x72,0x20,0x6e,0x6f,0x74,0x65,0x70,0x61,0x64,0x20,0xe0,0x05,0x73,0x07,0x64,
	0x69,0x76,0x20,0x63,0x6c,0x61,0x73,0x20,0x65,0x04,0x74,0x6f,0x70,0x2d,0x6d,0x20,0x30,0x00,0x22,0xe0,0x00,0x5a,0x40,0x00,0xe0,0x01,0x2b,0x40,0x00,0x02,0x3c,0x75,0x6c,0xc0,0x2e,0x40,0x2a,0x20,0x5b,0x00,0x22,
	0xe0,0x02,0x4d,0xc0,0x00,0x02,0x3c,0x6c,0x69,0xe0,0x0a,0x15,0x40,0x00,0x03,0x46,0x69,0x6c,0x65,0xe0,0x0d,0x19,0xe0,0x02,0x59,0x02,0x73,0x75,0x62,0xa0,0x88,0xe0,0x0d,0x2a,0x40,0x00,0x20,0x62,0x0a,0x20,0x23,
	0x6e,0x65,0x77,0x3e,0x4e,0x65,0x77,0x3c,0x2f,0xe0,0x10,0x6f,0x40,0x00,0x60,0x2a,0x05,0x6f,0x70,0x65,0x6e,0x3e,0x4f,0x20,0x04,0xe0,0x1b,0x2c,0x05,0x73,0x61,0x76,0x65,0x3e,0x53,0x20,0x04,0xe0,0x1b,0x2c,0x05,
	0x65,0x78,0x69,0x74,0x3e,0x45,0x20,0x04,0xe0,0x12,0x2c,0x02,0x3c,0x2f,0x75,0xc2,0x72,0xe0,0x03,0x00,0x60,0x31,0x20,0x07,0xe0,0x04,0x00,0xe0,0x05,0x26,0xe0,0x06,0x39,0x01,0x3c,0x2f,0x21,0xb0,0x40,0x32,0xe1,
	0x04,0xf5,0x0a,0x73,0x63,0x72,0x69,0x70,0x74,0x20,0x74,0x79,0x70,0x65,0x21,0xc6,0x22,0x20,0x02,0x2f,0x74,0x69,0x80,0x13,0xe1,0x03,0x42,0xe0,0x01,0x09,0x40,0x00,0x04,0x65,0x76,0x65,0x6e,0x74,0x21,0x71,0x05,
	0x69,0x63,0x6b,0x20,0x24,0x28,0x60,0xc5,0x01,0x29,0x7b,0xe0,0x05,0x22,0x40,0x00,0x11,0x76,0x69,0x65,0x77,0x2e,0x63,0x6c,0x6f,0x73,0x65,0x57,0x69,0x6e,0x64,0x6f,0x77,0x28,0x29,0xe0,0x05,0x23,0x00,0x7d,0xe0,
	0x01,0x0e,0xe0,0x01,0x09,0x40,0x00,0xe0,0x06,0x5f,0x41,0x7f,0xe0,0x0b,0x5f,0x14,0x63,0x6f,0x6e,0x73,0x74,0x20,0x48,0x54,0x4d,0x4c,0x5f,0x46,0x49,0x4c,0x45,0x53,0x20,0x3d,0x20,0x22,0x4e,0xa2,0xa9,0x07,0x50,
	0x6c,0x75,0x73,0x20,0x4d,0x69,0x6e,0x20,0x05,0x08,0x28,0x2a,0x2e,0x6e,0x70,0x70,0x6d,0x29,0x7c,0x80,0x07,0x01,0x22,0x3b,0x20,0x3a,0xe0,0x06,0x70,0x40,0x00,0x00,0x76,0x42,0xeb,0x22,0x62,0x03,0x70,0x61,0x74,
	0x68,0x20,0x49,0x60,0xbc,0x05,0x73,0x65,0x6c,0x65,0x63,0x74,0x42,0x78,0x80,0x85,0x00,0x2c,0xe0,0x03,0x6d,0x0a,0x2c,0x20,0x22,0x55,0x6e,0x6e,0x61,0x6d,0x65,0x64,0x22,0x20,0x0a,0x40,0x3a,0x40,0x07,0x40,0x2c,
	0x00,0x20,0x40,0x15,0x43,0xb6,0x42,0x0c,0x03,0x3a,0x22,0x29,0x3b,0xe0,0x09,0x6d,0x01,0x2f,0x2f,0x23,0xcf,0x07,0x63,0x61,0x73,0x65,0x20,0x6f,0x66,0x20,0x60,0x77,0x03,0x20,0x77,0x61,0x73,0x43,0x68,0x00,0x20,
	0x80,0x74,0x06,0x65,0x64,0x20,0x74,0x68,0x65,0x6e,0xe0,0x0c,0x39,0x0e,0x66,0x75,0x6e,0x63,0x74,0x69,0x6f,0x6e,0x20,0x72,0x65,0x74,0x75,0x72,0x6e,0x20,0x36,0x02,0x75,0x6c,0x6c,0xe0,0x09,0x29,0x03,0x69,0x66,
	0x20,0x28,0x40,0x59,0x40,0xd1,0xe1,0x0b,0x3f,0x40,0x00,0x20,0x6c,0x00,0x66,0x64,0x1e,0x00,0x2e,0x41,0xfb,0xc0,0xf7,0x40,0xeb,0xe0,0x01,0x3a,0xe1,0x0b,0x2d,0xe1,0x02,0xbb,0x40,0x00,0x20,0x0e,0xe0,0x05,0x10,
	0xe1,0x06,0xc2,0x41,0x08,0xe0,0x0b,0x82,0xe1,0x33,0xc2,0xe0,0x01,0x00,0xe0,0x09,0x57,0xe1,0x17,0xcc,0x40,0x8f,0xe1,0x41,0xcc,0x80,0x5e,0x20,0x1e,0xe1,0x00,0x33,0x00,0x2c,0x20,0x6c,0xe1,0x02,0x57,0x21,0x42,
	0xe0,0x00,0x00,0xe0,0x05,0x3e,0x21,0x38,0x20,0x05,0x60,0x00,0xe0,0x01,0x09,0x40,0x00,0xe1,0x06,0x40,0x24,0xae,0xe1,0x0b,0x3f,0xe0,0x05,0x66,0x20,0xe9,0x01,0x22,0x22,0xe0,0x05,0x24,0xe0,0x0c,0x5f,0x01,0x3c,
	0x2f,0x83,0xe1,0x23,0xe0,0x80,0x14,0x01,0x3c,0x2f,0x66,0x8f,0x01,0x0d,0x0a,0x26,0xa4,0x03,0x74,0x6d,0x6c,0x3e,0x02,0x0d,0x0a,0x20,0xa0,0x00,0x0a,0x2e,0x74,0x6f,0x70,0x2d,0x6d,0x65,0x6e,0x75,0x7b,0x0d,0xe0,
	0x00,0x13,0x40,0x00,0x12,0x70,0x6f,0x73,0x69,0x74,0x69,0x6f,0x6e,0x3a,0x20,0x61,0x62,0x73,0x6f,0x6c,0x75,0x74,0x65,0x3b,0xe0,0x05,0x20,0x0c,0x77,0x69,0x64,0x74,0x68,0x3a,0x20,0x33,0x30,0x30,0x64,0x69,0x70,
	0xe0,0x06,0x1b,0x20,0x53,0x02,0x3a,0x20,0x30,0xe0,0x06,0x14,0x03,0x6c,0x65,0x66,0x74,0xe0,0x09,0x15,0x06,0x70,0x61,0x64,0x64,0x69,0x6e,0x67,0xe0,0x09,0x18,0x04,0x6d,0x61,0x72,0x67,0x69,0x20,0x7e,0xe0,0x03,
	0x17,0x00,0x7d,0xe0,0x01,0x0a,0x02,0x75,0x6c,0x2e,0x40,0xb5,0x02,0x62,0x61,0x72,0xe0,0x06,0xb8,0xa0,0x97,0xe0,0x0d,0x98,0x14,0x62,0x61,0x63,0x6b,0x67,0x72,0x6f,0x75,0x6e,0x64,0x2d,0x63,0x6f,0x6c,0x6f,0x72,
	0x20,0x3a,0x20,0x23,0x63,0x60,0x00,0x00,0x3b,0x20,0x1d,0xe0,0x00,0x00,0xe0,0x05,0x34,0xe0,0x01,0xa2,0x04,0x20,0x20,0x34,0x70,0x78,0x40,0x29,0xe0,0x05,0x20,0xe0,0x01,0xaa,0x20,0x0c,0x60,0x00,0xe0,0x01,0x1f,
	0xe0,0x0c,0xb2,0x05,0x20,0x3e,0x20,0x6c,0x69,0x20,0xe0,0x06,0xb8,0x14,0x64,0x69,0x73,0x70,0x6c,0x61,0x79,0x3a,0x20,0x69,0x6e,0x6c,0x69,0x6e,0x65,0x2d,0x62,0x6c,0x6f,0x63,0x6b,0xe0,0x06,0xbf,0x60,0xb4,0x20,
	0xb3,0x00,0x36,0x60,0x00,0xe0,0x06,0x1c,0xe0,0x02,0xa7,0x00,0x35,0xe0,0x14,0xa6,0xe0,0x0d,0x1e,0xe1,0x01,0xf1,0x06,0x72,0x65,0x6c,0x61,0x74,0x69,0x76,0xe1,0x03,0xf1,0xe0,0x11,0xc6,0x05,0x3a,0x68,0x6f,0x76,
	0x65,0x72,0xe0,0x07,0xcc,0xc0,0xa8,0x21,0x59,0x20,0x00,0xe0,0x02,0x48,0x40,0x00,0xe1,0x01,0x85,0x20,0x21,0x20,0xc7,0x20,0x00,0x40,0x8e,0xe0,0x0d,0x00,0xe0,0x01,0x3a,0x00,0x7d,0x20,0x03,0x20,0x0e,0xe0,0x01,
	0x10,0x20,0x89,0x02,0x73,0x75,0x62,0x62,0xbd,0xe0,0x07,0x7f,0xe0,0x01,0xcc,0xe2,0x0e,0xbe,0xe1,0x00,0x6d,0x02,0x6e,0x6f,0x6e,0xe0,0x07,0x1b,0x62,0xbe,0x01,0x31,0x36,0xe2,0x09,0x3e,0xe2,0x0d,0xc2,0xa2,0x71,
	0x03,0x61,0x75,0x74,0x6f,0xe0,0x06,0x19,0x04,0x68,0x65,0x69,0x67,0x68,0x20,0x31,0xe0,0x0a,0x1a,0x09,0x6c,0x69,0x73,0x74,0x2d,0x73,0x74,0x79,0x6c,0x65,0xe0,0x0c,0x82,0xe1,0x00,0xab,0xe0,0x06,0x17,0xe1,0x01,
	0xe3,0xe0,0x06,0x18,0x09,0x62,0x6f,0x78,0x2d,0x73,0x68,0x61,0x64,0x6f,0x77,0x22,0xfb,0x00,0x32,0x20,0xba,0xe0,0x02,0x04,0x08,0x72,0x67,0x62,0x61,0x28,0x31,0x30,0x30,0x2c,0xc0,0x03,0x03,0x30,0x2e,0x36,0x29,
	0xe0,0x01,0x3e,0x61,0x57,0x21,0x58,0x20,0x04,0x60,0x00,0xe1,0x03,0x56,0x41,0xe1,0xe1,0x08,0x5b,0xc0,0x83,0x40,0x5c,0xe0,0x08,0x86,0x0a,0x72,0x64,0x65,0x72,0x2d,0x62,0x6f,0x74,0x74,0x6f,0x6d,0x40,0x89,0x03,
	0x70,0x78,0x20,0x64,0x20,0x0c,0x01,0x65,0x64,0xe2,0x0e,0x0b,0xc2,0x28,0x00,0x33,0x60,0x00,0xe0,0x02,0x1c,0xe2,0x18,0x71,0x00,0x3e,0x40,0x17,0xe0,0x00,0xa8,0x00,0x7b,0x20,0x27,0xe0,0x0e,0x00,0xe0,0x01,0x48,
	0x40,0x00,0xc1,0xf8,0xe3,0x0b,0x5e,0xe2,0x04,0x98,0x02,0x66,0x66,0x66,0xe2,0x03,0x95,0xe0,0x01,0x29,0xe1,0x07,0x2f,0x07,0x74,0x65,0x78,0x74,0x61,0x72,0x65,0x61,0x20,0x82,0xe0,0x01,0x14,0x40,0x00,0xe2,0x18,
	0x84,0x42,0x68,0xe2,0x10,0x67,0x64,0xbe,0x20,0x0c,0x20,0x00,0xe0,0x05,0x1d,0xa2,0x6f,0x02,0x32,0x39,0x35,0xe0,0x0f,0x21,0x02,0x66,0x6f,0x6e,0x22,0x5c,0x01,0x69,0x7a,0x22,0x5b,0x01,0x31,0x34,0xe0,0x00,0x24,
	0xe0,0x05,0x23,0xe0,0x02,0xeb,0x09,0x61,0x6e,0x74,0x69,0x71,0x75,0x65,0x77,0x68,0x69,0xe0,0x08,0xa1,0x81,0xcf,0xa2,0xa2,0x20,0x0f,0xc0,0x00,0xe0,0x05,0x25,0x41,0x8f,0x04,0x66,0x6c,0x6f,0x77,0x2d,0x23,0x4f,
	0xe2,0x06,0xeb,0x02,0x7d,0x0d,0x0a,
}
