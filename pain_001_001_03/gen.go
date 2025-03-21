package pain_001_001_03

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

// AmountType3Choice ...
type AmountType3Choice struct {
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty"`
	EqvtAmt  *EquivalentAmount2                 `xml:"EqvtAmt,omitempty"`
}

// AnyBICIdentifier ...
type AnyBICIdentifier string

// Authorisation1Choice ...
type Authorisation1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// Authorisation1Code ...
type Authorisation1Code string

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

// Cheque6 ...
type Cheque6 struct {
	ChqTp       string                       `xml:"ChqTp,omitempty"`
	ChqNb       string                       `xml:"ChqNb,omitempty"`
	ChqFr       *NameAndAddress10            `xml:"ChqFr,omitempty"`
	DlvryMtd    *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`
	DlvrTo      *NameAndAddress10            `xml:"DlvrTo,omitempty"`
	InstrPrty   string                       `xml:"InstrPrty,omitempty"`
	ChqMtrtyDt  string                       `xml:"ChqMtrtyDt,omitempty"`
	FrmsCd      string                       `xml:"FrmsCd,omitempty"`
	MemoFld     []string                     `xml:"MemoFld,omitempty"`
	RgnlClrZone string                       `xml:"RgnlClrZone,omitempty"`
	PrtLctn     string                       `xml:"PrtLctn,omitempty"`
}

// ChequeDelivery1Code ...
type ChequeDelivery1Code string

// ChequeDeliveryMethod1Choice ...
type ChequeDeliveryMethod1Choice struct {
	Cd    string `xml:"Cd,omitempty"`
	Prtry string `xml:"Prtry,omitempty"`
}

// ChequeType2Code ...
type ChequeType2Code string

// ClearingSystemIdentification2Choice ...
type ClearingSystemIdentification2Choice struct {
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

// CreditTransferTransactionInformation10 ...
type CreditTransferTransactionInformation10 struct {
	PmtId           *PaymentIdentification1                       `xml:"PmtId,omitempty"`
	PmtTpInf        *PaymentTypeInformation19                     `xml:"PmtTpInf,omitempty"`
	Amt             *AmountType3Choice                            `xml:"Amt,omitempty"`
	XchgRateInf     *ExchangeRateInformation1                     `xml:"XchgRateInf,omitempty"`
	ChrgBr          string                                        `xml:"ChrgBr,omitempty"`
	ChqInstr        *Cheque6                                      `xml:"ChqInstr,omitempty"`
	UltmtDbtr       *PartyIdentification32                        `xml:"UltmtDbtr,omitempty"`
	IntrmyAgt1      *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct  *CashAccount16                                `xml:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2      *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct  *CashAccount16                                `xml:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3      *BranchAndFinancialInstitutionIdentification4 `xml:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct  *CashAccount16                                `xml:"IntrmyAgt3Acct,omitempty"`
	CdtrAgt         *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`
	CdtrAgtAcct     *CashAccount16                                `xml:"CdtrAgtAcct,omitempty"`
	Cdtr            *PartyIdentification32                        `xml:"Cdtr,omitempty"`
	CdtrAcct        *CashAccount16                                `xml:"CdtrAcct,omitempty"`
	UltmtCdtr       *PartyIdentification32                        `xml:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt []*InstructionForCreditorAgent1               `xml:"InstrForCdtrAgt,omitempty"`
	InstrForDbtrAgt string                                        `xml:"InstrForDbtrAgt,omitempty"`
	Purp            *Purpose2Choice                               `xml:"Purp,omitempty"`
	RgltryRptg      []*RegulatoryReporting3                       `xml:"RgltryRptg,omitempty"`
	Tax             *TaxInformation3                              `xml:"Tax,omitempty"`
	RltdRmtInf      []*RemittanceLocation2                        `xml:"RltdRmtInf,omitempty"`
	RmtInf          *RemittanceInformation5                       `xml:"RmtInf,omitempty"`
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

// CustomerCreditTransferInitiationV03 ...
type CustomerCreditTransferInitiationV03 struct {
	GrpHdr *GroupHeader32                    `xml:"GrpHdr,omitempty"`
	PmtInf []*PaymentInstructionInformation3 `xml:"PmtInf,omitempty"`
}

// DateAndPlaceOfBirth ...
type DateAndPlaceOfBirth struct {
	BirthDt     string `xml:"BirthDt,omitempty"`
	PrvcOfBirth string `xml:"PrvcOfBirth,omitempty"`
	CityOfBirth string `xml:"CityOfBirth,omitempty"`
	CtryOfBirth string `xml:"CtryOfBirth,omitempty"`
}

// DatePeriodDetails ...
type DatePeriodDetails struct {
	FrDt string `xml:"FrDt,omitempty"`
	ToDt string `xml:"ToDt,omitempty"`
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

// EquivalentAmount2 ...
type EquivalentAmount2 struct {
	Amt      *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
	CcyOfTrf string                             `xml:"CcyOfTrf,omitempty"`
}

// ExchangeRateInformation1 ...
type ExchangeRateInformation1 struct {
	XchgRate float64 `xml:"XchgRate,omitempty"`
	RateTp   string  `xml:"RateTp,omitempty"`
	CtrctId  string  `xml:"CtrctId,omitempty"`
}

// ExchangeRateType1Code ...
type ExchangeRateType1Code string

// ExternalAccountIdentification1Code ...
type ExternalAccountIdentification1Code string

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

// GroupHeader32 ...
type GroupHeader32 struct {
	MsgId    string                                        `xml:"MsgId,omitempty"`
	CreDtTm  string                                        `xml:"CreDtTm,omitempty"`
	Authstn  []*Authorisation1Choice                       `xml:"Authstn,omitempty"`
	NbOfTxs  string                                        `xml:"NbOfTxs,omitempty"`
	CtrlSum  float64                                       `xml:"CtrlSum,omitempty"`
	InitgPty *PartyIdentification32                        `xml:"InitgPty,omitempty"`
	FwdgAgt  *BranchAndFinancialInstitutionIdentification4 `xml:"FwdgAgt,omitempty"`
}

// IBAN2007Identifier ...
type IBAN2007Identifier string

// ISODate ...
type ISODate string

// ISODateTime ...
type ISODateTime string

// Instruction3Code ...
type Instruction3Code string

// InstructionForCreditorAgent1 ...
type InstructionForCreditorAgent1 struct {
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

// Max128Text ...
type Max128Text string

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

// Number ...
type Number float64

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

// PaymentIdentification1 ...
type PaymentIdentification1 struct {
	InstrId    string `xml:"InstrId,omitempty"`
	EndToEndId string `xml:"EndToEndId,omitempty"`
}

// PaymentInstructionInformation3 ...
type PaymentInstructionInformation3 struct {
	PmtInfId        string                                        `xml:"PmtInfId,omitempty"`
	PmtMtd          string                                        `xml:"PmtMtd,omitempty"`
	BtchBookg       bool                                          `xml:"BtchBookg,omitempty"`
	NbOfTxs         string                                        `xml:"NbOfTxs,omitempty"`
	CtrlSum         float64                                       `xml:"CtrlSum,omitempty"`
	PmtTpInf        *PaymentTypeInformation19                     `xml:"PmtTpInf,omitempty"`
	ReqdExctnDt     string                                        `xml:"ReqdExctnDt,omitempty"`
	PoolgAdjstmntDt string                                        `xml:"PoolgAdjstmntDt,omitempty"`
	Dbtr            *PartyIdentification32                        `xml:"Dbtr,omitempty"`
	DbtrAcct        *CashAccount16                                `xml:"DbtrAcct,omitempty"`
	DbtrAgt         *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`
	DbtrAgtAcct     *CashAccount16                                `xml:"DbtrAgtAcct,omitempty"`
	UltmtDbtr       *PartyIdentification32                        `xml:"UltmtDbtr,omitempty"`
	ChrgBr          string                                        `xml:"ChrgBr,omitempty"`
	ChrgsAcct       *CashAccount16                                `xml:"ChrgsAcct,omitempty"`
	ChrgsAcctAgt    *BranchAndFinancialInstitutionIdentification4 `xml:"ChrgsAcctAgt,omitempty"`
	CdtTrfTxInf     []*CreditTransferTransactionInformation10     `xml:"CdtTrfTxInf,omitempty"`
}

// PaymentMethod3Code ...
type PaymentMethod3Code string

// PaymentTypeInformation19 ...
type PaymentTypeInformation19 struct {
	InstrPrty string                  `xml:"InstrPrty,omitempty"`
	SvcLvl    *ServiceLevel8Choice    `xml:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty"`
}

// PercentageRate ...
type PercentageRate float64

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

// TaxAmount1 ...
type TaxAmount1 struct {
	Rate         float64                            `xml:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`
	Dtls         []*TaxRecordDetails1               `xml:"Dtls,omitempty"`
}

// TaxAuthorisation1 ...
type TaxAuthorisation1 struct {
	Titl string `xml:"Titl,omitempty"`
	Nm   string `xml:"Nm,omitempty"`
}

// TaxInformation3 ...
type TaxInformation3 struct {
	Cdtr            *TaxParty1                         `xml:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"Dbtr,omitempty"`
	AdmstnZn        string                             `xml:"AdmstnZn,omitempty"`
	RefNb           string                             `xml:"RefNb,omitempty"`
	Mtd             string                             `xml:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty"`
	Dt              string                             `xml:"Dt,omitempty"`
	SeqNb           float64                            `xml:"SeqNb,omitempty"`
	Rcrd            []*TaxRecord1                      `xml:"Rcrd,omitempty"`
}

// TaxParty1 ...
type TaxParty1 struct {
	TaxId  string `xml:"TaxId,omitempty"`
	RegnId string `xml:"RegnId,omitempty"`
	TaxTp  string `xml:"TaxTp,omitempty"`
}

// TaxParty2 ...
type TaxParty2 struct {
	TaxId   string             `xml:"TaxId,omitempty"`
	RegnId  string             `xml:"RegnId,omitempty"`
	TaxTp   string             `xml:"TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty"`
}

// TaxPeriod1 ...
type TaxPeriod1 struct {
	Yr     string             `xml:"Yr,omitempty"`
	Tp     string             `xml:"Tp,omitempty"`
	FrToDt *DatePeriodDetails `xml:"FrToDt,omitempty"`
}

// TaxRecord1 ...
type TaxRecord1 struct {
	Tp       string      `xml:"Tp,omitempty"`
	Ctgy     string      `xml:"Ctgy,omitempty"`
	CtgyDtls string      `xml:"CtgyDtls,omitempty"`
	DbtrSts  string      `xml:"DbtrSts,omitempty"`
	CertId   string      `xml:"CertId,omitempty"`
	FrmsCd   string      `xml:"FrmsCd,omitempty"`
	Prd      *TaxPeriod1 `xml:"Prd,omitempty"`
	TaxAmt   *TaxAmount1 `xml:"TaxAmt,omitempty"`
	AddtlInf string      `xml:"AddtlInf,omitempty"`
}

// TaxRecordDetails1 ...
type TaxRecordDetails1 struct {
	Prd *TaxPeriod1                        `xml:"Prd,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`
}

// TaxRecordPeriod1Code ...
type TaxRecordPeriod1Code string
