package diaoracle

import "github.com/sirupsen/logrus"

var log = logrus.New()

type Oracalize interface {
	Update() error
}

type DataSource interface {
	getQuotation() error
}

type DataType interface {
	getQuotation() error
}
