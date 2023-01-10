package enums

import (
	"fmt"
	"strings"
)

type SectionType string

const (
	ScraperSection  = "scraper"
	ServiceSection  = "service"
	DatabaseSection = "database"
	UnknownSection  = "Unknown"
)

func GetSectionTypeFromString(section string) (SectionType, error) {
	switch strings.TrimSpace(strings.ToLower(section)) {
	case "scraper":
		return ScraperSection, nil
	case "service":
		return ServiceSection, nil
	case "database":
		return DatabaseSection, nil
	}
	return UnknownSection, fmt.Errorf("could not parse section type for section %s", section)
}
