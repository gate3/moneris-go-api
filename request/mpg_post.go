package request

import (
	"errors"
	"github.com/moneris-go-api/lib"
	"strings"
)

type MpgPost struct {
	StoreId    string
	Apitoken   string
	AppVersion string
	IsMpi      bool

	*MpgRequest
}

func NewMpgPost(storeId, apiToken, appVersion string, isMpi bool, mpgRequest *MpgRequest) *MpgPost {
	m := MpgPost{}
	m.StoreId = storeId
	m.Apitoken = apiToken
	m.AppVersion = appVersion
	m.IsMpi = isMpi
	m.MpgRequest = mpgRequest

	return &m
}

func (m *MpgPost) toXML() (string, error) {
	mpgRequestXml, err := m.MpgRequest.ToXml()
	if err != nil {
		return "", errors.New("Xml Could not be formed!")
	}

	var xmlString strings.Builder
	xmlString.WriteString(XMLHeaderString)

	mainRequestEnvelope := lib.MainRequestEnvelope{}
	mainRequestEnvelope.StoreId = m.StoreId
	mainRequestEnvelope.ApiToken = m.Apitoken
	mainRequestEnvelope.AppVersion = m.AppVersion
	mainRequestEnvelope.EmbeddedXml = mpgRequestXml

	if m.IsMpi {
		mpiXml := lib.MpiRequestType{
			EmbeddedXml: mainRequestEnvelope,
		}

		formedXml := CreateXmlString(mpiXml)
		xmlString.WriteString(formedXml)
		return xmlString.String(), nil
	}

	requestXml := lib.RequestType{
		EmbeddedXml: mainRequestEnvelope,
	}
	formedXml := CreateXmlString(requestXml)
	xmlString.WriteString(formedXml)
	return xmlString.String(), nil
}
