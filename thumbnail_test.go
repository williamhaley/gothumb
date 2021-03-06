package gothumb

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
	"github.com/koofr/goepeg"
)

var pngInput []byte = []byte{
  0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
  0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x10,
  0x08, 0x06, 0x00, 0x00, 0x00, 0x1f, 0xf3, 0xff, 0x61, 0x00, 0x00, 0x00,
  0x06, 0x62, 0x4b, 0x47, 0x44, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0xa0,
  0xbd, 0xa7, 0x93, 0x00, 0x00, 0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00,
  0x00, 0x0b, 0x13, 0x00, 0x00, 0x0b, 0x13, 0x01, 0x00, 0x9a, 0x9c, 0x18,
  0x00, 0x00, 0x00, 0x07, 0x74, 0x49, 0x4d, 0x45, 0x07, 0xe0, 0x08, 0x03,
  0x0d, 0x1f, 0x08, 0x33, 0x0c, 0x40, 0xf2, 0x00, 0x00, 0x00, 0x19, 0x74,
  0x45, 0x58, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x00, 0x43,
  0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20,
  0x47, 0x49, 0x4d, 0x50, 0x57, 0x81, 0x0e, 0x17, 0x00, 0x00, 0x00, 0x4b,
  0x49, 0x44, 0x41, 0x54, 0x38, 0xcb, 0xd5, 0xd1, 0xc1, 0x0a, 0x00, 0x20,
  0x08, 0x03, 0xd0, 0x2d, 0xfc, 0xff, 0x5f, 0x5e, 0xd7, 0x20, 0x43, 0xc5,
  0x3a, 0xb4, 0xab, 0x32, 0x7c, 0x48, 0x49, 0x42, 0x23, 0x03, 0xcd, 0x18,
  0x00, 0x90, 0xdc, 0x06, 0xd9, 0xc3, 0x6c, 0x5d, 0x26, 0x89, 0xaa, 0xe8,
  0x0e, 0xc1, 0x4b, 0x96, 0x75, 0x2c, 0xc8, 0xb2, 0xde, 0x11, 0xb2, 0x2c,
  0xab, 0xbc, 0xce, 0x63, 0xb5, 0x09, 0xef, 0x0b, 0x22, 0xd6, 0x07, 0x84,
  0x88, 0x35, 0x01, 0x11, 0x2a, 0x25, 0x19, 0x49, 0xc2, 0x6f, 0xa3, 0x00,
  0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,
}
var pngOutput []byte = []byte{
  0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
  0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x08,
  0x10, 0x02, 0x00, 0x00, 0x00, 0x1b, 0xfd, 0xf5, 0x9f, 0x00, 0x00, 0x00,
  0x4b, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0x62, 0xf9, 0x0f, 0x06, 0x0c,
  0x44, 0x03, 0xa6, 0x03, 0x0c, 0x20, 0xe8, 0x00, 0x86, 0x0d, 0x60, 0x48,
  0x00, 0x40, 0x6c, 0xa8, 0x07, 0x43, 0x64, 0x36, 0x04, 0xee, 0x07, 0x43,
  0x64, 0xc0, 0x84, 0x69, 0x04, 0xb2, 0x3d, 0x10, 0x9b, 0x51, 0x9c, 0x84,
  0x4b, 0x02, 0x02, 0xb0, 0x38, 0xf8, 0x3f, 0x0e, 0x80, 0xe9, 0x48, 0x9c,
  0x4e, 0xc2, 0x74, 0x18, 0x16, 0x27, 0xe1, 0x07, 0xc8, 0x0e, 0x06, 0x04,
  0x00, 0x00, 0xff, 0xff, 0x04, 0xec, 0x6b, 0xa9, 0x65, 0xee, 0x5f, 0xe7,
  0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42, 0x60, 0x82,	
}
var gifInput []byte = []byte{
  0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x10, 0x00, 0x10, 0x00, 0xa1, 0x01,
  0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
  0xff, 0x21, 0xf9, 0x04, 0x01, 0x0a, 0x00, 0x02, 0x00, 0x2c, 0x00, 0x00,
  0x00, 0x00, 0x10, 0x00, 0x10, 0x00, 0x00, 0x02, 0x20, 0x8c, 0x8f, 0xa9,
  0x9b, 0xe0, 0xce, 0x5a, 0x54, 0x60, 0xca, 0xf0, 0xe0, 0xaa, 0x87, 0x33,
  0x8f, 0x4d, 0x20, 0x97, 0x81, 0xc8, 0x28, 0x9e, 0x69, 0x67, 0xa9, 0xad,
  0x61, 0xae, 0xef, 0x51, 0x00, 0x00, 0x3b,	
}
var gifOutput []byte = []byte{
  0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x08, 0x00, 0x08, 0x00, 0x87, 0x00,
  0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x00, 0x00, 0x88, 0x00, 0x00,
  0xcc, 0x00, 0x44, 0x00, 0x00, 0x44, 0x44, 0x00, 0x44, 0x88, 0x00, 0x44,
  0xcc, 0x00, 0x88, 0x00, 0x00, 0x88, 0x44, 0x00, 0x88, 0x88, 0x00, 0x88,
  0xcc, 0x00, 0xcc, 0x00, 0x00, 0xcc, 0x44, 0x00, 0xcc, 0x88, 0x00, 0xcc,
  0xcc, 0x00, 0xdd, 0xdd, 0x11, 0x11, 0x11, 0x00, 0x00, 0x55, 0x00, 0x00,
  0x99, 0x00, 0x00, 0xdd, 0x00, 0x55, 0x00, 0x00, 0x55, 0x55, 0x00, 0x4c,
  0x99, 0x00, 0x49, 0xdd, 0x00, 0x99, 0x00, 0x00, 0x99, 0x4c, 0x00, 0x99,
  0x99, 0x00, 0x93, 0xdd, 0x00, 0xdd, 0x00, 0x00, 0xdd, 0x49, 0x00, 0xdd,
  0x93, 0x00, 0xee, 0x9e, 0x00, 0xee, 0xee, 0x22, 0x22, 0x22, 0x00, 0x00,
  0x66, 0x00, 0x00, 0xaa, 0x00, 0x00, 0xee, 0x00, 0x66, 0x00, 0x00, 0x66,
  0x66, 0x00, 0x55, 0xaa, 0x00, 0x4f, 0xee, 0x00, 0xaa, 0x00, 0x00, 0xaa,
  0x55, 0x00, 0xaa, 0xaa, 0x00, 0x9e, 0xee, 0x00, 0xee, 0x00, 0x00, 0xee,
  0x4f, 0x00, 0xff, 0x55, 0x00, 0xff, 0xaa, 0x00, 0xff, 0xff, 0x33, 0x33,
  0x33, 0x00, 0x00, 0x77, 0x00, 0x00, 0xbb, 0x00, 0x00, 0xff, 0x00, 0x77,
  0x00, 0x00, 0x77, 0x77, 0x00, 0x5d, 0xbb, 0x00, 0x55, 0xff, 0x00, 0xbb,
  0x00, 0x00, 0xbb, 0x5d, 0x00, 0xbb, 0xbb, 0x00, 0xaa, 0xff, 0x00, 0xff,
  0x00, 0x44, 0x00, 0x44, 0x44, 0x00, 0x88, 0x44, 0x00, 0xcc, 0x44, 0x44,
  0x00, 0x44, 0x44, 0x44, 0x44, 0x44, 0x88, 0x44, 0x44, 0xcc, 0x44, 0x88,
  0x00, 0x44, 0x88, 0x44, 0x44, 0x88, 0x88, 0x44, 0x88, 0xcc, 0x44, 0xcc,
  0x00, 0x44, 0xcc, 0x44, 0x44, 0xcc, 0x88, 0x44, 0xcc, 0xcc, 0x44, 0x00,
  0x00, 0x55, 0x00, 0x00, 0x55, 0x00, 0x55, 0x4c, 0x00, 0x99, 0x49, 0x00,
  0xdd, 0x55, 0x55, 0x00, 0x55, 0x55, 0x55, 0x4c, 0x4c, 0x99, 0x49, 0x49,
  0xdd, 0x4c, 0x99, 0x00, 0x4c, 0x99, 0x4c, 0x4c, 0x99, 0x99, 0x49, 0x93,
  0xdd, 0x49, 0xdd, 0x00, 0x49, 0xdd, 0x49, 0x49, 0xdd, 0x93, 0x49, 0xdd,
  0xdd, 0x4f, 0xee, 0xee, 0x66, 0x00, 0x00, 0x66, 0x00, 0x66, 0x55, 0x00,
  0xaa, 0x4f, 0x00, 0xee, 0x66, 0x66, 0x00, 0x66, 0x66, 0x66, 0x55, 0x55,
  0xaa, 0x4f, 0x4f, 0xee, 0x55, 0xaa, 0x00, 0x55, 0xaa, 0x55, 0x55, 0xaa,
  0xaa, 0x4f, 0x9e, 0xee, 0x4f, 0xee, 0x00, 0x4f, 0xee, 0x4f, 0x4f, 0xee,
  0x9e, 0x55, 0xff, 0xaa, 0x55, 0xff, 0xff, 0x77, 0x00, 0x00, 0x77, 0x00,
  0x77, 0x5d, 0x00, 0xbb, 0x55, 0x00, 0xff, 0x77, 0x77, 0x00, 0x77, 0x77,
  0x77, 0x5d, 0x5d, 0xbb, 0x55, 0x55, 0xff, 0x5d, 0xbb, 0x00, 0x5d, 0xbb,
  0x5d, 0x5d, 0xbb, 0xbb, 0x55, 0xaa, 0xff, 0x55, 0xff, 0x00, 0x55, 0xff,
  0x55, 0x88, 0x00, 0x88, 0x88, 0x00, 0xcc, 0x88, 0x44, 0x00, 0x88, 0x44,
  0x44, 0x88, 0x44, 0x88, 0x88, 0x44, 0xcc, 0x88, 0x88, 0x00, 0x88, 0x88,
  0x44, 0x88, 0x88, 0x88, 0x88, 0x88, 0xcc, 0x88, 0xcc, 0x00, 0x88, 0xcc,
  0x44, 0x88, 0xcc, 0x88, 0x88, 0xcc, 0xcc, 0x88, 0x00, 0x00, 0x88, 0x00,
  0x44, 0x99, 0x00, 0x4c, 0x99, 0x00, 0x99, 0x93, 0x00, 0xdd, 0x99, 0x4c,
  0x00, 0x99, 0x4c, 0x4c, 0x99, 0x4c, 0x99, 0x93, 0x49, 0xdd, 0x99, 0x99,
  0x00, 0x99, 0x99, 0x4c, 0x99, 0x99, 0x99, 0x93, 0x93, 0xdd, 0x93, 0xdd,
  0x00, 0x93, 0xdd, 0x49, 0x93, 0xdd, 0x93, 0x93, 0xdd, 0xdd, 0x99, 0x00,
  0x00, 0xaa, 0x00, 0x00, 0xaa, 0x00, 0x55, 0xaa, 0x00, 0xaa, 0x9e, 0x00,
  0xee, 0xaa, 0x55, 0x00, 0xaa, 0x55, 0x55, 0xaa, 0x55, 0xaa, 0x9e, 0x4f,
  0xee, 0xaa, 0xaa, 0x00, 0xaa, 0xaa, 0x55, 0xaa, 0xaa, 0xaa, 0x9e, 0x9e,
  0xee, 0x9e, 0xee, 0x00, 0x9e, 0xee, 0x4f, 0x9e, 0xee, 0x9e, 0x9e, 0xee,
  0xee, 0xaa, 0xff, 0xff, 0xbb, 0x00, 0x00, 0xbb, 0x00, 0x5d, 0xbb, 0x00,
  0xbb, 0xaa, 0x00, 0xff, 0xbb, 0x5d, 0x00, 0xbb, 0x5d, 0x5d, 0xbb, 0x5d,
  0xbb, 0xaa, 0x55, 0xff, 0xbb, 0xbb, 0x00, 0xbb, 0xbb, 0x5d, 0xbb, 0xbb,
  0xbb, 0xaa, 0xaa, 0xff, 0xaa, 0xff, 0x00, 0xaa, 0xff, 0x55, 0xaa, 0xff,
  0xaa, 0xcc, 0x00, 0xcc, 0xcc, 0x44, 0x00, 0xcc, 0x44, 0x44, 0xcc, 0x44,
  0x88, 0xcc, 0x44, 0xcc, 0xcc, 0x88, 0x00, 0xcc, 0x88, 0x44, 0xcc, 0x88,
  0x88, 0xcc, 0x88, 0xcc, 0xcc, 0xcc, 0x00, 0xcc, 0xcc, 0x44, 0xcc, 0xcc,
  0x88, 0xcc, 0xcc, 0xcc, 0xcc, 0x00, 0x00, 0xcc, 0x00, 0x44, 0xcc, 0x00,
  0x88, 0xdd, 0x00, 0x93, 0xdd, 0x00, 0xdd, 0xdd, 0x49, 0x00, 0xdd, 0x49,
  0x49, 0xdd, 0x49, 0x93, 0xdd, 0x49, 0xdd, 0xdd, 0x93, 0x00, 0xdd, 0x93,
  0x49, 0xdd, 0x93, 0x93, 0xdd, 0x93, 0xdd, 0xdd, 0xdd, 0x00, 0xdd, 0xdd,
  0x49, 0xdd, 0xdd, 0x93, 0xdd, 0xdd, 0xdd, 0xdd, 0x00, 0x00, 0xdd, 0x00,
  0x49, 0xee, 0x00, 0x4f, 0xee, 0x00, 0x9e, 0xee, 0x00, 0xee, 0xee, 0x4f,
  0x00, 0xee, 0x4f, 0x4f, 0xee, 0x4f, 0x9e, 0xee, 0x4f, 0xee, 0xee, 0x9e,
  0x00, 0xee, 0x9e, 0x4f, 0xee, 0x9e, 0x9e, 0xee, 0x9e, 0xee, 0xee, 0xee,
  0x00, 0xee, 0xee, 0x4f, 0xee, 0xee, 0x9e, 0xee, 0xee, 0xee, 0xee, 0x00,
  0x00, 0xff, 0x00, 0x00, 0xff, 0x00, 0x55, 0xff, 0x00, 0xaa, 0xff, 0x00,
  0xff, 0xff, 0x55, 0x00, 0xff, 0x55, 0x55, 0xff, 0x55, 0xaa, 0xff, 0x55,
  0xff, 0xff, 0xaa, 0x00, 0xff, 0xaa, 0x55, 0xff, 0xaa, 0xaa, 0xff, 0xaa,
  0xff, 0xff, 0xff, 0x00, 0xff, 0xff, 0x55, 0xff, 0xff, 0xaa, 0xff, 0xff,
  0xff, 0x2c, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x08, 0x00, 0x00, 0x08,
  0x22, 0xff, 0x04, 0x0e, 0x14, 0xb8, 0x8b, 0xc8, 0x1d, 0x82, 0x02, 0x11,
  0xfd, 0xbb, 0x83, 0x68, 0xd7, 0x40, 0x84, 0x0c, 0x09, 0x3e, 0xfc, 0xe7,
  0xf0, 0x20, 0x41, 0x84, 0x0b, 0x13, 0x4a, 0xdc, 0xf8, 0x30, 0x20, 0x00,
  0x3b,
}
var sourceInput []byte = []byte{
	255, 216, 255, 219, 0, 67, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 255, 219, 0, 67, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 255, 194, 0, 17, 8, 0, 16, 0, 16, 3, 1, 17, 0,
	2, 17, 1, 3, 17, 1, 255, 196, 0, 22, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 9, 7, 8, 255, 196, 0, 23, 1, 0, 3, 1, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 9, 10, 7, 255, 218, 0, 12, 3, 1, 0,
	2, 16, 3, 16, 0, 0, 1, 197, 226, 155, 252, 74, 5, 184, 87, 6, 237,
	49, 172, 66, 208, 30, 169, 255, 196, 0, 20, 16, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0, 8, 1, 1, 0, 1, 5, 2, 31,
	255, 196, 0, 20, 17, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	32, 255, 218, 0, 8, 1, 3, 1, 1, 63, 1, 31, 255, 196, 0, 20, 17, 1,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0, 8, 1,
	2, 1, 1, 63, 1, 31, 255, 196, 0, 20, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0, 8, 1, 1, 0, 6, 63, 2, 31, 255,
	196, 0, 20, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32,
	255, 218, 0, 8, 1, 1, 0, 1, 63, 33, 31, 255, 218, 0, 12, 3, 1, 0, 2,
	0, 3, 0, 0, 0, 16, 0, 15, 255, 196, 0, 20, 17, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0, 8, 1, 3, 1, 1, 63, 16,
	31, 255, 196, 0, 20, 17, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 32, 255, 218, 0, 8, 1, 2, 1, 1, 63, 16, 31, 255, 196, 0, 20,
	16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 255, 218, 0,
	8, 1, 1, 0, 1, 63, 16, 31, 255, 217,
}
var thumbnailOutput []byte = []byte{
	255, 216, 255, 219, 0, 132, 0, 8, 6, 6, 7, 6, 5, 8, 7, 7, 7, 9, 9, 8,
	10, 12, 20, 13, 12, 11, 11, 12, 25, 18, 19, 15, 20, 29, 26, 31, 30,
	29, 26, 28, 28, 32, 36, 46, 39, 32, 34, 44, 35, 28, 28, 40, 55, 41,
	44, 48, 49, 52, 52, 52, 31, 39, 57, 61, 56, 50, 60, 46, 51, 52, 50, 1,
	9, 9, 9, 12, 11, 12, 24, 13, 13, 24, 50, 33, 28, 33, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50,
	50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 50, 255, 192, 0, 17, 8, 0,
	4, 0, 4, 3, 1, 34, 0, 2, 17, 1, 3, 17, 1, 255, 196, 1, 162, 0, 0, 1,
	5, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 16, 0, 2, 1, 3, 3, 2, 4, 3, 5, 5, 4, 4, 0, 0, 1, 125, 1, 2,
	3, 0, 4, 17, 5, 18, 33, 49, 65, 6, 19, 81, 97, 7, 34, 113, 20, 50,
	129, 145, 161, 8, 35, 66, 177, 193, 21, 82, 209, 240, 36, 51, 98, 114,
	130, 9, 10, 22, 23, 24, 25, 26, 37, 38, 39, 40, 41, 42, 52, 53, 54,
	55, 56, 57, 58, 67, 68, 69, 70, 71, 72, 73, 74, 83, 84, 85, 86, 87,
	88, 89, 90, 99, 100, 101, 102, 103, 104, 105, 106, 115, 116, 117, 118,
	119, 120, 121, 122, 131, 132, 133, 134, 135, 136, 137, 138, 146, 147,
	148, 149, 150, 151, 152, 153, 154, 162, 163, 164, 165, 166, 167, 168,
	169, 170, 178, 179, 180, 181, 182, 183, 184, 185, 186, 194, 195, 196,
	197, 198, 199, 200, 201, 202, 210, 211, 212, 213, 214, 215, 216, 217,
	218, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 241, 242, 243,
	244, 245, 246, 247, 248, 249, 250, 1, 0, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 17, 0, 2, 1, 2,
	4, 4, 3, 4, 7, 5, 4, 4, 0, 1, 2, 119, 0, 1, 2, 3, 17, 4, 5, 33, 49, 6,
	18, 65, 81, 7, 97, 113, 19, 34, 50, 129, 8, 20, 66, 145, 161, 177,
	193, 9, 35, 51, 82, 240, 21, 98, 114, 209, 10, 22, 36, 52, 225, 37,
	241, 23, 24, 25, 26, 38, 39, 40, 41, 42, 53, 54, 55, 56, 57, 58, 67,
	68, 69, 70, 71, 72, 73, 74, 83, 84, 85, 86, 87, 88, 89, 90, 99, 100,
	101, 102, 103, 104, 105, 106, 115, 116, 117, 118, 119, 120, 121, 122,
	130, 131, 132, 133, 134, 135, 136, 137, 138, 146, 147, 148, 149, 150,
	151, 152, 153, 154, 162, 163, 164, 165, 166, 167, 168, 169, 170, 178,
	179, 180, 181, 182, 183, 184, 185, 186, 194, 195, 196, 197, 198, 199,
	200, 201, 202, 210, 211, 212, 213, 214, 215, 216, 217, 218, 226, 227,
	228, 229, 230, 231, 232, 233, 234, 242, 243, 244, 245, 246, 247, 248,
	249, 250, 255, 218, 0, 12, 3, 1, 0, 2, 17, 3, 17, 0, 63, 0, 242, 83,
	227, 95, 16, 216, 159, 179, 91, 106, 27, 33, 79, 186, 190, 76, 103,
	25, 231, 169, 95, 83, 73, 255, 0, 11, 3, 197, 31, 244, 19, 255, 0,
	201, 120, 191, 248, 154, 193, 188, 255, 0, 143, 183, 252, 63, 149, 65,
	90, 102, 88, 12, 44, 113, 181, 163, 26, 81, 73, 74, 93, 23, 119, 228,
	114, 75, 7, 135, 172, 221, 74, 148, 227, 41, 75, 86, 218, 77, 182,
	247, 109, 245, 111, 171, 63, 255, 217,
}

func TestThumbnail(t *testing.T) {
	reader := bytes.NewReader(sourceInput)

	out, err := Thumbnail(reader, 4, 75, goepeg.ScaleTypeFitMax)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	content, err := ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(content, thumbnailOutput) {
		t.Errorf("Jpeg thumbnail failed")
	}

	reader = bytes.NewReader(pngInput)
	out, err = Thumbnail(reader, 8, 75, goepeg.ScaleTypeFitMax)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	content, err = ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}	
	if !reflect.DeepEqual(content, pngOutput) {
		t.Errorf("PNG thumbnail failed")
	}


	reader = bytes.NewReader(gifInput)
	out, err = Thumbnail(reader, 8, 75, goepeg.ScaleTypeFitMax)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	content, err = ioutil.ReadAll(out)
	if err != nil {
		t.Fatal(err)
	}		
	if !reflect.DeepEqual(content, gifOutput) {
		t.Errorf("GIF thumbnail failed")
	}
	
}
