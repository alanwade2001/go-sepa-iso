package pacs_008_001_02

import (
	"time"
)

// Document ...
type Document *Document

// AccountIdentification4Choice ...
type AccountIdentification4Choice struct {
	IBAN string                         `xml:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"Othr,omitempty"`
}

// AccountSchemeName1Choice ...
type AccountSchemeName1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// ActiveCurrencyAndAmountSimpleType ...
type ActiveCurrencyAndAmountSimpleType float64

// ActiveCurrencyAndAmount ...
type ActiveCurrencyAndAmount struct {
	CcyAttr string  `xml:"Ccy,attr,omitempty"`
	Value   float64 `xml:",chardata,omitempty"`
}

// ActiveCurrencyCode ...
type ActiveCurrencyCode string

// ActiveOrHistoricCurrencyAndAmountSimpleType ...
type ActiveOrHistoricCurrencyAndAmountSimpleType float64

// ActiveOrHistoricCurrencyAndAmount ...
type ActiveOrHistoricCurrencyAndAmount struct {
	CcyAttr string  `xml:"Ccy,attr,omitempty"`
	Value   float64 `xml:",chardata,omitempty"`
}

// ActiveOrHistoricCurrencyCode ...
type ActiveOrHistoricCurrencyCode string

// AddressType2Code ...
type AddressType2Code string

// AnyBICIdentifier ...
type AnyBICIdentifier string

// BICIdentifier ...
type BICIdentifier string

// BaseOneRate ...
type BaseOneRate float64

// BatchBookingIndicator ...
type BatchBookingIndicator bool

// BranchAndFinancialInstitutionIdentification4 ...
type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId *FinancialInstitutionIdentification7 `xml:"FinInstnId,omitempty"`
	BrnchId    *BranchData2                         `xml:"BrnchId,omitempty"`
}

// BranchData2 ...
type BranchData2 struct {
	Id      string          `xml:"Id,omitempty"`
	Nm      string          `xml:"Nm,omitempty"`
	PstlAdr *PostalAddress6 `xml:"PstlAdr,omitempty"`
}

// CashAccount16 ...
type CashAccount16 struct {
	Id  *AccountIdentification4Choice `xml:"Id,omitempty"`
	Tp  *CashAccountType2             `xml:"Tp,omitempty"`
	Ccy string                        `xml:"Ccy,omitempty"`
	Nm  string                        `xml:"Nm,omitempty"`
}

// CashAccountType2 ...
type CashAccountType2 struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// CashAccountType4Code ...
type CashAccountType4Code string

// CategoryPurpose1Choice ...
type CategoryPurpose1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// ChargeBearerType1Code ...
type ChargeBearerType1Code string

// ChargesInformation5 ...
type ChargesInformation5 struct {
	Amt *ActiveOrHistoricCurrencyAndAmount            `xml:"Amt,omitempty"`
	Pty *BranchAndFinancialInstitutionIdentification4 `xml:"Pty,omitempty"`
}

// ClearingChannel2Code ...
type ClearingChannel2Code string

// ClearingSystemIdentification2Choice ...
type ClearingSystemIdentification2Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// ClearingSystemIdentification3Choice ...
type ClearingSystemIdentification3Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// ClearingSystemMemberIdentification2 ...
type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty"`
	MmbId    string                               `xml:"MmbId,omitempty"`
}

// ContactDetails2 ...
type ContactDetails2 struct {
	NmPrfx   string `xml:"NmPrfx,omitempty"`
	Nm       string `xml:"Nm,omitempty"`
	PhneNb   string `xml:"PhneNb,omitempty"`
	MobNb    string `xml:"MobNb,omitempty"`
	FaxNb    string `xml:"FaxNb,omitempty"`
	EmailAdr string `xml:"EmailAdr,omitempty"`
	Othr     string `xml:"Othr,omitempty"`
}

// CountryCode ...
type CountryCode string

// CreditDebitCode ...
type CreditDebitCode string

