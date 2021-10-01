package aave

import (
	"context"
	"sync"

	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/aave/bind/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type erc20Metadata struct {
	addr     common.Address
	name     string
	symbol   string
	decimals int
}

type erc20MetadataCache struct {
	backend bind.ContractBackend

	mu    sync.RWMutex
	cache map[common.Address]*erc20Metadata
}

func NewERC20MetadataCache(backend bind.ContractBackend) ERC20MetadataProvider {
	return &erc20MetadataCache{backend: backend, cache: make(map[common.Address]*erc20Metadata)}
}

func (c *erc20MetadataCache) ERC20(ctx context.Context, addr common.Address) (ERC20Metadata, error) {
	c.mu.RLock()
	md, ok := c.cache[addr]
	c.mu.RUnlock()

	if ok {
		return md, nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	contract, err := erc20.NewIERC20Metadata(addr, c.backend)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create ERC20 contract binding for address %s: %s", addr.String(), err.Error())
	}

	opts := &bind.CallOpts{Context: ctx}

	erc20md := &erc20Metadata{addr: addr}

	if erc20md.name, err = contract.Name(opts); err != nil {
		return nil, errors.Wrapf(err, "unable to read metadata(Name) on ERC20(%s): %s", addr.String(), err.Error())
	}

	if erc20md.symbol, err = contract.Symbol(opts); err != nil {
		return nil, errors.Wrapf(err, "unable to read metadata(Symbol) on ERC20(%s): %s", addr.String(), err.Error())
	}

	decimals, err := contract.Decimals(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read metadata(Decimals) on ERC20(%s): %s", addr.String(), err.Error())
	}

	erc20md.decimals = int(decimals)

	c.cache[addr] = erc20md

	return erc20md, nil
}

func (md *erc20Metadata) Name() string {
	return md.name
}

func (md *erc20Metadata) Symbol() string {
	return md.symbol
}

func (md *erc20Metadata) Decimals() int {
	return md.decimals
}
