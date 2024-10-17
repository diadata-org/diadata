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
	"errors"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/scale"
)

type Tranche struct {
	FirstVal  U64
	SecondVal [16]U8
}

func (t *Tranche) Decode(decoder scale.Decoder) error {
	if err := decoder.Decode(&t.FirstVal); err != nil {
		return err
	}

	return decoder.Decode(&t.SecondVal)
}

func (t Tranche) Encode(encoder scale.Encoder) error {
	if err := encoder.Encode(t.FirstVal); err != nil {
		return err
	}

	return encoder.Encode(t.SecondVal)
}

type PermissionedCurrency struct {
	//  At the moment of writing this, this enum is empty in altair.
}

func (p *PermissionedCurrency) Decode(_ scale.Decoder) error {
	return nil
}

func (p *PermissionedCurrency) Encode(_ scale.Encoder) error {
	return nil
}

type StakingCurrency struct {
	IsBlockRewards bool
}

func (s *StakingCurrency) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		s.IsBlockRewards = true

		return nil
	default:
		return errors.New("unsupported staking currency")
	}
}

func (s StakingCurrency) Encode(encoder scale.Encoder) error {
	switch {
	case s.IsBlockRewards:
		return encoder.PushByte(0)
	default:
		return errors.New("unsupported staking currency")
	}
}

type CurrencyID struct {
	IsNative bool

	IsTranche bool
	Tranche   Tranche

	IsKSM bool

	IsAUSD bool

	IsForeignAsset bool
	AsForeignAsset U32

	IsStaking bool
	AsStaking StakingCurrency
}

func (c *CurrencyID) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		c.IsNative = true

		return nil
	case 1:
		c.IsTranche = true

		return decoder.Decode(&c.Tranche)
	case 2:
		c.IsKSM = true

		return nil
	case 3:
		c.IsAUSD = true

		return nil
	case 4:
		c.IsForeignAsset = true

		return decoder.Decode(&c.AsForeignAsset)
	case 5:
		c.IsStaking = true

		return decoder.Decode(&c.AsStaking)
	default:
		return errors.New("unsupported currency ID")
	}
}

func (c CurrencyID) Encode(encoder scale.Encoder) error {
	switch {
	case c.IsNative:
		return encoder.PushByte(0)
	case c.IsTranche:
		if err := encoder.PushByte(1); err != nil {
			return err
		}

		return encoder.Encode(c.Tranche)
	case c.IsKSM:
		return encoder.PushByte(2)
	case c.IsAUSD:
		return encoder.PushByte(3)
	case c.IsForeignAsset:
		if err := encoder.PushByte(4); err != nil {
			return err
		}

		return encoder.Encode(c.AsForeignAsset)
	case c.IsStaking:
		if err := encoder.PushByte(5); err != nil {
			return err
		}

		return encoder.Encode(c.AsStaking)
	default:
		return errors.New("unsupported currency ID")
	}
}

type Price struct {
	CurrencyID CurrencyID
	Amount     U128
}

func (p *Price) Decode(decoder scale.Decoder) error {
	if err := decoder.Decode(&p.CurrencyID); err != nil {
		return err
	}

	return decoder.Decode(&p.Amount)
}

func (p Price) Encode(encoder scale.Encoder) error {
	if err := encoder.Encode(p.CurrencyID); err != nil {
		return err
	}

	return encoder.Encode(p.Amount)
}

type Sale struct {
	Seller AccountID
	Price  Price
}

func (s *Sale) Decode(decoder scale.Decoder) error {
	if err := decoder.Decode(&s.Seller); err != nil {
		return err
	}

	return decoder.Decode(&s.Price)
}

func (s Sale) Encode(encoder scale.Encoder) error {
	if err := encoder.Encode(s.Seller); err != nil {
		return err
	}

	return encoder.Encode(s.Price)
}