// CreditTransferTransactionInformation11 ...
type CreditTransferTransactionInformation11 struct {
	PmtId            *PaymentIdentification3                       `xml:"PmtId,omitempty"`
	PmtTpInf         *PaymentTypeInformation21                     `xml:"PmtTpInf,omitempty"`
	IntrBkSttlmAmt   *ActiveCurrencyAndAmount                      `xml:"IntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt    string                                        `xml:"IntrBkSttlmDt,omitempty"`
	SttlmPrty        string                                        `xml:"SttlmPrty,omitempty"`
	SttlmTmIndctn    *SettlementDateTimeIndication1                `xml:"SttlmTmIndctn,omitempty"`
	SttlmTmReq       *SettlementTimeRequest2                       `xml:"SttlmTmReq,omitempty"`
	AccptncDtTm      string                                        `xml:"AccptncDtTm,omitempty"`
	PoolgAdjstmntDt  string                                        `xml:"PoolgAdjstmntDt,omitempty"`
	InstdAmt         *ActiveOrHistoricCurrencyAndAmount            `xml:"InstdAmt,omitempty"`
	XchgRate         float64                                       `xml:"XchgRate,omitempty"`
	ChrgBr           string                                        `xml:"ChrgBr,omitempty"`
	ChrgsInf         []*ChargesInformation5                        `xml:"ChrgsInf,omitempty"`
	PrvsInstgAgt     *BranchAndFinancialInstitutionIdentification4 `xml:"PrvsInstgAgt,omitempty"`
	PrvsInstgAgtAcct *CashAccount16                                `xml:"PrvsInstgAgtAcct,omitempty"`
	InstgAgt         *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`
	InstdAgt         *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
	IntrmyAgt1       *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct   *CashAccount16                                `xml:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2       *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct   *CashAccount16                                `xml:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3       *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct   *CashAccount16                                `xml:"IntrmyAgt3Acct,omitempty"`
	UltmtDbtr        *PartyIdentification32                        `xml:"UltmtDbtr,omitempty"`
	InitgPty         *PartyIdentification32                        `xml:"InitgPty,omitempty"`
	Dbtr             *PartyIdentification32                        `xml:"Dbtr,omitempty"`
	DbtrAcct         *CashAccount16                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt          *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`
	DbtrAgtAcct      *CashAccount16                                `xml:"DbtrAgtAcct,omitempty"`
	CdtrAgt          *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct      *CashAccount16                                `xml:"CdtrAgtAcct,omitempty"`
	Cdtr             *PartyIdentification32                        `xml:"Cdtr,omitempty"`
	CdtrAcct         *CashAccount16                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr        *PartyIdentification32                        `xml:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt  []*InstructionForCreditorAgent1               `xml:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt   []*InstructionForNextAgent1                   `xml:"InstrForNxtAgt,omitempty"`
	Purp             *Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg       []*RegulatoryReporting3                       `xml:"RgltryRptg,omitempty"`
	RltdRmtInf       []*RemittanceLocation2                        `xml:"RltdRmtInf,omitempty"`
	RmtInf           *RemittanceInformation5                       `xml:"RmtInf,omitempty"`
}

// CreditorReferenceInformation2 ...
type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"Tp,omitempty"`
	Ref string                  `xml:"Ref,omitempty"`
}

