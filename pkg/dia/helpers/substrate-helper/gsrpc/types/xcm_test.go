// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types_test

import (
	"math/big"
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
)

var (
	testAssetID1 = AssetID{
		Parents:  1,
		Interior: testJunctionsV1n1,
	}
	testAssetID2 = AssetID{
		Parents:  1,
		Interior: testJunctionsV1n2,
	}

	assetIDFuzzOpts = CombineFuzzOpts(
		junctionsV1FuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(a *AssetID, c fuzz.Continue) {
				c.Fuzz(&a.Parents)
				c.Fuzz(&a.Interior)
			}),
		},
	)
)

func TestAssetID_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[AssetID](t, 100, assetIDFuzzOpts...)
	AssertDecodeNilData[AssetID](t)
	AssertEncodeEmptyObj[AssetID](t, 1)
}

func TestAssetID_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testAssetID1, MustHexDecodeString("0x0100")},
		{testAssetID2, MustHexDecodeString("0x0101002c")},
	})
}

func TestAssetID_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x0100"), testAssetID1},
		{MustHexDecodeString("0x0101002c"), testAssetID2},
	})
}

var (
	testAssetInstance1 = AssetInstance{
		IsUndefined: true,
	}
	testAssetInstance2 = AssetInstance{
		IsIndex: true,
		Index:   NewU128(*big.NewInt(123)),
	}
	testAssetInstance3 = AssetInstance{
		IsArray4: true,
		Array4:   [4]U8{1, 2, 3, 4},
	}
	testAssetInstance4 = AssetInstance{
		IsArray8: true,
		Array8:   [8]U8{6, 7, 4, 3},
	}
	testAssetInstance5 = AssetInstance{
		IsArray16: true,
		Array16:   [16]U8{5, 6, 7, 9, 4, 3},
	}
	testAssetInstance6 = AssetInstance{
		IsArray32: true,
		Array32:   [32]U8{4, 5, 43, 3, 2},
	}
	testAssetInstance7 = AssetInstance{
		IsBlob: true,
		Blob:   []U8{4, 5, 3, 2, 4, 5, 6, 2, 1},
	}

	assetInstanceFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(a *AssetInstance, c fuzz.Continue) {
			switch c.Intn(7) {
			case 0:
				a.IsUndefined = true
			case 1:
				a.IsIndex = true

				c.Fuzz(&a.Index)
			case 2:
				a.IsArray4 = true

				c.Fuzz(&a.Array4)
			case 3:
				a.IsArray8 = true

				c.Fuzz(&a.Array8)
			case 4:
				a.IsArray16 = true

				c.Fuzz(&a.Array16)
			case 5:
				a.IsArray32 = true

				c.Fuzz(&a.Array32)
			case 6:
				a.IsBlob = true

				c.Fuzz(&a.Blob)
			}
		}),
	}
)

func TestAssetInstance_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[AssetInstance](t, 1000, assetInstanceFuzzOpts...)
	AssertDecodeNilData[AssetInstance](t)
	AssertEncodeEmptyObj[AssetInstance](t, 0)
}

func TestAssetInstance_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testAssetInstance1, MustHexDecodeString("0x00")},
		{testAssetInstance2, MustHexDecodeString("0x017b000000000000000000000000000000")},
		{testAssetInstance3, MustHexDecodeString("0x0201020304")},
		{testAssetInstance4, MustHexDecodeString("0x030607040300000000")},
		{testAssetInstance5, MustHexDecodeString("0x0405060709040300000000000000000000")},
		{testAssetInstance6, MustHexDecodeString("0x0504052b0302000000000000000000000000000000000000000000000000000000")},
		{testAssetInstance7, MustHexDecodeString("0x0624040503020405060201")},
	})
}

func TestAssetInstance_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testAssetInstance1},
		{MustHexDecodeString("0x017b000000000000000000000000000000"), testAssetInstance2},
		{MustHexDecodeString("0x0201020304"), testAssetInstance3},
		{MustHexDecodeString("0x030607040300000000"), testAssetInstance4},
		{MustHexDecodeString("0x0405060709040300000000000000000000"), testAssetInstance5},
		{MustHexDecodeString("0x0504052b0302000000000000000000000000000000000000000000000000000000"), testAssetInstance6},
		{MustHexDecodeString("0x0624040503020405060201"), testAssetInstance7},
	})
}

