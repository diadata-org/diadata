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
	"testing"

	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/test_utils"
	fuzz "github.com/google/gofuzz"
)

var (
	xcmErrorFuzzOpts = []FuzzOpt{
		WithFuzzFuncs(func(x *XCMError, c fuzz.Continue) {
			switch c.Intn(26) {
			case 0:
				x.IsOverflow = true
			case 1:
				x.IsUnimplemented = true
			case 2:
				x.IsUntrustedReserveLocation = true
			case 3:
				x.IsUntrustedTeleportLocation = true
			case 4:
				x.IsMultiLocationFull = true
			case 5:
				x.IsMultiLocationNotInvertible = true
			case 6:
				x.IsBadOrigin = true
			case 7:
				x.IsInvalidLocation = true
			case 8:
				x.IsAssetNotFound = true
			case 9:
				x.IsFailedToTransactAsset = true
			case 10:
				x.IsNotWithdrawable = true
			case 11:
				x.IsLocationCannotHold = true
			case 12:
				x.IsExceedsMaxMessageSize = true
			case 13:
				x.IsDestinationUnsupported = true
			case 14:
				x.IsTransport = true

				c.Fuzz(&x.Transport)
			case 15:
				x.IsUnroutable = true
			case 16:
				x.IsUnknownClaim = true
			case 17:
				x.IsFailedToDecode = true
			case 18:
				x.IsMaxWeightInvalid = true
			case 19:
				x.IsNotHoldingFees = true
			case 20:
				x.IsTooExpensive = true
			case 21:
				x.IsTrap = true

				c.Fuzz(&x.TrapCode)
			case 22:
				x.IsUnhandledXcmVersion = true
			case 23:
				x.IsWeightLimitReached = true

				c.Fuzz(&x.Weight)
			case 24:
				x.IsBarrier = true
			case 25:
				x.IsWeightNotComputable = true
			}
		}),
	}
)

func TestXCMError_EncodeDecode(t *testing.T) {
	AssertRoundTripFuzz[XCMError](t, 1000, xcmErrorFuzzOpts...)
	AssertDecodeNilData[XCMError](t)
	AssertEncodeEmptyObj[XCMError](t, 0)
}