// CreditorReferenceType1Choice ...
type CreditorReferenceType1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// CreditorReferenceType2 ...
type CreditorReferenceType2 struct {
	CdOrPrtry *CreditorReferenceType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      string                        `xml:"Issr,omitempty"`
}

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	BirthDt     string `xml:"BirthDt,omitempty"`
	PrvcOfBirth string `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth string `xml:"CityOfBirth,omitempty"`
	CtryOfBirth string `xml:"CtryOfBirth,omitempty"`
}

// DecimalNumber ...
type DecimalNumber float64

// DocumentAdjustment1 ...
type DocumentAdjustment1 struct {
	Amt       *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	CdtDbtInd string                             `xml:"CdtDbtInd,omitempty"`
	Rsn       string                             `xml:"Rsn,omitempty"`
	AddtlInf  string                             `xml:"AddtlInf,omitempty"`
}

// DocumentType3Code ...
type DocumentType3Code string

// DocumentType5Code ...
type DocumentType5Code string

// ExternalAccountIdentification1Code ...
type ExternalAccountIdentification1Code string

// ExternalCashClearingSystem1Code ...
type ExternalCashClearingSystem1Code string

// ExternalCategoryPurpose1Code ...
type ExternalCategoryPurpose1Code string

// ExternalClearingSystemIdentification1Code ...
type ExternalClearingSystemIdentification1Code string

// ExternalFinancialInstitutionIdentification1Code ...
type ExternalFinancialInstitutionIdentification1Code string

// ExternalLocalInstrument1Code ...
type ExternalLocalInstrument1Code string

// ExternalOrganisationIdentification1Code ...
type ExternalOrganisationIdentification1Code string

// ExternalPersonIdentification1Code ...
type ExternalPersonIdentification1Code string

// ExternalPurpose1Code ...
type ExternalPurpose1Code string

// ExternalServiceLevel1Code ...
type ExternalServiceLevel1Code string

// FIToFICustomerCreditTransferV02 ...
type FIToFICustomerCreditTransferV02 struct {
	GrpHdr      *GroupHeader33                            `xml:"GrpHdr,omitempty"`
	CdtTrfTxInf []*CreditTransferTransactionInformation11 `xml:"CdtTrfTxInf,omitempty"`
}

// FinancialIdentificationSchemeName1Choice ...
type FinancialIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// FinancialInstitutionIdentification7 ...
type FinancialInstitutionIdentification7 struct {
	BIC         string                               `xml:"BIC,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty"`
	Nm          string                               `xml:"Nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty"`
}

// GenericAccountIdentification1 ...
type GenericAccountIdentification1 struct {
	Id      string                    `xml:"Id,omitempty"`
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    string                    `xml:"Issr,omitempty"`
}

// GenericFinancialIdentification1 ...
type GenericFinancialIdentification1 struct {
	Id      string                                    `xml:"Id,omitempty"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    string                                    `xml:"Issr,omitempty"`
}

// GenericOrganisationIdentification1 ...
type GenericOrganisationIdentification1 struct {
	Id      string                                       `xml:"Id,omitempty"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    string                                       `xml:"Issr,omitempty"`
}

// GenericPersonIdentification1 ...
type GenericPersonIdentification1 struct {
	Id      string                                 `xml:"Id,omitempty"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`
	Issr    string                                 `xml:"Issr,omitempty"`
}

// GroupHeader33 ...
type GroupHeader33 struct {
	MsgId             string                                        `xml:"MsgId,omitempty"`
	CreDtTm           string                                        `xml:"CreDtTm,omitempty"`
	BtchBookg         bool                                          `xml:"BtchBookg,omitempty"`
	NbOfTxs           string                                        `xml:"NbOfTxs,omitempty"`
	CtrlSum           float64                                       `xml:"CtrlSum,omitempty"`
	TtlIntrBkSttlmAmt *ActiveCurrencyAndAmount                      `xml:"TtlIntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt     string                                        `xml:"IntrBkSttlmDt,omitempty"`
	SttlmInf          *SettlementInformation13                      `xml:"SttlmInf,omitempty"`
	PmtTpInf          *PaymentTypeInformation21                     `xml:"PmtTpInf,omitempty"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification4 `xml:"InstgAgt,omitempty"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification4 `xml:"InstdAgt,omitempty"`
}

// IBAN2007Identifier ...
type IBAN2007Identifier string

// ISODate ...
type ISODate string

// ISODateTime ...
type ISODateTime string

// ISOTime ...
type ISOTime time.Time

// Instruction3Code ...
type Instruction3Code string

// Instruction4Code ...
type Instruction4Code string

// InstructionForCreditorAgent1 ...
type InstructionForCreditorAgent1 struct {
	Cd       string `xml:"Cd,omitempty"`
	InstrInf string `xml:"InstrInf,omitempty"`
}

// InstructionForNextAgent1 ...
type InstructionForNextAgent1 struct {
	Cd       string `xml:"Cd,omitempty"`
	InstrInf string `xml:"InstrInf,omitempty"`
}

