package request

import (
	"encoding/xml"
	"errors"
	"github.com/moneris-go-api/lib"
	"strings"
)

type MpgRequest struct {
	TransactionReqFields []lib.TransactionRequestFields
}

func NewMpgRequest(transactionReqFields []lib.TransactionRequestFields) *MpgRequest {
	return &MpgRequest{TransactionReqFields: transactionReqFields}
}

func (m *MpgRequest) ToXml() (string, error) {
	if len(m.TransactionReqFields) < 1 {
		return "", errors.New("No Transaction provided!")
	}

	var xmlToReturn string
	for _, transaction := range m.TransactionReqFields {
		transaction.XMLName = xml.Name{"", transaction.Type}

		transactionType := transaction.Type
		if transaction.CountryCode == "US" && strings.Contains(transactionType, "us_") {
			if strings.Compare(transactionType, "txn") != 0 || strings.Compare(transactionType, "acs") != 0 || strings.Compare(transactionType, "group") != 0 {
				transactionType = "us_" + transactionType
			}
		}

		formXmlString := func(transactionXmlByte []byte) string {
			transactionXmlStr := string(transactionXmlByte)
			var xmlData strings.Builder
			xmlData.WriteString(transactionXmlStr)
			return xmlData.String()
		}

		if strings.Compare(transactionType, "attribute_query") == 0 || strings.Compare(transactionType, "session_query") == 0 {
			risk := struct {
				XMLName xml.Name `xml:"risk"`
				txn     lib.TransactionRequestFields
			}{txn: transaction}

			transactionXmlByte, err := xml.MarshalIndent(risk, " ", " ")
			if err != nil {
				panic(err)
			}
			xmlToReturn = formXmlString(transactionXmlByte)
		} else {
			transactionXmlByte, err := xml.MarshalIndent(transaction, " ", " ")
			if err != nil {
				panic(err)
			}
			xmlToReturn = formXmlString(transactionXmlByte)
		}
	}
	return xmlToReturn, nil
}
