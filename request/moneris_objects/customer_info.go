package moneris_objects

import (
	"encoding/xml"
	"github.com/moneris-go-api/request"
)

type commonFields struct {
	FirstName    string `xml:"first_name"`
	LastName     string `xml:"last_name"`
	CompanyName  string `xml:"company_name"`
	Address      string `xml:"address"`
	City         string `xml:"city"`
	Province     string `xml:"province"`
	PostalCode   string `xml:"postal_code"`
	Country      string `xml:"country"`
	PhoneNumber  string `xml:"phone_number"`
	Fax          string `xml:"fax"`
	Tax1         string `xml:"tax1"`
	Tax2         string `xml:"tax2"`
	Tax3         string `xml:"tax3"`
	ShippingCost string `xml:"shipping_cost"`
}

type Billing struct {
	commonFields
}

type Shipping struct {
	commonFields
}

type Item struct {
	Name           string `xml:"name"`
	ProductCode    string `xml:"product_code"`
	ExtendedAmount string `xml:"extended_amount"`
	Quantity       int64  `xml:"quantity"`
}

type CustomerInfo struct {
	XMLName         xml.Name `xml:"cust_info"`
	Email           string   `xml:"email"`
	Instructions    string   `xml:"instructions"`
	BillingDetails  Billing  `xml:"billing"`
	ShippingDetails Shipping `xml:"shipping"`
	ItemDetails     Item     `xml:"item"`
}

func NewCustomerInfo(email, instructions string, billing Billing, shipping Shipping, item Item) *CustomerInfo {
	return &CustomerInfo{
		Email:           email,
		Instructions:    instructions,
		BillingDetails:  billing,
		ShippingDetails: shipping,
		ItemDetails:     item,
	}
}

func (c *CustomerInfo) ToXml() string {
	return request.CreateXmlString(c)
}
