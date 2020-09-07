package moneris_objects

import "encoding/xml"

type CustomerOnFile struct {
	XMLName            xml.Name `xml:"cof_info"`
	PaymentIndicator   string   `xml:"payment_indicator"`
	PaymentInformation string   `xml:"payment_information"`
	IssuerId           string   `xml:"issuer_id"`
}

func NewCustomerOnFile(paymentIndicator, paymentInformation, issuerId string) *CustomerOnFile {
	return &CustomerOnFile{
		PaymentInformation: paymentInformation,
		PaymentIndicator:   paymentIndicator,
		IssuerId:           issuerId,
	}
}

func (c *CustomerOnFile) ToXML() string {
	xmlString, _ := xml.MarshalIndent(c, "", " ")
	return string(xmlString)
}
