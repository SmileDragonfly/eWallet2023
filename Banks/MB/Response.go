package MB

// Response data object for all api
type DataObject struct {
	RequestId           string `json:"requestId"`
	Otp                 string `json:"otp"`
	FeePayment          string `json:"feePayment"`
	FeeNoVAT            string `json:"feeNoVAT"`
	Vat                 string `json:"vat"`
	RequestToken        string `json:"requestToken"`
	HtmlSource          string `json:"htmlSource"`
	ResourceNumber      string `json:"resourceNumber"`
	Amount              string `json:"amount"`
	Fee                 string `json:"fee"`
	Mobile              string `json:"mobile"`
	RefNo               string `json:"refNo"`
	ResultOFS           string `json:"resultOFS"`
	RootID              string `json:"rootID"`
	ReferenceNumber     string `json:"referenceNumber"`
	OrderId             string `json:"orderId"`
	NapasKey            string `json:"napasKey"`
	DataKey             string `json:"dataKey"`
	ResourceId          string `json:"resourceId"`
	T24ReferenceNumber  string `json:"t24ReferenceNumber"`
	Way4ReferenceNumber string `json:"way4ReferenceNumber"`
	Balance             string `json:"balance"`
	Time                string `json:"time"`
	TransType           string `json:"transType"`
	CustomerName        string `json:"customerName"`
	CustomerId          string `json:"customerId"`
	WalletId            string `json:"walletId"`
	LinkResourceNumber  string `json:"linkResourceNumber"`
	BalanceWallet       string `json:"balanceWallet"`
	MerchantId          string `json:"merchantId"`
	CreationTime        string `json:"creationTime"`
	Currency            string `json:"currency"`
	Result              string `json:"result"`
	NameCard            string `json:"nameCard"`
	NumberCard          string `json:"numberCard"`
	Type                string `json:"type"`
	TransactionId       string `json:"transactionId"`
}

// Reponse Authorization Token
type Response_AuthorizationToken struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Issued_at    string `json:"issued_at"`
}

// Reponse LINK
type Response_LINK struct {
	ClientMessageId string     `json:"clientMessageId"`
	ErrorCode       string     `json:"errorCode"`
	ErrorDesc       []string   `json:"errorDesc"`
	Data            DataObject `json:"data"`
}

// Reponse UNLINK
type Response_UNLINK struct {
	ClientMessageId string     `json:"clientMessageId"`
	ErrorCode       string     `json:"errorCode"`
	ErrorDesc       []string   `json:"errorDesc"`
	Data            DataObject `json:"data"`
}
