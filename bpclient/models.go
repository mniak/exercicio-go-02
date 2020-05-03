package bpclient

// AuthorizationRequest = AuthorizationRequest
type AuthorizationRequest struct {
	MerchantOrderID string  `json:"MerchantOrderId"`
	Payment         Payment `json:"Payment"`
}

// AuthorizationResponse = AuthorizationResponse
type AuthorizationResponse struct {
	MerchantOrderID string  `json:"MerchantOrderId"`
	Payment         Payment `json:"Payment"`
}

// Payment = Payment
type Payment struct {
	Type                   string               `json:"Type"`
	SoftDescriptor         string               `json:"SoftDescriptor"`
	PaymentDateTime        string               `json:"PaymentDateTime"`
	Amount                 int64                `json:"Amount"`
	Installments           int64                `json:"Installments"`
	Interest               string               `json:"Interest"`
	Capture                bool                 `json:"Capture"`
	ProductID              int64                `json:"ProductId"`
	CreditCard             CreditCard           `json:"CreditCard"`
	PinPadInformation      PinPadInformation    `json:"PinPadInformation"`
	ReceivedDate           string               `json:"ReceivedDate"`
	CapturedAmount         int64                `json:"CapturedAmount"`
	CapturedDate           string               `json:"CapturedDate"`
	Provider               *string              `json:"Provider"`
	Status                 int64                `json:"Status"`
	IsSplitted             bool                 `json:"IsSplitted"`
	ReturnMessage          string               `json:"ReturnMessage"`
	ReturnCode             string               `json:"ReturnCode"`
	PaymentID              string               `json:"PaymentId"`
	Currency               *string              `json:"Currency"`
	Country                *string              `json:"Country"`
	Links                  []Link               `json:"Links"`
	ServiceTaxAmount       int64                `json:"ServiceTaxAmount"`
	PrintMessage           []PrintMessage       `json:"PrintMessage"`
	ReceiptInformation     []ReceiptInformation `json:"ReceiptInformation"`
	Receipt                Receipt              `json:"Receipt"`
	AuthorizationCode      string               `json:"AuthorizationCode"`
	ProofOfSale            string               `json:"ProofOfSale"`
	InitializationVersion  int64                `json:"InitializationVersion"`
	ConfirmationStatus     int64                `json:"ConfirmationStatus"`
	EmvResponseData        string               `json:"EmvResponseData"`
	SubordinatedMerchantID *string              `json:"SubordinatedMerchantId"`
}

// Receipt = Receipt
type Receipt struct {
	MerchantName      string `json:"MerchantName"`
	MerchantCity      string `json:"MerchantCity"`
	InputMethod       string `json:"InputMethod"`
	Terminal          string `json:"Terminal"`
	IssuerName        string `json:"IssuerName"`
	Nsu               string `json:"Nsu"`
	MerchantCode      string `json:"MerchantCode"`
	MerchantAddress   string `json:"MerchantAddress"`
	AuthorizationCode string `json:"AuthorizationCode"`
	CardHolder        string `json:"CardHolder"`
	TransactionType   string `json:"TransactionType"`
	MerchantState     string `json:"MerchantState"`
	Date              string `json:"Date"`
	Hour              string `json:"Hour"`
	Value             string `json:"Value"`
	TransactionMode   string `json:"TransactionMode"`
	CardNumber        string `json:"CardNumber"`
}

//ReceiptInformation = ReceiptInformation
type ReceiptInformation struct {
	Field   string `json:"Field"`
	Label   string `json:"Label"`
	Content string `json:"Content"`
}

// Link = Link
type Link struct {
	Method string `json:"Method"`
	Rel    string `json:"Rel"`
	Href   string `json:"Href"`
}

// PrintMessage = PrintMessage
type PrintMessage struct {
	Position string `json:"Position"`
	Message  string `json:"Message"`
}

// CreditCard = CreditCard
type CreditCard struct {
	CardNumber                     string           `json:"CardNumber"`
	ExpirationDate                 string           `json:"ExpirationDate"`
	SecurityCodeStatus             string           `json:"SecurityCodeStatus"`
	SecurityCode                   string           `json:"SecurityCode"`
	BrandID                        int64            `json:"BrandId"`
	IssuerID                       int64            `json:"IssuerId"`
	TruncateCardNumberWhenPrinting bool             `json:"TruncateCardNumberWhenPrinting"`
	InputMode                      string           `json:"InputMode"`
	AuthenticationMethod           string           `json:"AuthenticationMethod"`
	EmvData                        string           `json:"EmvData"`
	IsFallback                     bool             `json:"IsFallback"`
	BrandInformation               BrandInformation `json:"BrandInformation"`
	SaveCard                       bool             `json:"SaveCard"`
}

// BrandInformation = BrandInformation
type BrandInformation struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// PinPadInformation = PinPadInformation
type PinPadInformation struct {
	PhysicalCharacteristics string `json:"PhysicalCharacteristics"`
	ReturnDataInfo          string `json:"ReturnDataInfo"`
	SerialNumber            string `json:"SerialNumber"`
	TerminalID              string `json:"TerminalId"`
}

// Error = Error
type Error struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}
