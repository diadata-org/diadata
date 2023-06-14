package enums

type HealthState string

const (
	HealthStateOperational HealthState = "Operational"
	HealthStateMinor       HealthState = "Minor Outage"
	HealthStateMajor       HealthState = "Major Outage"
	HealthStateMaintenance HealthState = "Maintenance"
)

func CompareStates(itemState HealthState, groupState HealthState) HealthState {
	switch itemState {
	case HealthStateMajor:
		return HealthStateMajor
	case HealthStateMinor:
		if groupState == HealthStateOperational || groupState == HealthStateMaintenance {
			return HealthStateMinor
		}
	case HealthStateMaintenance:
		if groupState == HealthStateOperational {
			return HealthStateOperational
		}
	}
	return groupState
}
