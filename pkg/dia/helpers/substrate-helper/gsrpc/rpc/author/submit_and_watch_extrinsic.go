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

package author

import (
	"context"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/config"
	gethrpc "github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/gethrpc"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/codec"
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types/extrinsic"
)

// ExtrinsicStatusSubscription is a subscription established through one of the Client's subscribe methods.
type ExtrinsicStatusSubscription struct {
	sub      *gethrpc.ClientSubscription
	channel  chan types.ExtrinsicStatus
	quitOnce sync.Once // ensures quit is closed once
}

// Chan returns the subscription channel.
//
// The channel is closed when Unsubscribe is called on the subscription.
func (s *ExtrinsicStatusSubscription) Chan() <-chan types.ExtrinsicStatus {
	return s.channel
}

// Err returns the subscription error channel. The intended use of Err is to schedule
// resubscription when the client connection is closed unexpectedly.
//
// The error channel receives a value when the subscription has ended due
// to an error. The received error is nil if Close has been called
// on the underlying client and no other error has occurred.
//
// The error channel is closed when Unsubscribe is called on the subscription.
func (s *ExtrinsicStatusSubscription) Err() <-chan error {
	return s.sub.Err()
}

// Unsubscribe unsubscribes the notification and closes the error channel.
// It can safely be called more than once.
func (s *ExtrinsicStatusSubscription) Unsubscribe() {
	s.sub.Unsubscribe()
	s.quitOnce.Do(func() {
		close(s.channel)
	})
}

// SubmitAndWatchExtrinsic will submit and subscribe to watch an extrinsic until unsubscribed, returning a subscription
// that will receive server notifications containing the extrinsic status updates.
func (a *author) SubmitAndWatchExtrinsic(xt extrinsic.Extrinsic) (*ExtrinsicStatusSubscription, error) { //nolint:lll
	hexEncodedExtrinsic, err := codec.EncodeToHex(xt)
	if err != nil {
		return nil, err
	}

	return a.submitAndWatchExtrinsic(hexEncodedExtrinsic)
}

func (a *author) submitAndWatchExtrinsic(hexEncodedExtrinsic string) (*ExtrinsicStatusSubscription, error) { //nolint:lll
	ctx, cancel := context.WithTimeout(context.Background(), config.Default().SubscribeTimeout)
	defer cancel()

	c := make(chan types.ExtrinsicStatus)

	sub, err := a.client.Subscribe(ctx, "author", "submitAndWatchExtrinsic", "unwatchExtrinsic", "extrinsicUpdate",
		c, hexEncodedExtrinsic)
	if err != nil {
		return nil, err
	}

	return &ExtrinsicStatusSubscription{sub: sub, channel: c}, nil
}
