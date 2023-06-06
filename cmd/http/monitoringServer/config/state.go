package config

import "github.com/diadata-org/diadata/http/monitoringServer/enums"

type State struct {
	Name       string
	State      enums.HealthState
	Subsection []State
}

func GetMajorHealthState(name string) (state State) {
	state.Name = name
	state.State = enums.HealthStateMajor
	state.Subsection = []State{}
	return
}

func GetOperationalHealthState(name string) (state State) {
	state.Name = name
	state.State = enums.HealthStateOperational
	state.Subsection = []State{}
	return
}
