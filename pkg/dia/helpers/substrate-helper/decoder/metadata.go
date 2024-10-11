package decoder

import (
	"strings"

	"github.com/itering/scale.go"
	"github.com/itering/scale.go/types"
)

type RuntimeRaw struct {
	Spec int
	Raw  string
}

type Instant types.MetadataStruct

var (
	latestSpec      = -1
	RuntimeMetadata = make(map[int]*Instant)
	Decoder         *scalecodec.MetadataDecoder
)

func Latest(runtime *RuntimeRaw) *Instant {
	if latestSpec != -1 {
		d := RuntimeMetadata[latestSpec]
		return d
	}
	if runtime == nil {
		return nil
	}
	m := scalecodec.MetadataDecoder{}
	m.Init(HexToBytes(runtime.Raw))
	_ = m.Process()

	Decoder = &m
	latestSpec = runtime.Spec
	instant := Instant(m.Metadata)
	RuntimeMetadata[latestSpec] = &instant
	return &instant
}

func Process(runtime *RuntimeRaw) *Instant {
	if runtime == nil {
		return nil
	}
	if d, ok := RuntimeMetadata[runtime.Spec]; ok {
		return d
	}

	m := scalecodec.MetadataDecoder{}
	m.Init(HexToBytes(runtime.Raw))
	_ = m.Process()

	instant := Instant(m.Metadata)
	RuntimeMetadata[runtime.Spec] = &instant

	return &instant
}

func RegNewMetadataType(spec int, coded string) *Instant {
	m := scalecodec.MetadataDecoder{}
	m.Init(HexToBytes(coded))
	_ = m.Process()

	instant := Instant(m.Metadata)
	RuntimeMetadata[spec] = &instant

	if spec > latestSpec {
		latestSpec = spec
	}
	return &instant
}

func (m *Instant) FindCallCallName(moduleName, callName string) *types.MetadataCalls {
	for index, v := range m.CallIndex {
		if strings.EqualFold(v.Call.Name, callName) && strings.EqualFold(v.Module.Name, moduleName) {
			call := v.Call
			call.Lookup = index
			return &call
		}
	}
	return nil
}
