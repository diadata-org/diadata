package types

import (
	"bytes"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/stretchr/testify/assert"
)

func TestBitVec_Decode(t *testing.T) {
	tests := []struct {
		ScaleEncodedVal []byte
		ExpectedResult  string
		BitOrder        BitOrder
	}{
		{
			ScaleEncodedVal: []byte{4, 0},
			ExpectedResult:  "0b00000000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{4, 1},
			ExpectedResult:  "0b10000000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{8, 0},
			ExpectedResult:  "0b00000000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{8, 2},
			ExpectedResult:  "0b01000000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{8, 1},
			ExpectedResult:  "0b10000000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{8, 3},
			ExpectedResult:  "0b11000000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{12, 5},
			ExpectedResult:  "0b10100000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{28, 106},
			ExpectedResult:  "0b01010110",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{28, 106},
			ExpectedResult:  "0b01010110",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{32, 106},
			ExpectedResult:  "0b01010110",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{36, 107, 1},
			ExpectedResult:  "0b1101011010000000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{60, 53, 53},
			ExpectedResult:  "0b1010110010101100",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{64, 106, 106},
			ExpectedResult:  "0b0101011001010110",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{68, 106, 106, 0},
			ExpectedResult:  "0b010101100101011000000000",
			BitOrder:        BitOrderLsb0,
		},
		{
			ScaleEncodedVal: []byte{4, 0},
			ExpectedResult:  "0b00000000",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{4, 128},
			ExpectedResult:  "0b10000000",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{8, 0},
			ExpectedResult:  "0b00000000",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{8, 128},
			ExpectedResult:  "0b10000000",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{8, 64},
			ExpectedResult:  "0b01000000",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{8, 192},
			ExpectedResult:  "0b11000000",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{12, 160},
			ExpectedResult:  "0b10100000",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{28, 86},
			ExpectedResult:  "0b01010110",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{32, 86},
			ExpectedResult:  "0b01010110",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{36, 214, 128},
			ExpectedResult:  "0b1101011010000000",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{60, 172, 172},
			ExpectedResult:  "0b1010110010101100",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{64, 86, 86},
			ExpectedResult:  "0b0101011001010110",
			BitOrder:        BitOrderMsb0,
		},
		{
			ScaleEncodedVal: []byte{68, 86, 86, 0},
			ExpectedResult:  "0b010101100101011000000000",
			BitOrder:        BitOrderMsb0,
		},
	}

	for _, test := range tests {
		b := NewBitVec(test.BitOrder)

		decoder := scale.NewDecoder(bytes.NewReader(test.ScaleEncodedVal))

		if err := b.Decode(*decoder); err != nil {
			assert.NoError(t, err)
			return
		}

		assert.Equal(t, test.ExpectedResult, b.String())
	}
}
