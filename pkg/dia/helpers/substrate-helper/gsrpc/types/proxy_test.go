package types_test

import (
	"math/big"
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
)

var (
	proxyDefinition1 = ProxyDefinition{
		Delegate:  newTestAccountID(),
		ProxyType: 6,
		Delay:     3,
	}
	proxyDefinition2 = ProxyDefinition{
		Delegate:  newTestAccountID(),
		ProxyType: 9,
		Delay:     0,
	}
)

func TestProxyDefinition_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[ProxyDefinition](t, 1000)
	AssertDecodeNilData[ProxyDefinition](t)
	AssertEncodeEmptyObj[ProxyDefinition](t, 37)
}

func TestProxyDefinition_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			proxyDefinition1,
			MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200603000000"),
		},
		{
			proxyDefinition2,
			MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200900000000"),
		},
	})
}

func TestProxyDefinition_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200603000000"),
			proxyDefinition1,
		},
		{
			MustHexDecodeString("0x0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200900000000"),
			proxyDefinition2,
		},
	})
}

var (
	proxyStorageEntry1 = ProxyStorageEntry{
		ProxyDefinitions: []ProxyDefinition{
			proxyDefinition1,
			proxyDefinition2,
		},
		Balance: NewU128(*big.NewInt(1234)),
	}
)

func TestProxyStorageEntry_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[ProxyStorageEntry](t, 1000)
	AssertEncodeEmptyObj[ProxyStorageEntry](t, 17)
}

func TestProxyStorageEntry_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			proxyStorageEntry1,
			MustHexDecodeString("0x080102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f2006030000000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200900000000d2040000000000000000000000000000"),
		},
	})
}

func TestProxyStorageEntry_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x080102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f2006030000000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f200900000000d2040000000000000000000000000000"),
			proxyStorageEntry1,
		},
	})
}
