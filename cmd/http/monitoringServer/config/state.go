package config

import (
	"github.com/diadata-org/diadata/http/monitoringServer/enums"
	log "github.com/sirupsen/logrus"
)

type State struct {
	Name        string
	State       enums.HealthSate
	Subsection  []State
	SectionType enums.SectionType
}

func getHealthState(name string, section string) (state State) {
	var err error
	state.Name = name
	state.Subsection = []State{}
	state.SectionType, err = enums.GetSectionTypeFromString(section)
	if err != nil {
		log.Error("error getting section ", err)
	}
	return state
}

func getHealthStateWithSection(name string, section enums.SectionType) (state State) {
	state.Name = name
	state.Subsection = []State{}
	state.SectionType = section
	return state
}

func GetMajorHealthState(name string, section string) (state State) {
	state = getHealthState(name, section)
	state.State = enums.HealthStateMajor
	return
}

func GetOperationalHealthState(name string, section string) (state State) {
	state = getHealthState(name, section)
	state.State = enums.HealthStateOperational
	return
}

func GetOperationalHealthStateWithSection(name string, section enums.SectionType) (state State) {
	state = getHealthStateWithSection(name, section)
	state.State = enums.HealthStateOperational
	return
}
