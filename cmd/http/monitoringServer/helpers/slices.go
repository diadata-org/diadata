package helpers

import (
	"github.com/diadata-org/diadata/http/monitoringServer/config"
	"github.com/diadata-org/diadata/http/monitoringServer/enums"
)

func MergeStateSlicesAsSubsection(name string, states []config.State) config.State {
	state := config.GetOperationalHealthState(name)

	for _, oneSlice := range states {
		state.Subsection = append(state.Subsection, oneSlice)
		state.State = CompareHealthStates(state.State, oneSlice.State)
	}

	return state
}

func CompareHealthStates(currentSate enums.HealthSate, newState enums.HealthSate) enums.HealthSate {
	switch newState {
	case enums.HealthStateMajor:
		return enums.HealthStateMajor
	case enums.HealthStateMinor:
		if currentSate == enums.HealthStateOperational || currentSate == enums.HealthStateMaintenance {
			return enums.HealthStateMinor
		}
	case enums.HealthStateMaintenance:
		if currentSate == enums.HealthStateOperational {
			return enums.HealthStateMaintenance
		}
	}
	return enums.HealthStateOperational
}