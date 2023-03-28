package Interfaces

type TransactionInterface interface {
	Link(data []byte) ([]byte, error)
	RequestOTP(data []byte) ([]byte, error)
	RequestOTPForLink(data []byte) ([]byte, error)
	TopupByCard(data []byte) ([]byte, error)
	TopupByAccount(data []byte) ([]byte, error)
	PayByCard(data []byte) ([]byte, error)
	PayByAccount(data []byte) ([]byte, error)
	Withdraw(data []byte) ([]byte, error)
	TransferToInternalCard(data []byte) ([]byte, error)
	TransferToInternalAccount(data []byte) ([]byte, error)
	CancelLink(data []byte) ([]byte, error)
	TransferToExternalCard(data []byte) ([]byte, error)
	TransferToExternalAccount(data []byte) ([]byte, error)
	TransferCardless(data []byte) ([]byte, error)
	TransferVisaMasterCard(data []byte) ([]byte, error)
	QueryBalance(data []byte) ([]byte, error)
	QueryStatus(data []byte) ([]byte, error)
	CheckGuaranteeBlance(data []byte) ([]byte, error)
	CheckCustomerInfo(data []byte) ([]byte, error)
	PushNotification(data []byte) ([]byte, error)
}
