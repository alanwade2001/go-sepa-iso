package schema

import (
	"encoding/xml"

	"github.com/alanwade2001/go-sepa-iso/pacs_008_001_02"
	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

// Document did not generate properly...
type Pain001Document struct {
	XMLName          xml.Name                                            `xml:"Document"`
	CstmrCdtTrfInitn pain_001_001_03.CustomerCreditTransferInitiationV03 `xml:"CstmrCdtTrfInitn"`
}

type Pacs008Document struct {
	XMLName           xml.Name                                        `xml:"Document"`
	FIToFICstmrCdtTrf pacs_008_001_02.FIToFICustomerCreditTransferV02 `xml:"FIToFICstmrCdtTrf"`
}

func NewPain001XsdHandler() (*xsdvalidate.XsdHandler, error) {
	return xsdvalidate.NewXsdHandlerUrl("../xsd/pain.001.001.03.xsd", xsdvalidate.ParsErrVerbose)

}

func NewPacs008XsdHandler() (*xsdvalidate.XsdHandler, error) {
	return xsdvalidate.NewXsdHandlerUrl("../xsd/pacs.008.001.02.xsd", xsdvalidate.ParsErrVerbose)

}
