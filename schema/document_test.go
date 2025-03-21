package schema_test

import (
	"encoding/xml"
	"log/slog"
	"strings"
	"testing"

	"github.com/alanwade2001/go-sepa-iso/pacs_008_001_02"
	"github.com/alanwade2001/go-sepa-iso/pain_001_001_03"
	"github.com/alanwade2001/go-sepa-iso/schema"
	"github.com/stretchr/testify/assert"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

var P1_GOOD_XML string = `<?xml version="1.0" encoding="UTF-8"?>
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
</Document>`

var P1_BAD_XML_MISSING_MSGID string = `<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.001.001.03">
    <CstmrCdtTrfInitn>
        <GrpHdr>            
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
</Document>`

var P1_BAD_XML_NBOFTXS_NOT_NUMBER string = `<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.001.001.03">
    <CstmrCdtTrfInitn>
        <GrpHdr>            
            <CreDtTm>2021-09-28T13:16:37.219430288</CreDtTm>
            <MsgId>m1</MsgId>
            <NbOfTxs>a</NbOfTxs>
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
</Document>`

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestDocumentGood(t *testing.T) {
	doc := &pain_001_001_03.Document{}

	err := xml.Unmarshal([]byte(P1_GOOD_XML), doc)
	if err != nil {
		t.Fatalf(`Failed to parse GOOD_XML %v error`, err)
	}
}

func TestDocumentNotXml(t *testing.T) {
	doc := &pacs_008_001_02.Document{}

	err := xml.Unmarshal([]byte("not xml"), doc)
	if err == nil {
		t.Fatalf(`Parsed NOT_XML with no error`)
	}
}

func TestDocumentHasMsgId(t *testing.T) {
	doc := &pain_001_001_03.Document{}

	err := xml.Unmarshal([]byte(P1_GOOD_XML), doc)
	if err != nil {
		t.Fatalf(`Failed to parse GOOD_XML %v error`, err)
	}

	expected := "MSG_ID1"
	actual := doc.CstmrCdtTrfInitn.GrpHdr.MsgId

	if expected != actual {
		t.Fatalf(`Failed to get MsgId: expected[%q] != actual[%q]`, expected, actual)
	}
}

func TestPain001Schema(t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()

	pain001Handler, err := schema.NewPain001XsdHandler()
	assert.NoError(t, err)

	defer pain001Handler.Free()

	testCases := []struct {
		desc    string
		content string
		pass    bool
	}{
		{
			desc:    "Good Xml",
			content: P1_GOOD_XML,
			pass:    true,
		},
		{
			desc:    "Bad Xml Missing Msg Id",
			content: P1_BAD_XML_MISSING_MSGID,
			pass:    false,
		},
		{
			desc:    "Bad Xml NbOfTxs not number",
			content: P1_BAD_XML_NBOFTXS_NOT_NUMBER,
			pass:    false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := pain001Handler.ValidateMem([]byte(tC.content), xsdvalidate.ParsErrVerbose)

			if tC.pass {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

		})
	}
}

func TestEmpty(t *testing.T) {

	c := &pacs_008_001_02.AccountSchemeName1Choice{
		Prtry: "prtry",
	}

	data, err := xml.Marshal(c)

	assert.NoError(t, err)

	str := string(data)

	assert.False(t, strings.Contains(str, "Cd"))

}

func Test(t *testing.T) {

	data := `<Document>
  <FIToFICstmrCdtTrf>
    <GrpHdr>
      <MsgId>SCTORD156820211213000000012649</MsgId>
      <CreDtTm>2025-322-36T36:2240:69</CreDtTm>
      <BtchBookg>false</BtchBookg>
      <NbOfTxs>1</NbOfTxs>
      <CtrlSum>11</CtrlSum>
      <TtlIntrBkSttlmAmt Ccy="EUR">11</TtlIntrBkSttlmAmt>
      <IntrBkSttlmDt></IntrBkSttlmDt>
      <SttlmInf>
        <SttlmMtd>CLRG</SttlmMtd>
        <ClrSys>
          <Prtry>ST2</Prtry>
        </ClrSys>
      </SttlmInf>
      <InstgAgt>
        <FinInstnId>
          <BIC>BTRLRO22</BIC>
          <Nm></Nm>
        </FinInstnId>
      </InstgAgt>
      <InstdAgt>
        <FinInstnId>
          <BIC>BTRLRO22</BIC>
          <Nm></Nm>
        </FinInstnId>
      </InstdAgt>
    </GrpHdr>
    <CdtTrfTxInf>
      <PmtId>
        <InstrId></InstrId>
        <EndToEndId>E2EID-2</EndToEndId>
        <TxId></TxId>
        <ClrSysRef></ClrSysRef>
      </PmtId>
      <PmtTpInf>
        <InstrPrty></InstrPrty>
        <ClrChanl></ClrChanl>
      </PmtTpInf>
      <IntrBkSttlmAmt Ccy="EUR">11</IntrBkSttlmAmt>
      <IntrBkSttlmDt>2025-01-10</IntrBkSttlmDt>
      <SttlmPrty></SttlmPrty>
      <AccptncDtTm></AccptncDtTm>
      <PoolgAdjstmntDt></PoolgAdjstmntDt>
      <XchgRate>0</XchgRate>
      <ChrgBr>CRED</ChrgBr>
      <Dbtr>
        <Nm>Mr. Debtor</Nm>
        <CtryOfRes></CtryOfRes>
      </Dbtr>
      <DbtrAcct>
        <Id>
          <IBAN>IE30BOFI90909012345678</IBAN>
        </Id>
        <Ccy></Ccy>
        <Nm></Nm>
      </DbtrAcct>
      <DbtrAgt>
        <FinInstnId>
          <BIC>BOFIIE2D</BIC>
          <Nm></Nm>
        </FinInstnId>
      </DbtrAgt>
      <CdtrAgt>
        <FinInstnId>
          <BIC>BOFIIE2D</BIC>
          <Nm></Nm>
        </FinInstnId>
      </CdtrAgt>
      <Cdtr>
        <Nm>Mr. Cdtr</Nm>
        <CtryOfRes></CtryOfRes>
      </Cdtr>
      <CdtrAcct>
        <Id>
          <IBAN>IE16BOFI90909187654321</IBAN>
        </Id>
        <Ccy></Ccy>
        <Nm></Nm>
      </CdtrAcct>
    </CdtTrfTxInf>
  </FIToFICstmrCdtTrf>
</Document>`

	p8 := &pacs_008_001_02.Document{}
	xml.Unmarshal([]byte(data), p8)

	output, err := xml.Marshal(p8)

	assert.NoError(t, err)

	assert.False(t, strings.Contains(string(output), "<Ccy></Ccy>"))

	slog.Info("test", "output", output)

}
