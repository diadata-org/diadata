package stateFilter

import (
	"github.com/diadata-org/diadata/http/monitoringServer/config"
	"github.com/diadata-org/diadata/http/monitoringServer/enums"
)

func CheckHealthState(states []config.State, operationalThreshold, minorThreshold float64) enums.HealthState {
	totalStates := len(states)
	operationalStates := 0

	for _, state := range states {
		if state.State == enums.HealthStateOperational {
			operationalStates++
		}
	}

	operationalPercentage := float64(operationalStates) / float64(totalStates)

	switch {
	case operationalPercentage >= operationalThreshold:
		return enums.HealthStateOperational
	case operationalPercentage >= minorThreshold:
		return enums.HealthStateMinor
	default:
		return enums.HealthStateMajor
	}
}
