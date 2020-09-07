package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/moneris-go-api/lib"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpHelper struct {
	baseUrl           string
	IsLiveEnvironment bool
}

var urlComponents map[string]string
var XMLHeaderString = `<?xml version="1.0" encoding="UTF-8"?>`

const (
	MONERIS_PROTOCOL     = "MONERIS_PROTOCOL"
	MONERIS_HOST         = "MONERIS_HOST"
	MONERIS_TEST_HOST    = "MONERIS_TEST_HOST"
	MONERIS_US_HOST      = "MONERIS_US_HOST"
	MONERIS_US_TEST_HOST = "MONERIS_US_TEST_HOST"
	MONERIS_PORT         = "MONERIS_PORT"
	MONERIS_FILE         = "MONERIS_FILE"
	MONERIS_US_FILE      = "MONERIS_US_FILE"
	MONERIS_MPI_FILE     = "MONERIS_MPI_FILE"
	MONERIS_US_MPI_FILE  = "MONERIS_US_MPI_FILE"
	API_VERSION          = ""
	CONNECT_TIMEOUT      = 20
	CLIENT_TIMEOUT       = 35
)

func New(isLiveEnvironment bool) *HttpHelper {
	httpHelper := new(HttpHelper)
	httpHelper.IsLiveEnvironment = isLiveEnvironment

	urlComponents = make(map[string]string)

	urlComponents[MONERIS_PROTOCOL] = "https://"
	urlComponents[MONERIS_HOST] = "www3.moneris.com"
	urlComponents[MONERIS_TEST_HOST] = "esqa.moneris.com"
	urlComponents[MONERIS_US_HOST] = "esplus.moneris.com"
	urlComponents[MONERIS_US_TEST_HOST] = "esplusqa.moneris.com"
	urlComponents[MONERIS_PORT] = "443"
	urlComponents[MONERIS_FILE] = "/gateway2/servlet/MpgRequest"
	urlComponents[MONERIS_US_FILE] = "/gateway_us/servlet/MpgRequest"
	urlComponents[MONERIS_MPI_FILE] = "/mpi/servlet/MpiServlet"
	urlComponents[MONERIS_US_MPI_FILE] = "/mpi/servlet/MpiServlet"
	urlComponents[API_VERSION] = ""

	return httpHelper
}

func (h *HttpHelper) isMpi(transactionType string) bool {
	if transactionType == "txn" || transactionType == "acs" {
		return true
	}
	return false
}

func (h *HttpHelper) formRequestUrl(data lib.TransactionRequestFields) string {
	var formedUrl strings.Builder
	var hostKey strings.Builder
	var path strings.Builder

	formedUrl.WriteString(urlComponents[MONERIS_PROTOCOL]) // Todo - Do a check on all used items in urlComponents and panic if not initialized

	hostKey.WriteString("MONERIS_")
	if data.CountryCode == "US" {
		hostKey.WriteString(data.CountryCode)
		path.WriteString("MONERIS_US")
	}
	if !h.IsLiveEnvironment {
		hostKey.WriteString("_TEST")
	}
	hostKey.WriteString("_HOST")

	if h.isMpi(data.Type) {
		path.WriteString("_MPI")
	}

	path.WriteString("_FILE")

	formedUrl.WriteString(urlComponents[hostKey.String()])
	formedUrl.WriteString(urlComponents[path.String()])

	return formedUrl.String()
}

func (h *HttpHelper) PostRequest(data lib.TransactionRequestFields) []byte {
	requestUrl := h.formRequestUrl(data)
	reqBody, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	resp, err := http.NewRequest("post", requestUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func CreateXmlString(xmlStruct interface{}) string {
	transactionXmlByte, err := xml.MarshalIndent(xmlStruct, " ", " ")
	if err != nil {
		panic(err)
	}
	return string(transactionXmlByte)
}
