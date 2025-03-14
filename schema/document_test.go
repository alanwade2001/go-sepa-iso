package schema_test

import (
	"encoding/xml"
	"testing"

	"github.com/alanwade2001/go-sepa-iso/schema"
)

var GOOD_XML string = `
<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.001.001.03">
    <CstmrCdtTrfInitn>
        <GrpHdr>
            <MsgId>MSG_ID1</MsgId>
            <CreDtTm>2021-09-28T13:16:37.219430288</CreDtTm>
            <NbOfTxs>1</NbOfTxs>
            <CtrlSum>11</CtrlSum>
            <InitgPty>
                <Id>
                    <PrvtId>
                        <Othr>
                            <Id>IP-123456</Id>
                        </Othr>
                    </PrvtId>
                </Id>
            </InitgPty>
        </GrpHdr>
        <PmtInf>
            <PmtInfId>PIID-1</PmtInfId>
            <PmtMtd>TRF</PmtMtd>
            <NbOfTxs>1</NbOfTxs>
            <CtrlSum>11</CtrlSum>
            <PmtTpInf/>
            <ReqdExctnDt>2025-01-10</ReqdExctnDt>
            <Dbtr>
                <Nm>Mr. Debtor</Nm>
            </Dbtr>
            <DbtrAcct>
                <Id>
                    <IBAN>IE30BOFI90909012345678</IBAN>
                </Id>
            </DbtrAcct>
            <DbtrAgt>
                <FinInstnId>
                    <BIC>BOFIIE2D</BIC>
                </FinInstnId>
            </DbtrAgt>
            <ChrgBr>CRED</ChrgBr>
            <CdtTrfTxInf>
                <PmtId>
                    <EndToEndId>E2EID-2</EndToEndId>
                </PmtId>
                <Amt>
                    <InstdAmt Ccy="EUR">11</InstdAmt>
                </Amt>
                <CdtrAgt>
                    <FinInstnId>
                        <BIC>BOFIIE2D</BIC>
                    </FinInstnId>
                </CdtrAgt>
                <Cdtr>
                    <Nm>Mr. Cdtr</Nm>
                </Cdtr>
                <CdtrAcct>
                    <Id>
                        <IBAN>IE16BOFI90909187654321</IBAN>
                    </Id>
                </CdtrAcct>
            </CdtTrfTxInf>
        </PmtInf>
    </CstmrCdtTrfInitn>
</Document>
`

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestDocumentGood(t *testing.T) {
	doc := &schema.P1Document{}

	err := xml.Unmarshal([]byte(GOOD_XML), doc)
	if err != nil {
		t.Fatalf(`Failed to parse GOOD_XML %v error`, err)
	}
}

func TestDocumentNotXml(t *testing.T) {
	doc := &schema.P1Document{}

	err := xml.Unmarshal([]byte("not xml"), doc)
	if err == nil {
		t.Fatalf(`Parsed NOT_XML with no error`)
	}
}

func TestDocumentHasMsgId(t *testing.T) {
	doc := &schema.P1Document{}

	err := xml.Unmarshal([]byte(GOOD_XML), doc)
	if err != nil {
		t.Fatalf(`Failed to parse GOOD_XML %v error`, err)
	}

	expected := "MSG_ID1"
	actual := doc.CstmrCdtTrfInitn.GrpHdr.MsgId

	if expected != actual {
		t.Fatalf(`Failed to get MsgId: expected[%q] != actual[%q]`, expected, actual)
	}
}
