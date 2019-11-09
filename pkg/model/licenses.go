package models

import (
	"http"
	"io"
	"encoding/json"
	"time"
)

type License struct {
	Name				string
	Copyleft		bool
	Symbol			string
	Commercial	bool
	Derivation	bool
	Inheriting	bool
	FullTerms		string
}

type Creators {
	Licenses		[]License
	Names				[]string
	LastUpdate	time.Date
}
