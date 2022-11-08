package source

import (
	"fmt"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia"
)

func TestSerum(t *testing.T) {
	service := NewSerumAssetSource(dia.Exchange{})
	names, err := service.getTokenNames()
	fmt.Println(err)

	i := 0
	for k, v := range names {
		i++
		fmt.Println(i)
		fmt.Println(k)
		fmt.Println("name:", v.name)
		fmt.Println("symbol:", v.symbol)
		fmt.Println("decimals:", v.decimals)
		fmt.Println("mint:", v.mint)
		fmt.Println("")
	}
}
