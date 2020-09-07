package purchase

import (
	"github.com/moneris-go-api/config"
	"github.com/moneris-go-api/request"
)

type Purchase struct {
	config *config.Config
	*request.HttpHelper
	*request.MpgRequest
}

func NewPurchaseTransaction(cfg *config.Config) *Purchase {
	httpHelper := request.New(cfg.IsLiveEnvironment)
	mpgRequest := new(request.MpgRequest)

	return &Purchase{
		config:     cfg,
		HttpHelper: httpHelper,
		MpgRequest: mpgRequest,
	}
}
