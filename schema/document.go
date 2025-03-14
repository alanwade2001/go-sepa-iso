package schema

import (
	"encoding/xml"

	"github.com/alanwade2001/go-sepa-iso/gen"
)

// Document did not generate properly...
type P1Document struct {
	XMLName          xml.Name                                `xml:"Document"`
	CstmrCdtTrfInitn gen.CustomerCreditTransferInitiationV03 `xml:"CstmrCdtTrfInitn"`
}