var (
	testFungibility1 = Fungibility{
		IsFungible: true,
		Amount:     NewUCompactFromUInt(123),
	}
	testFungibility2 = Fungibility{
		IsNonFungible: true,
		AssetInstance: testAssetInstance3,
	}

	fungibilityFuzzOpts = CombineFuzzOpts(
		assetInstanceFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(f *Fungibility, c fuzz.Continue) {
				if c.RandBool() {
					f.IsFungible = true
					c.Fuzz(&f.Amount)
					return
				}

				f.IsNonFungible = true
				c.Fuzz(&f.AssetInstance)
			}),
		},
	)
)

func TestFungibility_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Fungibility](t, 1000, fungibilityFuzzOpts...)
	AssertDecodeNilData[Fungibility](t)
	AssertEncodeEmptyObj[Fungibility](t, 0)
}

func TestFungibility_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testFungibility1, MustHexDecodeString("0x00ed01")},
		{testFungibility2, MustHexDecodeString("0x010201020304")},
	})
}

func TestFungibility_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00ed01"), testFungibility1},
		{MustHexDecodeString("0x010201020304"), testFungibility2},
	})
}

var (
	testMultiAssetV1 = MultiAssetV1{
		ID:          testAssetID2,
		Fungibility: testFungibility1,
	}
	testMultiAssetV2 = MultiAssetV1{
		ID:          testAssetID1,
		Fungibility: testFungibility2,
	}

	multiAssetV1FuzzOpts = CombineFuzzOpts(assetIDFuzzOpts, fungibilityFuzzOpts)
)

func TestMultiAssetV1_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[MultiAssetV1](t, 1000, multiAssetV1FuzzOpts...)
	AssertDecodeNilData[MultiAssetV1](t)
	AssertEncodeEmptyObj[MultiAssetV1](t, 1)
}

func TestMultiAssetV1_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testMultiAssetV1,
			MustHexDecodeString("0x0101002c00ed01"),
		},
		{
			testMultiAssetV2,
			MustHexDecodeString("0x0100010201020304"),
		},
	})
}

func TestMultiAssetV1_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x0101002c00ed01"),
			testMultiAssetV1,
		},
		{
			MustHexDecodeString("0x0100010201020304"),
			testMultiAssetV2,
		},
	})
}

var (
	testMultiAssetsV1 = MultiAssetsV1{testMultiAssetV1, testMultiAssetV2}
)

func TestMultiAssetsV1_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[MultiAssetsV1](t, 100, multiAssetV1FuzzOpts...)
	AssertEncodeEmptyObj[MultiAssetsV1](t, 1)
}

func TestMultiAssetsV1_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testMultiAssetsV1,
			MustHexDecodeString("0x080101002c00ed010100010201020304"),
		},
	})
}

func TestMultiAssetsV1_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x080101002c00ed010100010201020304"),
			testMultiAssetsV1,
		},
	})
}