// LocalInstrument2Choice ...
type LocalInstrument2Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// Max10Text ...
type Max10Text string

// Max140Text ...
type Max140Text string

// Max15NumericText ...
type Max15NumericText string

// Max16Text ...
type Max16Text string

// Max2048Text ...
type Max2048Text string

// Max34Text ...
type Max34Text string

// Max35Text ...
type Max35Text string

// Max4Text ...
type Max4Text string

// Max70Text ...
type Max70Text string

// NameAndAddress10 ...
type NameAndAddress10 struct {
	Nm  string          `xml:"Nm,omitempty"`
	Adr *PostalAddress6 `xml:"Adr,omitempty"`
}

// NamePrefix1Code ...
type NamePrefix1Code string

// OrganisationIdentification4 ...
type OrganisationIdentification4 struct {
	BICOrBEI string                                `xml:"BICOrBEI,omitempty"`
	Othr     []*GenericOrganisationIdentification1 `xml:"Othr,omitempty"`
}

// OrganisationIdentificationSchemeName1Choice ...
type OrganisationIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// Party6Choice ...
type Party6Choice struct {
	OrgId  *OrganisationIdentification4 `xml:"OrgId,omitempty"`
	PrvtId *PersonIdentification5       `xml:"PrvtId,omitempty"`
}

