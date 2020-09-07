package moneris_objects

type MpgTransaction struct {
	Transaction string
	Recurring   bool
	// Cvd
	CardVerificationDigit      string
	Cof                        CustomerOnFile
	McpRateInfo                string
	AddressVerificationService string
	AchInfo                    string
	// Convenience Fee
	ConvFee              string
	SessionAccountInfo   string
	AttributeAccountInfo string
	Level23Data          string
}

func NewMpgTransaction(
	transaction, cardVerificationDigit, mcpRateInfo, addressVerificationService, achInfo, convFee, sessionAccountInfo, attributeAccountInfo, level23Data string,
	recurring bool,
	customerOnFile CustomerOnFile) *MpgTransaction {

	m := MpgTransaction{}

	m.Transaction = transaction
	m.Recurring = recurring
	m.CardVerificationDigit = cardVerificationDigit
	m.Cof = customerOnFile
	m.McpRateInfo = mcpRateInfo
	m.AddressVerificationService = addressVerificationService
	m.AchInfo = achInfo
	m.ConvFee = convFee
	m.SessionAccountInfo = sessionAccountInfo
	m.AttributeAccountInfo = attributeAccountInfo
	m.Level23Data = level23Data

	return &m
}