var (
	testMultiAssetV0n1 = MultiAssetV0{
		IsNone: true,
	}
	testMultiAssetV0n2 = MultiAssetV0{
		IsAll: true,
	}
	testMultiAssetV0n3 = MultiAssetV0{
		IsAllFungible: true,
	}
	testMultiAssetV0n4 = MultiAssetV0{
		IsAllNonFungible: true,
	}
	testMultiAssetV0n5 = MultiAssetV0{
		IsAllAbstractFungible: true,
		AllAbstractFungibleID: []U8{4, 5, 6},
	}
	testMultiAssetV0n6 = MultiAssetV0{
		IsAllAbstractNonFungible:    true,
		AllAbstractNonFungibleClass: []U8{5, 6, 7, 8},
	}
	testMultiAssetV0n7 = MultiAssetV0{
		IsAllConcreteFungible: true,
		AllConcreteFungibleID: testMultiLocationV1n1,
	}
	testMultiAssetV0n8 = MultiAssetV0{
		IsAllConcreteNonFungible:    true,
		AllConcreteNonFungibleClass: testMultiLocationV1n1,
	}
	testMultiAssetV0n9 = MultiAssetV0{
		IsAbstractFungible: true,
		AbstractFungibleID: []U8{12, 4, 3},
		AbstractFungible:   NewU128(*big.NewInt(432)),
	}
	testMultiAssetV0n10 = MultiAssetV0{
		IsAbstractNonFungible:       true,
		AbstractNonFungibleClass:    []U8{7, 6, 5},
		AbstractNonFungibleInstance: testAssetInstance4,
	}
	testMultiAssetV0n11 = MultiAssetV0{
		IsConcreteFungible:     true,
		ConcreteFungibleID:     testMultiLocationV1n1,
		ConcreteFungibleAmount: NewU128(*big.NewInt(654)),
	}
	testMultiAssetV0n12 = MultiAssetV0{
		IsConcreteNonFungible:       true,
		ConcreteNonFungibleClass:    testMultiLocationV1n1,
		ConcreteNonFungibleInstance: testAssetInstance3,
	}

	multiAssetV0FuzzOpts = CombineFuzzOpts(
		multiLocationV1FuzzOpts,
		assetInstanceFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(m *MultiAssetV0, c fuzz.Continue) {
				switch c.Intn(12) {
				case 0:
					m.IsNone = true
				case 1:
					m.IsAll = true
				case 2:
					m.IsAllFungible = true
				case 3:
					m.IsAllNonFungible = true
				case 4:
					m.IsAllAbstractFungible = true

					c.Fuzz(&m.AllAbstractFungibleID)
				case 5:
					m.IsAllAbstractNonFungible = true

					c.Fuzz(&m.AllAbstractNonFungibleClass)
				case 6:
					m.IsAllConcreteFungible = true

					c.Fuzz(&m.AllConcreteFungibleID)
				case 7:
					m.IsAllConcreteNonFungible = true

					c.Fuzz(&m.AllConcreteNonFungibleClass)
				case 8:
					m.IsAbstractFungible = true

					c.Fuzz(&m.AbstractFungibleID)

					c.Fuzz(&m.AbstractFungible)
				case 9:
					m.IsAbstractNonFungible = true

					c.Fuzz(&m.AbstractNonFungibleClass)

					c.Fuzz(&m.AbstractNonFungibleInstance)
				case 10:
					m.IsConcreteFungible = true

					c.Fuzz(&m.ConcreteFungibleID)

					c.Fuzz(&m.ConcreteFungibleAmount)
				case 11:
					m.IsConcreteNonFungible = true

					c.Fuzz(&m.ConcreteNonFungibleClass)

					c.Fuzz(&m.ConcreteNonFungibleInstance)
				}
			}),
		},
	)
)

func TestMultiAssetV0_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[MultiAssetV0](t, 1000, multiAssetV0FuzzOpts...)
	AssertDecodeNilData[MultiAssetV0](t)
	AssertEncodeEmptyObj[MultiAssetV0](t, 0)
}

func TestMultiAssetV0_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testMultiAssetV0n1, MustHexDecodeString("0x00")},
		{testMultiAssetV0n2, MustHexDecodeString("0x01")},
		{testMultiAssetV0n3, MustHexDecodeString("0x02")},
		{testMultiAssetV0n4, MustHexDecodeString("0x03")},
		{testMultiAssetV0n5, MustHexDecodeString("0x040c040506")},
		{testMultiAssetV0n6, MustHexDecodeString("0x051005060708")},
		{testMultiAssetV0n7, MustHexDecodeString("0x060408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807")},
		{testMultiAssetV0n8, MustHexDecodeString("0x070408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807")},
		{testMultiAssetV0n9, MustHexDecodeString("0x080c0c0403b0010000000000000000000000000000")},
		{testMultiAssetV0n10, MustHexDecodeString("0x090c070605030607040300000000")},
		{testMultiAssetV0n11, MustHexDecodeString("0x0a0408002c01000c010203020010000000000000000303000404052a00000000000000000000000000000006080608078e020000000000000000000000000000")},
		{testMultiAssetV0n12, MustHexDecodeString("0x0b0408002c01000c010203020010000000000000000303000404052a00000000000000000000000000000006080608070201020304")},
	})
}

