package enums

type HealthSate string

const (
	HealthStateOperational HealthSate = "Operational"
	HealthStateMinor       HealthSate = "Minor Outage"
	HealthStateMajor       HealthSate = "Major Outage"
	HealthStateMaintenance HealthSate = "Maintenance"
)
