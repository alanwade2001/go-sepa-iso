package schema

import (
	_ "embed"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

//go:embed pain.001.001.03.xsd
var p1Bytes []byte

func NewPain001XsdHandler() (*xsdvalidate.XsdHandler, error) {

	return xsdvalidate.NewXsdHandlerMem(p1Bytes, xsdvalidate.ParsErrVerbose)

}

//go:embed pacs.008.001.02.xsd
var p8Bytes []byte

func NewPacs008XsdHandler() (*xsdvalidate.XsdHandler, error) {

	return xsdvalidate.NewXsdHandlerMem(p8Bytes, xsdvalidate.ParsErrVerbose)

}
