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
	"fmt"
	"reflect"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	. "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
)

// PhaseEnum is an enum example. Since Go has no enums, it is implemented as a struct with flags for each
// potential value and a corresponding value if needed. This enum represents phases, the values are `ApplyExtrinsic`
// which can be a uint32 and `Finalization`. By implementing Encode and Decode methods on this struct that satisfy the
// scale.Encodeable and scale.Decodeable interfaces, we encode our enum struct to correspond to the scale codec
// (see https://substrate.dev/docs/en/overview/low-level-data-format for a description).
type PhaseEnum struct {
	IsApplyExtrinsic bool
	AsApplyExtrinsic uint32
	IsFinalization   bool
}

func (m *PhaseEnum) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsApplyExtrinsic = true
		err = decoder.Decode(&m.AsApplyExtrinsic)
	} else if b == 1 {
		m.IsFinalization = true
	}

	if err != nil {
		return err
	}

	return nil
}

func (m PhaseEnum) Encode(encoder scale.Encoder) error {
	var err1, err2 error
	if m.IsApplyExtrinsic {
		err1 = encoder.PushByte(0)
		err2 = encoder.Encode(m.AsApplyExtrinsic)
	} else if m.IsFinalization {
		err1 = encoder.PushByte(1)
	}

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil
}

func ExampleExampleEnum_applyExtrinsic() {
	applyExtrinsic := PhaseEnum{
		IsApplyExtrinsic: true,
		AsApplyExtrinsic: 1234,
	}

	enc, err := EncodeToHex(applyExtrinsic)
	if err != nil {
		panic(err)
	}

	var dec PhaseEnum
	err = DecodeFromHex(enc, &dec)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.DeepEqual(applyExtrinsic, dec))
}

func ExampleExampleEnum_finalization() {
	finalization := PhaseEnum{
		IsFinalization: true,
	}

	enc, err := EncodeToHex(finalization)
	if err != nil {
		panic(err)
	}

	var dec PhaseEnum
	err = DecodeFromHex(enc, &dec)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.DeepEqual(finalization, dec))
}
