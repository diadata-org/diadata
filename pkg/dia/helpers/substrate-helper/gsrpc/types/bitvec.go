package types

import (
	"errors"
	"fmt"
	"math"
	"math/bits"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
)

type BitOrder uint

const (
	BitOrderLsb0 BitOrder = iota
	BitOrderMsb0
)

var (
	BitOrderName = map[BitOrder]string{
		BitOrderLsb0: "Lsb0",
		BitOrderMsb0: "Msb0",
	}

	BitOrderValue = map[string]BitOrder{
		"Lsb0": BitOrderLsb0,
		"Msb0": BitOrderMsb0,
	}
)

func (b *BitOrder) String() string {
	return BitOrderName[*b]
}

func NewBitOrderFromString(s string) (BitOrder, error) {
	bitOrder, ok := BitOrderValue[s]

	if !ok {
		return 0, fmt.Errorf("bit order '%s' not supported", s)
	}

	return bitOrder, nil
}

type BitVec struct {
	BitOrder BitOrder

	ByteSlice []uint8
}

func NewBitVec(bitOrder BitOrder) *BitVec {
	return &BitVec{
		BitOrder: bitOrder,
	}
}

func (b *BitVec) Decode(decoder scale.Decoder) error {
	numBytes, err := b.GetMinimumNumberOfBytes(decoder)

	if err != nil {
		return err
	}

	for i := uint(0); i < numBytes; i++ {
		decodedByte, err := decoder.ReadOneByte()

		if err != nil {
			return err
		}

		if b.BitOrder == BitOrderLsb0 {
			decodedByte = bits.Reverse8(decodedByte)
		}

		b.ByteSlice = append(b.ByteSlice, decodedByte)
	}

	return nil
}

func (b *BitVec) GetMinimumNumberOfBytes(decoder scale.Decoder) (uint, error) {
	nb, err := decoder.DecodeUintCompact()

	if err != nil {
		return 0, err
	}

	numberOfBits := nb.Uint64()

	if numberOfBits == 0 {
		return 0, errors.New("invalid number of bits")
	}

	return uint(math.Ceil(float64(numberOfBits) / 8)), nil
}

func (b *BitVec) String() string {
	var sb strings.Builder

	sb.WriteString("0b")

	for _, byte := range b.ByteSlice {
		sb.WriteString(fmt.Sprintf("%08b", byte))
	}

	return sb.String()
}