// PartyIdentification32 ...
type PartyIdentification32 struct {
	Nm        string           `xml:"Nm,omitempty"`
	PstlAdr   *PostalAddress6  `xml:"PstlAdr,omitempty"`
	Id        *Party6Choice    `xml:"Id,omitempty"`
	CtryOfRes string           `xml:"CtryOfRes,omitempty"`
	CtctDtls  *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

// PaymentIdentification3 ...
type PaymentIdentification3 struct {
	InstrId    string `xml:"InstrId,omitempty"`
	EndToEndId string `xml:"EndToEndId,omitempty"`
	TxId       string `xml:"TxId,omitempty"`
	ClrSysRef  string `xml:"ClrSysRef,omitempty"`
}

// PaymentTypeInformation21 ...
type PaymentTypeInformation21 struct {
	InstrPrty string                  `xml:"InstrPrty,omitempty"`
	ClrChanl  string                  `xml:"ClrChanl,omitempty"`
	SvcLvl    *ServiceLevel8Choice    `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

// PersonIdentification5 ...
type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth            `xml:"DtAndPlcOfBirth,omitempty"`
	Othr            []*GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

// PersonIdentificationSchemeName1Choice ...
type PersonIdentificationSchemeName1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// PhoneNumber ...
type PhoneNumber string

// PostalAddress6 ...
type PostalAddress6 struct {
	AdrTp       string   `xml:"AdrTp,omitempty"`
	Dept        string   `xml:"Dept,omitempty"`
	SubDept     string   `xml:"SubDept,omitempty"`
	StrtNm      string   `xml:"StrtNm,omitempty"`
	BldgNb      string   `xml:"BldgNb,omitempty"`
	PstCd       string   `xml:"PstCd,omitempty"`
	TwnNm       string   `xml:"TwnNm,omitempty"`
	CtrySubDvsn string   `xml:"CtrySubDvsn,omitempty"`
	Ctry        string   `xml:"Ctry,omitempty"`
	AdrLine     []string `xml:"AdrLine,omitempty"`
}

// Priority2Code ...
type Priority2Code string

// Priority3Code ...
type Priority3Code string

// Purpose2Choice ...
type Purpose2Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// ReferredDocumentInformation3 ...
type ReferredDocumentInformation3 struct {
	Tp     *ReferredDocumentType2 `xml:"Tp,omitempty"`
	Nb     string                 `xml:"Nb,omitempty"`
	RltdDt string                 `xml:"RltdDt,omitempty"`
}

// ReferredDocumentType1Choice ...
type ReferredDocumentType1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// ReferredDocumentType2 ...
type ReferredDocumentType2 struct {
	CdOrPrtry *ReferredDocumentType1Choice `xml:"CdOrPrtry,omitempty"`
	Issr      string                       `xml:"Issr,omitempty"`
}

// RegulatoryAuthority2 ...
type RegulatoryAuthority2 struct {
	Nm   string `xml:"Nm,omitempty"`
	Ctry string `xml:"Ctry,omitempty"`
}

// RegulatoryReporting3 ...
type RegulatoryReporting3 struct {
	DbtCdtRptgInd string                            `xml:"DbtCdtRptgInd,omitempty"`
	Authrty       *RegulatoryAuthority2             `xml:"Authrty,omitempty"`
	Dtls          []*StructuredRegulatoryReporting3 `xml:"Dtls,omitempty"`
}

// RegulatoryReportingType1Code ...
type RegulatoryReportingType1Code string

// RemittanceAmount1 ...
type RemittanceAmount1 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`
	DscntApldAmt      *ActiveOrHistoricCurrencyAndAmount `xml:"DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`
	TaxAmt            *ActiveOrHistoricCurrencyAndAmount `xml:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []*DocumentAdjustment1             `xml:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

// RemittanceInformation5 ...
type RemittanceInformation5 struct {
	Ustrd []string                            `xml:"Ustrd,omitempty"`
	Strd  []*StructuredRemittanceInformation7 `xml:"Strd,omitempty"`
}

// RemittanceLocation2 ...
type RemittanceLocation2 struct {
	RmtId             string            `xml:"RmtId,omitempty"`
	RmtLctnMtd        string            `xml:"RmtLctnMtd,omitempty"`
	RmtLctnElctrncAdr string            `xml:"RmtLctnElctrncAdr,omitempty"`
	RmtLctnPstlAdr    *NameAndAddress10 `xml:"RmtLctnPstlAdr,omitempty"`
}

// RemittanceLocationMethod2Code ...
type RemittanceLocationMethod2Code string

// ServiceLevel8Choice ...
type ServiceLevel8Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// SettlementDateTimeIndication1 ...
type SettlementDateTimeIndication1 struct {
	DbtDtTm string `xml:"DbtDtTm,omitempty"`
	CdtDtTm string `xml:"CdtDtTm,omitempty"`
}

// SettlementInformation13 ...
type SettlementInformation13 struct {
	SttlmMtd             string                                        `xml:"SttlmMtd,omitempty"`
	SttlmAcct            *CashAccount16                                `xml:"SttlmAcct,omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"ClrSys,omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification4 `xml:"InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount16                                `xml:"InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification4 `xml:"InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount16                                `xml:"InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification4 `xml:"ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount16                                `xml:"ThrdRmbrsmntAgtAcct,omitempty"`
}

// SettlementMethod1Code ...
type SettlementMethod1Code string

// SettlementTimeRequest2 ...
type SettlementTimeRequest2 struct {
	CLSTm  time.Time `xml:"CLSTm,omitempty"`
	TillTm time.Time `xml:"TillTm,omitempty"`
	FrTm   time.Time `xml:"FrTm,omitempty"`
	RjctTm time.Time `xml:"RjctTm,omitempty"`
}

// StructuredRegulatoryReporting3 ...
type StructuredRegulatoryReporting3 struct {
	Tp   string                             `xml:"Tp,omitempty"`
	Dt   string                             `xml:"Dt,omitempty"`
	Ctry string                             `xml:"Ctry,omitempty"`
	Cd   string                             `xml:"Cd,omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	Inf  []string                           `xml:"Inf,omitempty"`
}

// StructuredRemittanceInformation7 ...
type StructuredRemittanceInformation7 struct {
	RfrdDocInf  []*ReferredDocumentInformation3 `xml:"RfrdDocInf,omitempty"`
	RfrdDocAmt  *RemittanceAmount1              `xml:"RfrdDocAmt,omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2  `xml:"CdtrRefInf,omitempty"`
	Invcr       *PartyIdentification32          `xml:"Invcr,omitempty"`
	Invcee      *PartyIdentification32          `xml:"Invcee,omitempty"`
	AddtlRmtInf []string                        `xml:"AddtlRmtInf,omitempty"`
}