func TestMultiAssetV0_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testMultiAssetV0n1},
		{MustHexDecodeString("0x01"), testMultiAssetV0n2},
		{MustHexDecodeString("0x02"), testMultiAssetV0n3},
		{MustHexDecodeString("0x03"), testMultiAssetV0n4},
		{MustHexDecodeString("0x040c040506"), testMultiAssetV0n5},
		{MustHexDecodeString("0x051005060708"), testMultiAssetV0n6},
		{MustHexDecodeString("0x060408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807"), testMultiAssetV0n7},
		{MustHexDecodeString("0x070408002c01000c010203020010000000000000000303000404052a0000000000000000000000000000000608060807"), testMultiAssetV0n8},
		{MustHexDecodeString("0x080c0c0403b0010000000000000000000000000000"), testMultiAssetV0n9},
		{MustHexDecodeString("0x090c070605030607040300000000"), testMultiAssetV0n10},
		{MustHexDecodeString("0x0a0408002c01000c010203020010000000000000000303000404052a00000000000000000000000000000006080608078e020000000000000000000000000000"), testMultiAssetV0n11},
		{MustHexDecodeString("0x0b0408002c01000c010203020010000000000000000303000404052a00000000000000000000000000000006080608070201020304"), testMultiAssetV0n12},
	})
}

var (
	testVersionedMultiAssets1 = VersionedMultiAssets{
		IsV0: true,
		MultiAssetsV0: []MultiAssetV0{
			testMultiAssetV0n1,
			testMultiAssetV0n5,
			testMultiAssetV0n11,
		},
	}
	testVersionedMultiAssets2 = VersionedMultiAssets{
		IsV1:          true,
		MultiAssetsV1: testMultiAssetsV1,
	}

	versionedMultiAssetsFuzzOpts = CombineFuzzOpts(
		multiAssetV0FuzzOpts,
		multiAssetV1FuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(v *VersionedMultiAssets, c fuzz.Continue) {
				if c.RandBool() {
					v.IsV0 = true
					c.Fuzz(&v.MultiAssetsV0)
					return
				}

				v.IsV1 = true
				c.Fuzz(&v.MultiAssetsV1)
			}),
		},
	)
)

func TestVersionedMultiAssets_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[VersionedMultiAssets](t, 1000, versionedMultiAssetsFuzzOpts...)
	AssertDecodeNilData[VersionedMultiAssets](t)
	AssertEncodeEmptyObj[VersionedMultiAssets](t, 0)
}

func TestVersionedMultiAssets_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testVersionedMultiAssets1,
			MustHexDecodeString("0x000c00040c0405060a0408002c01000c010203020010000000000000000303000404052a00000000000000000000000000000006080608078e020000000000000000000000000000"),
		},
		{
			testVersionedMultiAssets2,
			MustHexDecodeString("0x01080101002c00ed010100010201020304"),
		},
	})
}

func TestVersionedMultiAssets_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x000c00040c0405060a0408002c01000c010203020010000000000000000303000404052a00000000000000000000000000000006080608078e020000000000000000000000000000"),
			testVersionedMultiAssets1,
		},
		{
			MustHexDecodeString("0x01080101002c00ed010100010201020304"),
			testVersionedMultiAssets2,
		},
	})
}

var (
	testResponse1 = Response{
		IsNull: true,
	}
	testResponse2 = Response{
		IsAssets:    true,
		MultiAssets: testMultiAssetsV1,
	}
	testResponse3 = Response{
		IsExecutionResult: true,
		ExecutionResult:   testExecutionResult,
	}
	testResponse4 = Response{
		IsVersion: true,
		Version:   NewU32(431),
	}

	responseFuzzOpts = CombineFuzzOpts(
		multiAssetV1FuzzOpts,
		executionResultFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(r *Response, c fuzz.Continue) {
				switch c.Intn(4) {
				case 0:
					r.IsNull = true
				case 1:
					r.IsAssets = true

					c.Fuzz(&r.MultiAssets)
				case 2:
					r.IsExecutionResult = true

					c.Fuzz(&r.ExecutionResult)
				case 3:
					r.IsVersion = true

					c.Fuzz(&r.Version)
				}
			}),
		},
	)
)

func TestResponse_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Response](t, 1000, responseFuzzOpts...)
	AssertDecodeNilData[Response](t)
	AssertEncodeEmptyObj[Response](t, 0)
}

func TestResponse_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testResponse1, MustHexDecodeString("0x00")},
		{testResponse2, MustHexDecodeString("0x01080101002c00ed010100010201020304")},
		{testResponse3, MustHexDecodeString("0x020100000000")},
		{testResponse4, MustHexDecodeString("0x03af010000")},
	})
}

