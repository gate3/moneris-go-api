package purchase

import (
	"github.com/moneris-go-api/lib"
	"github.com/moneris-go-api/request"
	mo "github.com/moneris-go-api/request/moneris_objects"

	// "github.com/moneris-go-api/request/moneris_objects"
	"log"
)

func (p *Purchase) ExecuteTransaction(transactionFields lib.TransactionRequestFields) {

	mpgTxn := new(mo.MpgTransaction)

	cof := mo.NewCustomerOnFile("U", "2", "168451306048014")

	mpgTxn.Cof = *cof

	mpgReq := &request.MpgRequest{TransactionReqFields: []lib.TransactionRequestFields{transactionFields}}

	log.Println(mpgReq.ToXml())
	//cinfo moneris_objects.CustomerInfo
	// p.httpHelper.PostRequest(transactionFields)
	//trxFields := []lib.TransactionRequestFields{transactionFields}
	//mpgRequest := request.NewMpgRequest(trxFields)
	//log.Println(mpgRequest.ToXml())

	//c := cinfo.ToXml()
	//log.Println(c)
}
