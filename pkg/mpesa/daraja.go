package mpesa

import (
	"time"
)

type C2BPayload struct {
	ShortCode     string `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        string `json:"Amount"`
	Msisdn        string `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}

type C2BResponse struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
	ResponseCode             string `json:"ResponseCode"`
	ResultDesc               string `json:"ResultDesc"`
	ResultCode               string `json:"ResultCode"`
}

// darajaAuthorizationError is a custom error type that extends the default error.
type darajaAuthorizationError struct {
	Context string
	Err     error
}

func (e *darajaAuthorizationError) Error() string {
	return e.Context + ": " + e.Err.Error()
}

const (
	ENVIROMENT_SANDBOX    = "sandbox"
	ENVIROMENT_PRODUCTION = "production"
)

type Environment string

type Daraja struct {
	authorization  Authorization
	environment    Environment
	nextAuthTime   time.Time
	ConsumerKey    string
	ConsumerSecret string
}

// DarajaAPI defines the interface for M-Pesa API operations
type DarajaAPI interface {
	Authorize() (*Authorization, error)
	ReverseTransaction(transaction ReversePayload) (*ReversalResponse, *ErrorResponse)
	MakeSTKPushRequest(mpesaConfig LipaNaMpesaPayload) (*LipaNaMpesaResponse, *ErrorResponse)
	MakeC2BPaymentV2(c2b C2BPayload) (*C2BResponse, *ErrorResponse)
}

var darajaAPI *Daraja

func NewDaraja(consumerKey, consumerSecret string, env Environment) *Daraja {
	if darajaAPI == nil {
		darajaAPI = &Daraja{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			environment:    env,
		}
	}
	return darajaAPI
}

func (d *Daraja) Authorize() (*Authorization, error) {
	authTimeStart := time.Now()
	auth, err := NewAuthorization(d.ConsumerKey, d.ConsumerSecret, d.environment)
	if err != nil {
		return nil, &darajaAuthorizationError{Context: "authorization failed", Err: err}
	}

	expiry, err := time.ParseDuration(auth.ExpiresIn + "s")
	if err != nil {
		return nil, &darajaAuthorizationError{Context: "failed to parse authorization expiry", Err: err}
	}

	d.nextAuthTime = authTimeStart.Add(expiry)
	d.authorization = *auth

	return auth, nil
}
