package lib

import "encoding/xml"

type PaymentInstrumentFields struct {
	CreditCardNumber string `xml:"pan,omitempty"`
	ExpiryDate       string `xml:"expdate,omitempty"`

	EcIndicator int `xml:"crypt_type,omitempty"`

	WalletIndicator string `xml:"wallet_indicator,omitempty"`

	MarketIndicator string `xml:"market_indicator,omitempty"`
}

type TransactionRequestFields struct {
	XMLName  xml.Name
	StoreId  string `xml:"-"`
	ApiToken string `xml:"-"`

	Type string `xml:"-"`

	CustomerFields
	PaymentInstrumentFields

	TransactionDescription string `xml:"dynamic_descriptor,omitempty"`

	CountryCode string `xml:"-"`
}

type CustomerFields struct {
	CustomerId  string  `xml:"cust_id,omitempty"`
	CardMatchId string  `xml:"cm_id,omitempty"`
	OrderId     string  `xml:"order_id,omitempty"`
	Amount      float32 `xml:"amount,omitempty"`
}

type Receipt struct {
	XMLName       xml.Name `xml:"receipt"`
	ReceiptId     *string
	ReferenceNum  *string
	ResponseCode  *string
	AuthCode      *string
	TransTime     *string
	TransDate     *string
	TransType     *string
	Complete      bool
	Message       *string
	TransAmount   *string
	CardType      *string
	TransID       *string
	TimedOut      *string
	CorporateCard bool
	MessageId     *string
}

type XmlErrorResponse struct {
	XMLName xml.Name `xml:"response"`
	receipt Receipt
}

type MainRequestEnvelope struct {
	StoreId     string `xml:"store_id"`
	ApiToken    string `xml:"api_token"`
	AppVersion  string `xml:"app_version"`
	EmbeddedXml interface{}
}

type MpiRequestType struct {
	EmbeddedXml interface{}
}
type RequestType struct {
	EmbeddedXml interface{}
}

/*func (tr *TransactionRequestFields) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if tr.XMLName.Local == "structF" {
		return nil
	}
}*/
