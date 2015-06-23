// Copyright 2014 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.
package layers

import (
	"gopacket"
	"reflect"
	"testing"
)

// Test packet collected from live network. See the test below for contents
var SFlowTestPacket1 = []byte{
	0x84, 0x2b, 0x2b, 0x16, 0x8b, 0x62, 0xf0, 0x50, 0x56, 0x85, 0x3a, 0xfd, 0x08, 0x00, 0x45, 0x00,
	0x05, 0xbc, 0x9c, 0x04, 0x40, 0x00, 0xff, 0x11, 0xc7, 0x00, 0x0a, 0x01, 0xff, 0x0e, 0x0a, 0x01,
	0x00, 0x1b, 0xc7, 0x57, 0x18, 0xc7, 0x05, 0xa8, 0x22, 0x3b, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00,
	0x00, 0x01, 0x0a, 0x01, 0xf8, 0x16, 0x00, 0x00, 0x00, 0x11, 0x00, 0x00, 0x9d, 0xfb, 0x40, 0x49,
	0xc6, 0xcd, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x00, 0x26,
	0x27, 0xe8, 0x00, 0x00, 0x02, 0x13, 0x00, 0x00, 0x3e, 0x80, 0x50, 0xbd, 0xe5, 0x80, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xd2, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6d,
	0xf0, 0x40, 0x08, 0x00, 0x45, 0x00, 0x05, 0xc0, 0x6b, 0xaa, 0x40, 0x00, 0x40, 0x06, 0x8f, 0x41,
	0x0a, 0x01, 0x0e, 0x16, 0x36, 0xf0, 0xeb, 0x45, 0x76, 0xfd, 0x00, 0x50, 0xca, 0x77, 0xef, 0x96,
	0xfc, 0x28, 0x63, 0x40, 0x50, 0x10, 0x00, 0x3c, 0x64, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00,
	0xf4, 0x00, 0x00, 0x02, 0x77, 0x00, 0x00, 0x00, 0xfd, 0x3b, 0x8c, 0xe7, 0x04, 0x4a, 0x2d, 0xb2,
	0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1c, 0x00, 0x00, 0x01, 0x48, 0xcc, 0x11, 0x0d, 0xe3, 0x00,
	0x26, 0x85, 0x30, 0x00, 0x00, 0x07, 0x66, 0x00, 0x02, 0xd0, 0x8a, 0x00, 0x02, 0xce, 0xf0, 0x00,
	0x29, 0x7e, 0x80, 0x00, 0x02, 0xd0, 0x98, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x26, 0x85, 0x30, 0x00,
	0x00, 0x00, 0xf4, 0x00, 0x00, 0x02, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x01, 0x5e, 0x5c, 0x1e, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x07, 0xd0, 0xb1, 0x2f, 0xa2, 0x90, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x05, 0xee, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x3c, 0x8a,
	0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x08, 0x00, 0x45, 0x00, 0x05, 0xdc,
	0x9f, 0xfd, 0x40, 0x00, 0x40, 0x06, 0x6a, 0xfa, 0x0a, 0x01, 0x0e, 0x10, 0x0a, 0x01, 0x08, 0x13,
	0x23, 0x84, 0xb7, 0x22, 0x8a, 0xc9, 0x50, 0xb5, 0x4e, 0x10, 0x2a, 0x87, 0x80, 0x10, 0x06, 0x01,
	0x10, 0xa6, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xef, 0x1f, 0xf4, 0x07, 0x99, 0x3a, 0xd8, 0x5b,
	0x01, 0x46, 0x09, 0x00, 0x0c, 0x00, 0x0c, 0x3c, 0xac, 0x4a, 0x1b, 0x06, 0x04, 0x78, 0x78, 0x4e,
	0xc2, 0x05, 0x46, 0x43, 0x06, 0x04, 0x78, 0x78, 0xee, 0x9c, 0x00, 0x41, 0xef, 0x05, 0x81, 0x32,
	0x1b, 0x06, 0x04, 0x78, 0x78, 0x56, 0x72, 0x05, 0x4e, 0x92, 0x00, 0x96, 0x39, 0x00, 0xea, 0x3f,
	0x01, 0x15, 0xa3, 0x08, 0x04, 0x42, 0x6a, 0x82, 0x87, 0x08, 0x05, 0xcc, 0x00, 0x04, 0x00, 0x00,
	0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x01, 0x5a,
	0xcd, 0xd0, 0x00, 0x00, 0x02, 0x55, 0x00, 0x00, 0x07, 0xd0, 0x95, 0x67, 0xe1, 0x30, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x55, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x46, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x3c, 0x8a, 0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f,
	0x11, 0x28, 0x08, 0x00, 0x45, 0x00, 0x04, 0x34, 0xdb, 0x36, 0x40, 0x00, 0x40, 0x06, 0x38, 0xac,
	0x0a, 0x01, 0x0e, 0x11, 0x0a, 0x01, 0x00, 0xcf, 0x23, 0x84, 0xa0, 0x3f, 0x3c, 0xce, 0xd5, 0x4a,
	0x72, 0x0b, 0x5d, 0x1a, 0x80, 0x10, 0x06, 0x01, 0x8a, 0x50, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a,
	0xef, 0x1f, 0xa2, 0xba, 0xe6, 0xfa, 0xae, 0xb3, 0xfe, 0xcf, 0x00, 0x19, 0xcf, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x01, 0xb9, 0x79, 0xdd, 0x42, 0x00, 0x00, 0x02, 0x84, 0x9b, 0xa9, 0x02, 0xe2, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x06, 0x32, 0x39, 0x35, 0x34, 0x33, 0x36, 0x00, 0x00, 0x02, 0x70, 0xcd,
	0x16, 0x40, 0xa6, 0x98, 0x88, 0x24, 0x06, 0x50, 0xb0, 0xf4, 0xee, 0x03, 0xa6, 0xfa, 0x87, 0xaf,
	0xc1, 0x99, 0x52, 0x0d, 0x07, 0xa8, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x02, 0x00, 0x00, 0x00, 0xa8, 0x00, 0x00, 0x20, 0xf2, 0x00, 0x00, 0x02, 0x0a, 0x00, 0x00,
	0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x58, 0x00, 0x00, 0x02, 0x0a, 0x00, 0x00,
	0x00, 0x06, 0x00, 0x00, 0x00, 0x02, 0x54, 0x0b, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0x03, 0x00, 0x01, 0x29, 0x82, 0x6d, 0xb0, 0x6c, 0x0b, 0xcb, 0x0d, 0xdd, 0x96, 0x00, 0x06,
	0xa8, 0xc6, 0x00, 0x00, 0x00, 0x7b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x34, 0x02, 0x35, 0x58, 0x7c, 0x9e, 0x56, 0x64, 0x25, 0x71, 0x00, 0x70,
	0x5a, 0xc4, 0x00, 0x09, 0x08, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x01, 0x5e, 0x5c, 0x1f, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x07, 0xd0, 0xb1, 0x2f, 0xaa, 0x60, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x05, 0xee, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x3c, 0x8a,
	0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x08, 0x00, 0x45, 0x00, 0x05, 0xdc,
	0x0f, 0xba, 0x40, 0x00, 0x40, 0x06, 0xf4, 0x3f, 0x0a, 0x01, 0x0e, 0x10, 0x0a, 0x01, 0x0f, 0x11,
	0x23, 0x84, 0xcd, 0xc0, 0xf4, 0x0e, 0x90, 0x23, 0xd7, 0x32, 0x8b, 0x31, 0x80, 0x10, 0x00, 0x1d,
	0x6b, 0x12, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xef, 0x1f, 0xf4, 0x28, 0xef, 0x1f, 0xec, 0x76,
	0xaa, 0x25, 0x01, 0x04, 0xc0, 0xac, 0xfe, 0x25, 0x01, 0x8e, 0x25, 0x01, 0x16, 0xc7, 0x28, 0xfe,
	0x7e, 0x70, 0xfe, 0x7e, 0x70, 0x52, 0x7e, 0x70, 0x15, 0x9b, 0xfe, 0x35, 0x01, 0xfe, 0x35, 0x01,
	0x42, 0x35, 0x01, 0xfe, 0x95, 0x77, 0xfe, 0x95, 0x77, 0xfe, 0x95, 0x77, 0x52, 0x95, 0x77, 0x00,
	0xd2, 0xfe, 0x70, 0x02, 0x92, 0x70, 0x02, 0x16, 0x60, 0x22, 0x00, 0x7e, 0xb2, 0x15, 0x00, 0x00,
	0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xd0, 0x01, 0x5a,
	0xcd, 0xd1, 0x00, 0x00, 0x02, 0x55, 0x00, 0x00, 0x07, 0xd0, 0x95, 0x67, 0xe9, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x02, 0x55, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xee, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0xb8, 0xca, 0x3a, 0x6f,
	0x11, 0x28, 0x08, 0x00, 0x45, 0x00, 0x05, 0xdc, 0xfe, 0x05, 0x40, 0x00, 0x40, 0x06, 0x06, 0xf4,
	0x0a, 0x01, 0x0e, 0x11, 0x0a, 0x01, 0x0e, 0x10, 0x23, 0x84, 0xfa, 0x29, 0xae, 0xd4, 0x95, 0x03,
	0x99, 0xb8, 0x77, 0xd0, 0x80, 0x10, 0x00, 0x1d, 0x6f, 0x4f, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a,
	0xef, 0x1f, 0xa2, 0xcc, 0xef, 0x1f, 0xf4, 0x2c, 0xfe, 0xdb, 0x05, 0xa1, 0xdb, 0x04, 0x9e, 0xc0,
	0xfe, 0x30, 0x08, 0xb2, 0x30, 0x08, 0xda, 0x2b, 0xbd, 0xfe, 0x2a, 0x01, 0xfe, 0x2a, 0x01, 0x21,
	0x2a, 0x00, 0xb2, 0xfe, 0x57, 0xb0, 0xb6, 0x57, 0xb0, 0x14, 0x74, 0xf4, 0xf0, 0x4c, 0x05, 0x68,
	0xfe, 0x54, 0x02, 0xfe, 0x54, 0x02, 0xd2, 0x54, 0x02, 0x00, 0xbe, 0xfe, 0x32, 0x0f, 0xb6, 0x32,
	0x0f, 0x14, 0x2e, 0x16, 0xaf, 0x47, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00,
	0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x94, 0x01, 0x5e, 0x5c, 0x20, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x07, 0xd0, 0xb1, 0x2f, 0xb2, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x57, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x54, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x46, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x42, 0x3c, 0x8a,
	0xb0, 0xe7, 0x54, 0x41, 0xb8, 0xca, 0x3a, 0x6f, 0xbe, 0xd8, 0x08, 0x00, 0x45, 0x00, 0x00, 0x34,
	0xa8, 0x23, 0x40, 0x00, 0x40, 0x06, 0x61, 0x7f, 0x0a, 0x01, 0x0e, 0x10, 0x0a, 0x01, 0x0f, 0x10,
	0x97, 0x91, 0x23, 0x84, 0x24, 0xfa, 0x91, 0xf7, 0xb4, 0xe8, 0xf3, 0x2d, 0x80, 0x10, 0x00, 0xab,
	0x7b, 0x7d, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xef, 0x1f, 0xf4, 0x36, 0xef, 0x1f, 0xdc, 0xde,
	0x00, 0x00, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x02, 0x02, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

// Test collected from the SFlow reference agent. Contains dummy data for several record types
// that wern't available on an actual network for sampling.
var SFlowTestPacket2 = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x45, 0x00,
	0x04, 0x88, 0x00, 0x00, 0x40, 0x00, 0x40, 0x11, 0x38, 0x63, 0x7f, 0x00, 0x00, 0x01, 0x7f, 0x00,
	0x00, 0x01, 0xdc, 0xb8, 0x18, 0xc7, 0x04, 0x74, 0x02, 0x88, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00,
	0x00, 0x01, 0xc0, 0xa8, 0x5b, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xb5, 0x3a, 0x00, 0x00,
	0xcb, 0x20, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01, 0x54, 0x00, 0x02,
	0x1f, 0x6e, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, 0x1f, 0x6e, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x3f, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00,
	0x03, 0xed, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x14, 0x68, 0x74,
	0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x6f,
	0x72, 0x67, 0x00, 0x00, 0x00, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x31, 0x2e, 0x73, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x6f, 0x72, 0x67, 0x06, 0x00, 0x00, 0x03, 0xec, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00,
	0x00, 0x6a, 0x00, 0x00, 0x00, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x75, 0x73, 0x65,
	0x72, 0xdc, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x00, 0x00, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x00, 0x00, 0x03, 0xeb, 0x00, 0x00,
	0x00, 0x64, 0x00, 0x00, 0x00, 0x01, 0x0d, 0x0c, 0x0b, 0x0a, 0x00, 0x00, 0xfd, 0xe9, 0x00, 0x00,
	0x00, 0x7b, 0x00, 0x00, 0x03, 0xe7, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x03, 0x00, 0x00, 0x00, 0x7b, 0x00, 0x00, 0x01, 0xc8, 0x00, 0x00, 0x03, 0x15, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x2b, 0x67, 0x00, 0x00, 0x56, 0xce, 0x00, 0x00,
	0x82, 0x35, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x03, 0x09, 0x00, 0x00,
	0x03, 0x78, 0x00, 0x00, 0x03, 0xe7, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00,
	0x00, 0x0d, 0x00, 0x00, 0x01, 0xb0, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x54, 0x00, 0x00,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x46, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x42, 0x00, 0x0c,
	0x29, 0x67, 0xa0, 0xe5, 0x00, 0x50, 0x56, 0xc0, 0x00, 0x09, 0x08, 0x00, 0x45, 0x10, 0x00, 0x34,
	0x92, 0xc3, 0x40, 0x00, 0x40, 0x06, 0x70, 0x8d, 0xc0, 0xa8, 0x5b, 0x01, 0xc0, 0xa8, 0x5b, 0x11,
	0xd3, 0xdd, 0x00, 0x16, 0xe3, 0x2e, 0x84, 0x77, 0x13, 0x6d, 0xc5, 0x53, 0x80, 0x10, 0x1f, 0xf7,
	0xe7, 0x7d, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x2e, 0xc6, 0x70, 0x3a, 0x00, 0x0f, 0x84, 0x7a,
	0xbc, 0xd2, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x01, 0x90, 0x00, 0x02, 0x1f, 0x6f, 0x00, 0x00,
	0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, 0x1f, 0x6f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0xff,
	0xff, 0xff, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x03, 0xed, 0x00, 0x00,
	0x00, 0x30, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x14, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f,
	0x2f, 0x77, 0x77, 0x77, 0x2e, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x6f, 0x72, 0x67, 0x00, 0x00,
	0x00, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x31, 0x2e, 0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x6f, 0x72,
	0x67, 0x03, 0x00, 0x00, 0x03, 0xec, 0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x00,
	0x00, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x77, 0x00, 0x00,
	0x00, 0x6a, 0x00, 0x00, 0x00, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x00, 0x00, 0x03, 0xeb, 0x00, 0x00, 0x00, 0x64, 0x00, 0x00,
	0x00, 0x01, 0x0d, 0x0c, 0x0b, 0x0a, 0x00, 0x00, 0xfd, 0xe9, 0x00, 0x00, 0x00, 0x7b, 0x00, 0x00,
	0x03, 0xe7, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
	0x00, 0x7b, 0x00, 0x00, 0x01, 0xc8, 0x00, 0x00, 0x03, 0x15, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0x03, 0x00, 0x00, 0x2b, 0x67, 0x00, 0x00, 0x56, 0xce, 0x00, 0x00, 0x82, 0x35, 0x00, 0x00,
	0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x03, 0x09, 0x00, 0x00, 0x03, 0x78, 0x00, 0x00,
	0x03, 0xe7, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00,
	0x01, 0xb0, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x01, 0x86, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x00, 0x50, 0x56, 0xc0, 0x00, 0x09,
	0x00, 0x0c, 0x29, 0x67, 0xa0, 0xe5, 0x08, 0x00, 0x45, 0x10, 0x01, 0x74, 0xbb, 0xfa, 0x40, 0x00,
	0x40, 0x06, 0x46, 0x16, 0xc0, 0xa8, 0x5b, 0x11, 0xc0, 0xa8, 0x5b, 0x01, 0x00, 0x16, 0xd3, 0xdd,
	0x13, 0x6d, 0xc5, 0x53, 0xe3, 0x2e, 0x84, 0x77, 0x80, 0x18, 0x01, 0x10, 0x38, 0xca, 0x00, 0x00,
	0x01, 0x01, 0x08, 0x0a, 0x00, 0x0f, 0x84, 0x7d, 0x2e, 0xc6, 0x70, 0x3a, 0xe3, 0x92, 0x97, 0x1a,
	0x67, 0x3b, 0xac, 0xec, 0xfa, 0x43, 0x71, 0x5e, 0x36, 0xa1, 0x0a, 0xc6, 0x1a, 0x6a, 0xed, 0x08,
	0xac, 0xf4, 0xbe, 0xd8, 0x36, 0x59, 0xf6, 0xe2, 0x3d, 0x34, 0x26, 0xf2, 0x42, 0xbd, 0x32, 0xd3,
	0x37, 0x52, 0xb8, 0xf4, 0x38, 0xf0, 0xf4, 0xeb, 0x76, 0x3b, 0xda, 0x23, 0xf1, 0x92, 0x96, 0xca,
	0xbb, 0x9c, 0x20, 0x0a, 0x38, 0x37, 0x6f, 0xd9, 0x26, 0xe6, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x01, 0x54, 0x00, 0x02, 0x1f, 0x70, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02,
	0x1f, 0x70, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x3f, 0xff, 0xff, 0xff, 0x00, 0x00,
	0x00, 0x04, 0x00, 0x00, 0x03, 0xed, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0x14, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x73, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x6f, 0x72, 0x67, 0x00, 0x00, 0x00, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x31, 0x2e,
	0x73, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x6f, 0x72, 0x67, 0xff, 0x00, 0x00, 0x03, 0xec, 0x00, 0x00,
	0x00, 0x2c, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x00, 0x00, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x77, 0x00, 0x00, 0x00, 0x6a, 0x00, 0x00, 0x00, 0x10, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x00, 0x00,
	0x03, 0xeb, 0x00, 0x00, 0x00, 0x64, 0x00, 0x00, 0x00, 0x01, 0x0d, 0x0c, 0x0b, 0x0a, 0x00, 0x00,
	0xfd, 0xe9, 0x00, 0x00, 0x00, 0x7b, 0x00, 0x00, 0x03, 0xe7, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
	0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x7b, 0x00, 0x00, 0x01, 0xc8, 0x00, 0x00,
	0x03, 0x15, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x2b, 0x67, 0x00, 0x00,
	0x56, 0xce, 0x00, 0x00, 0x82, 0x35, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
	0x03, 0x09, 0x00, 0x00, 0x03, 0x78, 0x00, 0x00, 0x03, 0xe7, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
	0x00, 0x0c, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x00, 0x01, 0xb0, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0x54, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x46, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00,
	0x00, 0x42, 0x00, 0x0c, 0x29, 0x67, 0xa0, 0xe5, 0x00, 0x50, 0x56, 0xc0, 0x00, 0x09, 0x08, 0x00,
	0x45, 0x10, 0x00, 0x34, 0x65, 0x7d, 0x40, 0x00, 0x40, 0x06, 0x9d, 0xd3, 0xc0, 0xa8, 0x5b, 0x01,
	0xc0, 0xa8, 0x5b, 0x11, 0xd3, 0xdd, 0x00, 0x16, 0xe3, 0x2e, 0x84, 0x77, 0x13, 0x6d, 0xc6, 0x93,
	0x80, 0x10, 0x1f, 0xec, 0xe6, 0x43, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x2e, 0xc6, 0x70, 0x3c,
	0x00, 0x0f, 0x84, 0x7d, 0x00, 0x50,
}

