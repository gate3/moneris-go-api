package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HttpHelper struct {
	baseUrl string
}

const (
	MONERIS_PROTOCOL = "https://"
	MONERIS_HOST = "www3.moneris.com"
	MONERIS_TEST_HOST = "esqa.moneris.com"
	MONERIS_US_HOST = "esplus.moneris.com"
	MONERIS_US_TEST_HOST = "esplusqa.moneris.com"
	MONERIS_PORT = "443"
	MONERIS_FILE = "/gateway2/servlet/MpgRequest"
	MONERIS_US_FILE = "/gateway_us/servlet/MpgRequest"
	MONERIS_MPI_FILE = "/mpi/servlet/MpiServlet"
	MONERIS_US_MPI_FILE = "/mpi/servlet/MpiServlet"
	API_VERSION  = ""
	CONNECT_TIMEOUT = 20
	CLIENT_TIMEOUT = 35
)

func (h *HttpHelper) FormRequestUrl (path string) string {
	if h.baseUrl == "" {
		h.baseUrl = MONERIS_HOST + MONERIS_TEST_HOST
	}
	return h.baseUrl + "/" + path
}

func (h *HttpHelper) Post (path string, data map[string]string) []byte {
	requestUrl := h.FormRequestUrl(path)
	reqBody, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(reqBody))
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
