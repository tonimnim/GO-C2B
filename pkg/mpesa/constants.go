package mpesa

const (
	baseUrlSandbox                    = "https://sandbox.safaricom.co.ke"
	baseUrlLive                       = "https://api.safaricom.co.ke"
	endpointAuth                      = "/oauth/v1/generate?grant_type=client_credentials" // The endpoint for generating an access token
	endpointReversal                  = "/mpesa/reversal/v1/request"                       // The endpoint for reversing a transaction
	endpointB2BReq                    = "/mpesa/b2b/v1/paymentrequest"                     // The endpoint for making a B2B payment request
	endpointTxStatus                  = "/mpesa/transactionstatus/v1/query"                // The endpoint for querying a transaction status
	endpointQueryLipanaMpesa          = "/mpesa/stkpushquery/v1/query"                     // The endpoint for querying a Lipa Na Mpesa transaction
	endpointAccountBalance            = "/mpesa/accountbalance/v1/query"                   // The endpoint for querying an M-Pesa account balance
	endpointRegisterConfirmValidation = "/mpesa/c2b/v1/registerurl"                        // The endpoint for confirming a C2BSimulation transaction
	endpointLipaNaMpesa               = "/mpesa/stkpush/v1/processrequest"                 // The endpoint for initiating a Lipa Na Mpesa transaction
	endpointC2B                       = "/mpesa/c2b/v1/simulate"                           // The endpoint for simulating a C2B transaction
)
