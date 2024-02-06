package horizonhelper

import "github.com/stellar/go/clients/horizonclient"

const (
	TestChain   = "test"
	PublicChain = "public"
)

func HorizonClient(chain string) *horizonclient.Client {
	switch chain {
	case TestChain:
		return horizonclient.DefaultTestNetClient
	default:
		return horizonclient.DefaultPublicNetClient
	}
}
