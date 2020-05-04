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
	Type                   PaymentType          `json:"Type"`
	SoftDescriptor         string               `json:"SoftDescriptor"`
	PaymentDateTime        string               `json:"PaymentDateTime"`
	Amount                 int64                `json:"Amount"`
	Installments           int64                `json:"Installments"`
	Interest               Interest             `json:"Interest"`
	Capture                bool                 `json:"Capture"`
	ProductID              int64                `json:"ProductId"`
	CreditCard             CreditCard           `json:"CreditCard"`
	PinPadInformation      PinPadInformation    `json:"PinPadInformation"`
	ReceivedDate           string               `json:"ReceivedDate"`
	CapturedAmount         int64                `json:"CapturedAmount"`
	CapturedDate           string               `json:"CapturedDate"`
	Provider               *string              `json:"Provider"`
	Status                 Status               `json:"Status"`
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
	ConfirmationStatus     ConfirmationStatus   `json:"ConfirmationStatus"`
	EmvResponseData        string               `json:"EmvResponseData"`
	SubordinatedMerchantID *string              `json:"SubordinatedMerchantId"`
}

// Status indicates the status of a payment
type Status int

const (
	// StatusNotFinished = StatusNotFinished
	StatusNotFinished Status = 0
	// StatusAuthorized = StatusAuthorized
	StatusAuthorized Status = 1
	// StatusPaid = StatusPaid
	StatusPaid Status = 2
	// StatusDenied = StatusDenied
	StatusDenied Status = 3
	// StatusVoided = StatusVoided
	StatusVoided Status = 10
	// StatusAborted = StatusAborted
	StatusAborted Status = 13
)

// Interest indicates who is charging the interest
type Interest string

const (
	// InterestByMerchant is an interest type
	InterestByMerchant Interest = "ByMerchant"
)

// PaymentType indicates a type of payment (credit/debit/voucher)
type PaymentType string

const (
	// PhysicalCreditCard is a physical payment with credit card
	PhysicalCreditCard PaymentType = "PhysicalCreditCard"
	// PhysicalDebitCard is a physical payment with debit card
	PhysicalDebitCard PaymentType = "PhysicalDebitCard"
	// PhysicalVoucherCard is a physical payment with voucher card
	PhysicalVoucherCard PaymentType = "PhysicalVoucherCard"
)

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
	InputMode                      InputMode        `json:"InputMode"`
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

type ConfirmationResponse struct {
	ConfirmationStatus ConfirmationStatus `json:"ConfirmationStatus"`
	Status             Status             `json:"Status"`
	ReturnCode         string             `json:"ReturnCode"`
	ReturnMessage      string             `json:"ReturnMessage"`
	Links              []Link             `json:"Links"`
}

// ConfirmationStatus indicates the state of the confirmation
type ConfirmationStatus int

const (
	// Pending is pending
	Pending ConfirmationStatus = 0
	// Confirmed is confirmed
	Confirmed ConfirmationStatus = 1
	// Reversed is reversed
	Reversed ConfirmationStatus = 2
)

// CancellationRequest = CancellationRequest
type CancellationRequest struct {
	MerchantVoidID   string `json:"MerchantVoidId"`
	MerchantVoidDate string `json:"MerchantVoidDate"`
	Card             Card   `json:"Card"`
}

type CancellationResponse struct {
	VoidID                string         `json:"VoidId"`
	InitializationVersion int64          `json:"InitializationVersion"`
	PrintMessage          []PrintMessage `json:"PrintMessage"`
	Receipt               Receipt        `json:"Receipt"`
	Status                Status         `json:"Status"`
	ReturnCode            string         `json:"ReturnCode"`
	ReturnMessage         string         `json:"ReturnMessage"`
	Links                 []Link         `json:"Links"`
}

// Card = Card
type Card struct {
	InputMode  InputMode
	CardNumber string `json:"CardNumber"`
}

// InputMode indicates a input mode (typed, strip or emv)
type InputMode string

const (
	// InputTyped is the typed input mode
	InputTyped InputMode = "Typed"
)

type CancellationReversalResponse struct {
	CancellationStatus CancellationStatus `json:"CancellationStatus"`
	Status             Status             `json:"Status"`
	ReturnCode         string             `json:"ReturnCode"`
	ReturnMessage      string             `json:"ReturnMessage"`
	Links              []Link             `json:"Links"`
}

// CancellationStatus represents the status of the cancellation
type CancellationStatus int

const (
	// CancellationNotFinished = CancellationNotFinished
	CancellationNotFinished CancellationStatus = 0
	// CancellationAuthorized = CancellationAuthorized
	CancellationAuthorized CancellationStatus = 1
	// CancellationDenied = CancellationDenied
	CancellationDenied CancellationStatus = 2
	// CancellationReversed = CancellationReversed
	CancellationReversed CancellationStatus = 4
)
