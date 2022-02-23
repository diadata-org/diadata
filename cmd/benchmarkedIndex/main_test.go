package main

import (
	"testing"
)

const (
	filenameDia            = "/data/DIAData_DIA_Index_History.csv"
	filenameBtc            = "/data/DIAData_BTC_Index_History.csv"
	filenameWithCurrentDir = "./DIAData_BTC_Index_History.csv"
)

func TestGetFilename(t *testing.T) {
	filename := getFilename(filenameDia)
	if filename != "DIAData_DIA_Index_History.csv" {
		t.Errorf("expected DIAData_DIA_Index_History.csv got  %v", filename)
	}
	filename = getFilename(filenameBtc)
	if filename != "DIAData_BTC_Index_History.csv" {
		t.Errorf("expected DIAData_BTC_Index_History.csv got  %v", filename)
	}
}

func TestGetSymbol(t *testing.T) {

	symbol := getSymbol(filenameDia)
	if symbol != "DIA" {
		t.Errorf("expected symbol DIA got %v", symbol)
	}
	symbol = getSymbol(filenameBtc)
	if symbol != "BTC" {
		t.Errorf("expected symbol DIA got %v", symbol)
	}
}
