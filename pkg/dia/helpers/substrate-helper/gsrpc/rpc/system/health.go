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

package system

import (
	"github.com/diadata-org/diadata/pkg/dia/helpers/substrate-helper/gsrpc/types"
)

// Health retrieves the health status of the connected node
func (c *system) Health() (types.Health, error) {
	var h types.Health
	err := c.client.Call(&h, "system_health")
	return h, err
}
