package MB

type RequestBody interface {
}

// Request LINK
type Request_LINK struct {
	WalletID         string `json:"walletID"`
	ActionType       string `json:"actionType"`
	SourceName       string `json:"sourceName"`
	SourceNumber     string `json:"sourceNumber"`
	SourceType       string `json:"sourceType"`
	AuthenType       string `json:"authenType"`
	Mobile           string `json:"mobile"`
	NationalId       string `json:"nationalId"`
	DateOpenCard     string `json:"dateOpenCard"`
	DeepLinkAuthen   string `json:"deepLinkAuthen"`
	DeepLinkRegister string `json:"deepLinkRegister"`
	DeviceId         string `json:"deviceId"`
	DueDate          string `json:"dueDate"`
	Fee              int64  `json:"fee"`
	Amount           int64  `json:"amount"`
	PaymentDetails   string `json:"paymentDetails"`
	RequestId        string `json:"requestId"`
	ResourceId       string `json:"resourceId"`
}
