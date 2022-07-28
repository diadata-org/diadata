package platform

import (
	"github.com/diadata-org/diadata/http/monitoringServer/config"
)

func GetAllStates() (states []config.State) {
	state := config.GetOperationalHealthState("Services")
	states = append(states, state)
	return
}