func TestResponse_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testResponse1},
		{MustHexDecodeString("0x01080101002c00ed010100010201020304"), testResponse2},
		{MustHexDecodeString("0x020100000000"), testResponse3},
		{MustHexDecodeString("0x03af010000"), testResponse4},
	})
}

var (
	testOriginKind1 = OriginKind{
		IsNative: true,
	}
	testOriginKind2 = OriginKind{
		IsSovereignAccount: true,
	}
	testOriginKind3 = OriginKind{
		IsSuperuser: true,
	}
	testOriginKind4 = OriginKind{
		IsXcm: true,
	}

	originKindFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(o *OriginKind, c fuzz.Continue) {
			switch c.Intn(4) {
			case 0:
				o.IsNative = true
			case 1:
				o.IsSovereignAccount = true
			case 2:
				o.IsSuperuser = true
			case 3:
				o.IsXcm = true
			}
		}),
	}
)

func TestOriginKind_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[OriginKind](t, 1000, originKindFuzzOpts...)
	AssertDecodeNilData[OriginKind](t)
	AssertEncodeEmptyObj[OriginKind](t, 0)
}

func TestOriginKind_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testOriginKind1, MustHexDecodeString("0x00")},
		{testOriginKind2, MustHexDecodeString("0x01")},
		{testOriginKind3, MustHexDecodeString("0x02")},
		{testOriginKind4, MustHexDecodeString("0x03")},
	})
}

func TestOriginKind_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testOriginKind1},
		{MustHexDecodeString("0x01"), testOriginKind2},
		{MustHexDecodeString("0x02"), testOriginKind3},
		{MustHexDecodeString("0x03"), testOriginKind4},
	})
}

var (
	testEncodedCall = EncodedCall{
		Call: []U8{5, 6, 7, 9},
	}
)

func TestEncodedCall_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[EncodedCall](t, 100)
	AssertEncodeEmptyObj[EncodedCall](t, 1)
}

func TestEncodedCall_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testEncodedCall, MustHexDecodeString("0x1005060709")},
	})
}

func TestEncodedCall_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x1005060709"), testEncodedCall},
	})
}

var (
	testWildFungibility1 = WildFungibility{IsFungible: true}
	testWildFungibility2 = WildFungibility{IsNonFungible: true}

	wildFungibilityFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(w *WildFungibility, c fuzz.Continue) {
			if c.RandBool() {
				w.IsFungible = true
				return
			}

			w.IsNonFungible = true
		}),
	}
)

func TestWildFungibility_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[WildFungibility](t, 100, wildFungibilityFuzzOpts...)
	AssertDecodeNilData[WildFungibility](t)
	AssertEncodeEmptyObj[WildFungibility](t, 0)
}

func TestWildFungibility_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testWildFungibility1, MustHexDecodeString("0x00")},
		{testWildFungibility2, MustHexDecodeString("0x01")},
	})
}

func TestWildFungibility_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testWildFungibility1},
		{MustHexDecodeString("0x01"), testWildFungibility2},
	})
}

var (
	testWildMultiAsset1 = WildMultiAsset{
		IsAll: true,
	}
	testWildMultiAsset2 = WildMultiAsset{
		IsAllOf: true,
		ID:      testAssetID2,
		Fun:     testWildFungibility1,
	}

	wildMultiAssetFuzzOpts = CombineFuzzOpts(
		assetIDFuzzOpts,
		wildFungibilityFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(w *WildMultiAsset, c fuzz.Continue) {
				if c.RandBool() {
					w.IsAll = true
					return
				}

				w.IsAllOf = true
				c.Fuzz(&w.ID)
				c.Fuzz(&w.Fun)
			}),
		},
	)
)

func TestWildMultiAsset_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[WildMultiAsset](t, 100, wildMultiAssetFuzzOpts...)
	AssertDecodeNilData[WildMultiAsset](t)
	AssertEncodeEmptyObj[WildMultiAsset](t, 0)
}

func TestWildMultiAsset_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testWildMultiAsset1, MustHexDecodeString("0x00")},
		{testWildMultiAsset2, MustHexDecodeString("0x010101002c00")},
	})
}

func TestWildMultiAsset_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testWildMultiAsset1},
		{MustHexDecodeString("0x010101002c00"), testWildMultiAsset2},
	})
}

