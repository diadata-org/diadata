package enums

type HealthSate string

const (
	HealthStateOperational HealthSate = "Operational"
	HealthStateMinor       HealthSate = "Minor Outage"
	HealthStateMajor       HealthSate = "Major Outage"
	HealthStateMaintenance HealthSate = "Maintenance"
)

func CompareStates(itemState HealthSate, groupState HealthSate) HealthSate {
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
