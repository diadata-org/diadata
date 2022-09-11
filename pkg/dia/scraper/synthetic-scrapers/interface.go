package synthscrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
)

type SynthScraperInterface interface {
	// Synth asset data should be streamed through dia.SynthAssetSupply.
	GetSynthSupplyChannel() chan dia.SynthAssetSupply
	// Should fetch synth data and send them to the channel.
	FetchSynthSupply() error
}
