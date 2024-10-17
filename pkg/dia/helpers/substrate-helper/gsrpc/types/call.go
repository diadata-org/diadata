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

package types

import (
	"io"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
)

// Call is the extrinsic function descriptor
type Call struct {
	CallIndex CallIndex
	Args      Args
}

func NewCall(m *Metadata, call string, args ...interface{}) (Call, error) {
	c, err := m.FindCallIndex(call)
	if err != nil {
		return Call{}, err
	}

	var a []byte
	for _, arg := range args {
		e, err := codec.Encode(arg)
		if err != nil {
			return Call{}, err
		}
		a = append(a, e...)
	}

	return Call{c, a}, nil
}

// Callindex is a 16 bit wrapper around the `[sectionIndex, methodIndex]` value that uniquely identifies a method
type CallIndex struct {
	SectionIndex uint8
	MethodIndex  uint8
}

func (m *CallIndex) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.SectionIndex)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.MethodIndex)
	if err != nil {
		return err
	}

	return nil
}

func (m CallIndex) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.SectionIndex)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.MethodIndex)
	if err != nil {
		return err
	}

	return nil
}

// Args are the encoded arguments for a Call
type Args []byte

// Encode implements encoding for Args, which just unwraps the bytes of Args
func (a Args) Encode(encoder scale.Encoder) error {
	return encoder.Write(a)
}

// Decode implements decoding for Args, which just reads all the remaining bytes into Args
func (a *Args) Decode(decoder scale.Decoder) error {
	for i := 0; true; i++ {
		b, err := decoder.ReadOneByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		*a = append((*a)[:i], b)
	}
	return nil
}

type Justification Bytes
