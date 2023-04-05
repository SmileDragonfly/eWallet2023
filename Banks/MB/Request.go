package MB

type RequestBody interface {
}

// Request LINK
type Request_LINK struct {
	WalletId         string `json:"walletId"`
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

// Request LINK
type Request_UNLINK struct {
	WalletId   string `json:"walletId"`
	ResourceId string `json:"resourceId"`
	AuthenType string `json:"authenType"`
	Mobile     string `json:"mobile"`
	Fee        int64  `json:"fee"`
}