func TestDecodeUDPSFlow(t *testing.T) {
	p := gopacket.NewPacket(SFlowTestPacket1, LayerTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv4, LayerTypeUDP, LayerTypeSFlow}, t)
	if got, ok := p.TransportLayer().(*UDP); ok {
		want := &UDP{
			BaseLayer: BaseLayer{SFlowTestPacket1[34:42], SFlowTestPacket1[42:]},
			sPort:     []byte{199, 87},
			dPort:     []byte{24, 199},
			SrcPort:   51031,
			DstPort:   6343,
			Checksum:  8763,
			Length:    1448,
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("UDP layer mismatch, \nwant  %#v\ngot %#v\n", want, got)
		}
	} else {
		t.Error("Transport layer packet not UDP")
	}
}

func TestDecodeSFlowDatagram(t *testing.T) {
	p := gopacket.NewPacket(SFlowTestPacket1, LayerTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv4, LayerTypeUDP, LayerTypeSFlow}, t)
	if got, ok := p.ApplicationLayer().(*SFlowDatagram); ok {
		want := &SFlowDatagram{
			DatagramVersion: uint32(5),
			AgentAddress:    []byte{0xa, 0x1, 0xf8, 0x16},
			SubAgentID:      uint32(17),
			SequenceNumber:  uint32(40443),
			AgentUptime:     uint32(1078576845),
			SampleCount:     uint32(7),
			FlowSamples: []SFlowFlowSample{
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x2627e8,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x213,
					SamplingRate:    0x3e80,
					SamplePool:      0x50bde580,
					Dropped:         0x0,
					InputInterface:  0x213,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x5d2,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(SFlowTestPacket1[134:262], LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x15e5c1e,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x257,
					SamplingRate:    0x7d0,
					SamplePool:      0xb12fa290,
					Dropped:         0x0,
					InputInterface:  0x257,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x5ee,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(SFlowTestPacket1[350:478], LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x15acdd0,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x255,
					SamplingRate:    0x7d0,
					SamplePool:      0x9567e130,
					Dropped:         0x0,
					InputInterface:  0x255,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x446,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(SFlowTestPacket1[566:694], LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x15e5c1f,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x257,
					SamplingRate:    0x7d0,
					SamplePool:      0xb12faa60,
					Dropped:         0x0,
					InputInterface:  0x257,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x5ee,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(SFlowTestPacket1[958:1086], LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0xd0,
					SequenceNumber:  0x15acdd1,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x255,
					SamplingRate:    0x7d0,
					SamplePool:      0x9567e900,
					Dropped:         0x0,
					InputInterface:  0x255,
					OutputInterface: 0x257,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x90,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x5ee,
							PayloadRemoved: 0x4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(SFlowTestPacket1[1174:1302], LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x202,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    0x94,
					SequenceNumber:  0x15e5c20,
					SourceIDClass:   0x0,
					SourceIDIndex:   0x257,
					SamplingRate:    0x7d0,
					SamplePool:      0xb12fb230,
					Dropped:         0x0,
					InputInterface:  0x257,
					OutputInterface: 0x0,
					RecordCount:     0x2,
					Records: []SFlowRecord{
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x54,
							},
							HeaderProtocol: 0x1,
							FrameLength:    0x46,
							PayloadRemoved: 0x4,
							HeaderLength:   0x42,
							Header:         gopacket.NewPacket(SFlowTestPacket1[1390:1458], LayerTypeEthernet, gopacket.Default),
						},
						SFlowExtendedSwitchFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x3e9,
								FlowDataLength: 0x10,
							},
							IncomingVLAN:         0x202,
							IncomingVLANPriority: 0x0,
							OutgoingVLAN:         0x0,
							OutgoingVLANPriority: 0x0,
						},
					},
				},
			},
			CounterSamples: []SFlowCounterSample{
				SFlowCounterSample{
					Format:         0x2,
					SampleLength:   0xa8,
					SequenceNumber: 0x20f2,
					SourceIDClass:  0x0,
					SourceIDIndex:  0x20a,
					RecordCount:    0x2,
					Records: []SFlowRecord{
						SFlowGenericInterfaceCounters{
							SFlowBaseCounterRecord: SFlowBaseCounterRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 0x58,
							},
							IfIndex:            0x20a,
							IfType:             0x6,
							IfSpeed:            0x2540be400,
							IfDirection:        0x1,
							IfStatus:           0x3,
							IfInOctets:         0x129826db06c0b,
							IfInUcastPkts:      0xcb0ddd96,
							IfInMulticastPkts:  0x6a8c6,
							IfInBroadcastPkts:  0x7b,
							IfInDiscards:       0x0,
							IfInErrors:         0x0,
							IfInUnknownProtos:  0x0,
							IfOutOctets:        0x340235587c9e,
							IfOutUcastPkts:     0x56642571,
							IfOutMulticastPkts: 0x705ac4,
							IfOutBroadcastPkts: 0x908f1,
							IfOutDiscards:      0x0,
							IfOutErrors:        0x0,
							IfPromiscuousMode:  0x0,
						},
						SFlowEthernetCounters{
							SFlowBaseCounterRecord: SFlowBaseCounterRecord{
								EnterpriseID:   0x0,
								Format:         0x2,
								FlowDataLength: 0x34,
							},
							AlignmentErrors:           0x0,
							FCSErrors:                 0x0,
							SingleCollisionFrames:     0x0,
							MultipleCollisionFrames:   0x0,
							SQETestErrors:             0x0,
							DeferredTransmissions:     0x0,
							LateCollisions:            0x0,
							ExcessiveCollisions:       0x0,
							InternalMacTransmitErrors: 0x0,
							CarrierSenseErrors:        0x0,
							FrameTooLongs:             0x0,
							InternalMacReceiveErrors:  0x0,
							SymbolErrors:              0x0,
						},
					},
				},
			},
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("SFlow layer mismatch, \nwant:\n\n%#v\ngot:\n\n\n%#v\n\n", want, got)
		}
	} else {
		t.Error("Application layer packet not UDP")
	}
}