var (
	testMultiAssetFilter1 = MultiAssetFilter{
		IsDefinite:  true,
		MultiAssets: testMultiAssetsV1,
	}
	testMultiAssetFilter2 = MultiAssetFilter{
		IsWild:         true,
		WildMultiAsset: testWildMultiAsset1,
	}
	multiAssetFilterFuzzOpts = CombineFuzzOpts(
		multiAssetV1FuzzOpts,
		wildMultiAssetFuzzOpts,
		[]FuzzOpt{
			WithFuzzFuncs(func(m *MultiAssetFilter, c fuzz.Continue) {
				if c.RandBool() {
					m.IsDefinite = true
					c.Fuzz(&m.MultiAssets)
					return
				}

				m.IsWild = true
				c.Fuzz(&m.WildMultiAsset)
			}),
		},
	)
)

func TestMultiAssetFilter_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[MultiAssetFilter](t, 100, multiAssetFilterFuzzOpts...)
	AssertDecodeNilData[MultiAssetFilter](t)
	AssertEncodeEmptyObj[MultiAssetFilter](t, 0)
}

func TestMultiAssetFilter_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{
			testMultiAssetFilter1,
			MustHexDecodeString("0x00080101002c00ed010100010201020304"),
		},
		{
			testMultiAssetFilter2,
			MustHexDecodeString("0x0100"),
		},
	})
}

func TestMultiAssetFilter_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{
			MustHexDecodeString("0x00080101002c00ed010100010201020304"),
			testMultiAssetFilter1,
		},
		{
			MustHexDecodeString("0x0100"),
			testMultiAssetFilter2,
		},
	})
}

var (
	testWeightLimit1 = WeightLimit{
		IsUnlimited: true,
	}
	testWeightLimit2 = WeightLimit{
		IsLimited: true,
		Limit:     54,
	}
	weightLimitFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(w *WeightLimit, c fuzz.Continue) {
			if c.RandBool() {
				w.IsUnlimited = true
				return
			}

			w.IsLimited = true
			c.Fuzz(&w.Limit)
		}),
	}
)

func TestWeightLimit_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[WeightLimit](t, 100, weightLimitFuzzOpts...)
	AssertDecodeNilData[WeightLimit](t)
	AssertEncodeEmptyObj[WeightLimit](t, 0)
}

func TestWeightLimit_Encode(t *testing.T) {
	AssertEncode(t, []EncodingAssert{
		{testWeightLimit1, MustHexDecodeString("0x00")},
		{testWeightLimit2, MustHexDecodeString("0x013600000000000000")},
	})
}

func TestWeightLimit_Decode(t *testing.T) {
	AssertDecode(t, []DecodingAssert{
		{MustHexDecodeString("0x00"), testWeightLimit1},
		{MustHexDecodeString("0x013600000000000000"), testWeightLimit2},
	})
}

var (
	instructionFuzzOpts = CombineFuzzOpts(
		multiAssetV1FuzzOpts,
		responseFuzzOpts,
		multiLocationV1FuzzOpts,
		originKindFuzzOpts,
		multiAssetFilterFuzzOpts,
		weightLimitFuzzOpts,
		[]FuzzOpt{
			WithNumElements(1, 3),
			WithFuzzFuncs(func(i *Instruction, c fuzz.Continue) {
				switch c.Intn(28) {
				case 0:
					i.IsWithdrawAsset = true

					c.Fuzz(&i.WithdrawAssetMultiAssets)
				case 1:
					i.IsReserveAssetDeposited = true

					c.Fuzz(&i.ReserveAssetDepositedMultiAssets)
				case 2:
					i.IsReceiveTeleportedAsset = true

					c.Fuzz(&i.ReceiveTeleportedAssetMultiAssets)
				case 3:
					i.IsQueryResponse = true

					c.Fuzz(&i.QueryResponseQueryID)

					c.Fuzz(&i.QueryResponseResponse)

					c.Fuzz(&i.QueryResponseMaxWeight)
				case 4:
					i.IsTransferAsset = true

					c.Fuzz(&i.TransferAssetAssets)

					c.Fuzz(&i.TransferAssetBeneficiary)
				case 5:
					i.IsTransferReserveAsset = true

					c.Fuzz(&i.TransferReserveAssetMultiAssets)

					c.Fuzz(&i.TransferReserveAssetDest)

					c.Fuzz(&i.TransferReserveAssetXCM)
				case 6:
					i.IsTransact = true

					c.Fuzz(&i.TransactOriginType)

					c.Fuzz(&i.TransactRequireWeightAtMost)

					c.Fuzz(&i.TransactCall)
				case 7:
					i.IsHrmpNewChannelOpenRequest = true

					c.Fuzz(&i.HrmpNewChannelOpenRequestSender)

					c.Fuzz(&i.HrmpNewChannelOpenRequestMaxMessageSize)

					c.Fuzz(&i.HrmpNewChannelOpenRequestMaxCapacity)
				case 8:
					i.IsHrmpChannelAccepted = true

					c.Fuzz(&i.HrmpChannelAcceptedRecipient)
				case 9:
					i.IsHrmpChannelClosing = true

					c.Fuzz(&i.HrmpChannelClosingInitiator)

					c.Fuzz(&i.HrmpChannelClosingSender)

					c.Fuzz(&i.HrmpChannelClosingRecipient)
				case 10:
					i.IsClearOrigin = true
				case 11:
					i.IsDescendOrigin = true

					c.Fuzz(&i.DescendOriginLocation)
				case 12:
					i.IsReportError = true

					c.Fuzz(&i.ReportErrorQueryID)

					c.Fuzz(&i.ReportErrorDestination)

					c.Fuzz(&i.ReportErrorMaxResponseWeight)
				case 13:
					i.IsDepositAsset = true

					c.Fuzz(&i.DepositAssetMultiAssetFilter)

					c.Fuzz(&i.DepositAssetMaxAssets)

					c.Fuzz(&i.DepositAssetBeneficiary)
				case 14:
					i.IsDepositReserveAsset = true

					c.Fuzz(&i.DepositReserveAssetMultiAssetFilter)

					c.Fuzz(&i.DepositReserveAssetMaxAssets)

					c.Fuzz(&i.DepositReserveAssetDest)

					c.Fuzz(&i.DepositReserveAssetXCM)
				case 15:
					i.IsExchangeAsset = true

					c.Fuzz(&i.ExchangeAssetGive)

					c.Fuzz(&i.ExchangeAssetReceive)
				case 16:
					i.IsInitiateReserveWithdraw = true

					c.Fuzz(&i.InitiateReserveWithdrawAssets)

					c.Fuzz(&i.InitiateReserveWithdrawReserve)

					c.Fuzz(&i.InitiateReserveWithdrawXCM)
				case 17:
					i.IsInitiateTeleport = true

					c.Fuzz(&i.InitiateTeleportAssets)

					c.Fuzz(&i.InitiateTeleportDest)

					c.Fuzz(&i.InitiateTeleportXCM)
				case 18:
					i.IsQueryHolding = true

					c.Fuzz(&i.QueryHoldingQueryID)

					c.Fuzz(&i.QueryHoldingDest)

					c.Fuzz(&i.QueryHoldingAssets)

					c.Fuzz(&i.QueryHoldingMaxResponseWeight)
				case 19:
					i.IsBuyExecution = true

					c.Fuzz(&i.BuyExecutionFees)

					c.Fuzz(&i.BuyExecutionWeightLimit)
				case 20:
					i.IsRefundSurplus = true
				case 21:
					i.IsSetErrorHandler = true

					c.Fuzz(&i.SetErrorHandlerXCM)
				case 22:
					i.IsSetAppendix = true

					c.Fuzz(&i.SetAppendixXCM)
				case 23:
					i.IsClearError = true
				case 24:
					i.IsClaimAsset = true

					c.Fuzz(&i.ClaimAssetAssets)

					c.Fuzz(&i.ClaimAssetTicket)
				case 25:
					i.IsTrap = true

					c.Fuzz(&i.TrapCode)
				case 26:
					i.IsSubscribeVersion = true

					c.Fuzz(&i.SubscribeVersionQueryID)

					c.Fuzz(&i.SubscribeVersionMaxResponseWeight)
				case 27:
					i.IsUnsubscribeVersion = true
				}
			}),
		},
	)
)

func TestInstruction_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[Instruction](t, 1000, instructionFuzzOpts...)
	AssertDecodeNilData[Instruction](t)
	AssertEncodeEmptyObj[Instruction](t, 0)
}