func TestPacketPacket0(t *testing.T) {
	p := gopacket.NewPacket(SFlowTestPacket2, LinkTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		t.Error("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{LayerTypeEthernet, LayerTypeIPv4, LayerTypeUDP, LayerTypeSFlow}, t)
	if got, ok := p.ApplicationLayer().(*SFlowDatagram); ok {
		want := &SFlowDatagram{
			DatagramVersion: uint32(5),
			AgentAddress:    []byte{192, 168, 91, 17},
			SubAgentID:      uint32(0),
			SequenceNumber:  uint32(46394),
			AgentUptime:     uint32(52000),
			SampleCount:     uint32(3),
			FlowSamples: []SFlowFlowSample{
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    340,
					SequenceNumber:  139118,
					SourceIDClass:   0,
					SourceIDIndex:   3,
					SamplingRate:    1,
					SamplePool:      139118,
					Dropped:         0,
					InputInterface:  3,
					OutputInterface: 1073741823,
					RecordCount:     4,
					Records: []SFlowRecord{
						SFlowExtendedURLRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1005,
								FlowDataLength: 48,
							},
							Direction: SFlowURLsrc,
							URL:       "http://www.sflow.org",
							Host:      "host1.sflow.org",
						},
						SFlowExtendedUserFlow{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1004,
								FlowDataLength: 44,
							},
							SourceCharSet:      SFlowCSUTF8,
							SourceUserID:       "source user",
							DestinationCharSet: SFlowCSUTF8,
							DestinationUserID:  "destination user",
						},
						SFlowExtendedGatewayFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1003,
								FlowDataLength: 100,
							},
							NextHop:     []byte{0x0d, 0x0c, 0x0b, 0x0a},
							AS:          65001,
							SourceAS:    123,
							PeerAS:      999,
							ASPathCount: 3,
							ASPath: []SFlowASDestination{
								SFlowASDestination{
									Type:    SFlowASSequence,
									Count:   3,
									Members: []uint32{123, 456, 789},
								},
								SFlowASDestination{
									Type:    SFlowASSet,
									Count:   3,
									Members: []uint32{11111, 22222, 33333},
								},
								SFlowASDestination{
									Type:    SFlowASSequence,
									Count:   3,
									Members: []uint32{777, 888, 999},
								},
							},
							Communities: []uint32{12, 13},
							LocalPref:   432,
						},
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 84,
							},
							HeaderProtocol: 1,
							FrameLength:    70,
							PayloadRemoved: 4,
							HeaderLength:   0x42,
							Header:         gopacket.NewPacket(SFlowTestPacket2[350:418], LayerTypeEthernet, gopacket.Default),
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    400,
					SequenceNumber:  139119,
					SourceIDClass:   0,
					SourceIDIndex:   3,
					SamplingRate:    1,
					SamplePool:      139119,
					Dropped:         0,
					InputInterface:  1073741823,
					OutputInterface: 3,
					RecordCount:     4,
					Records: []SFlowRecord{
						SFlowExtendedURLRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1005,
								FlowDataLength: 48,
							},
							Direction: SFlowURLsrc,
							URL:       "http://www.sflow.org",
							Host:      "host1.sflow.org",
						},
						SFlowExtendedUserFlow{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1004,
								FlowDataLength: 44,
							},
							SourceCharSet:      SFlowCSUTF8,
							SourceUserID:       "source user",
							DestinationCharSet: SFlowCSUTF8,
							DestinationUserID:  "destination user",
						},
						SFlowExtendedGatewayFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1003,
								FlowDataLength: 100,
							},
							NextHop:     []byte{0x0d, 0x0c, 0x0b, 0x0a},
							AS:          65001,
							SourceAS:    123,
							PeerAS:      999,
							ASPathCount: 3,
							ASPath: []SFlowASDestination{
								SFlowASDestination{
									Type:    SFlowASSequence,
									Count:   3,
									Members: []uint32{123, 456, 789},
								},
								SFlowASDestination{
									Type:    SFlowASSet,
									Count:   3,
									Members: []uint32{11111, 22222, 33333},
								},
								SFlowASDestination{
									Type:    SFlowASSequence,
									Count:   3,
									Members: []uint32{777, 888, 999},
								},
							},
							Communities: []uint32{12, 13},
							LocalPref:   432,
						},
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 144,
							},
							HeaderProtocol: 1,
							FrameLength:    390,
							PayloadRemoved: 4,
							HeaderLength:   0x80,
							Header:         gopacket.NewPacket(SFlowTestPacket2[698:826], LayerTypeEthernet, gopacket.Default),
						},
					},
				},
				SFlowFlowSample{
					EnterpriseID:    0x0,
					Format:          0x1,
					SampleLength:    340,
					SequenceNumber:  139120,
					SourceIDClass:   0,
					SourceIDIndex:   3,
					SamplingRate:    1,
					SamplePool:      139120,
					Dropped:         0,
					InputInterface:  3,
					OutputInterface: 1073741823,
					RecordCount:     4,
					Records: []SFlowRecord{
						SFlowExtendedURLRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1005,
								FlowDataLength: 48,
							},
							Direction: SFlowURLsrc,
							URL:       "http://www.sflow.org",
							Host:      "host1.sflow.org",
						},
						SFlowExtendedUserFlow{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1004,
								FlowDataLength: 44,
							},
							SourceCharSet:      SFlowCSUTF8,
							SourceUserID:       "source user",
							DestinationCharSet: SFlowCSUTF8,
							DestinationUserID:  "destination user",
						},
						SFlowExtendedGatewayFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0,
								Format:         1003,
								FlowDataLength: 100,
							},
							NextHop:     []byte{0x0d, 0x0c, 0x0b, 0x0a},
							AS:          65001,
							SourceAS:    123,
							PeerAS:      999,
							ASPathCount: 3,
							ASPath: []SFlowASDestination{
								SFlowASDestination{
									Type:    SFlowASSequence,
									Count:   3,
									Members: []uint32{123, 456, 789},
								},
								SFlowASDestination{
									Type:    SFlowASSet,
									Count:   3,
									Members: []uint32{11111, 22222, 33333},
								},
								SFlowASDestination{
									Type:    SFlowASSequence,
									Count:   3,
									Members: []uint32{777, 888, 999},
								},
							},
							Communities: []uint32{12, 13},
							LocalPref:   432,
						},
						SFlowRawPacketFlowRecord{
							SFlowBaseFlowRecord: SFlowBaseFlowRecord{
								EnterpriseID:   0x0,
								Format:         0x1,
								FlowDataLength: 84,
							},
							HeaderProtocol: 1,
							FrameLength:    70,
							PayloadRemoved: 4,
							HeaderLength:   0x42,
							Header:         gopacket.NewPacket(SFlowTestPacket2[1106:1174], LayerTypeEthernet, gopacket.Default),
						},
					},
				},
			},
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("SFlow layer mismatch, \nwant:\n\n%#v\ngot:\n\n\n%#v\n\n", want, got)
		}
	} else {
		t.Error("Application layer packet not UDP")
	}
}

func BenchmarkDecodeSFlowPacket1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(SFlowTestPacket1, LinkTypeEthernet, gopacket.NoCopy)
	}
}

func BenchmarkDecodeSFlowPacket2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopacket.NewPacket(SFlowTestPacket2, LinkTypeEthernet, gopacket.NoCopy)
	}
}

func BenchmarkDecodeSFlowLayerPacket1(b *testing.B) {
	var sflow SFlowDatagram
	for i := 0; i < b.N; i++ {
		sflow.DecodeFromBytes(SFlowTestPacket1[ /*eth*/ 14+ /*ipv4*/ 20+ /*udp*/ 8:], gopacket.NilDecodeFeedback)
	}
}

func BenchmarkDecodeSFlowLayerPacket2(b *testing.B) {
	var sflow SFlowDatagram
	for i := 0; i < b.N; i++ {
		sflow.DecodeFromBytes(SFlowTestPacket2[ /*eth*/ 14+ /*ipv4*/ 20+ /*udp*/ 8:], gopacket.NilDecodeFeedback)
	}
}
