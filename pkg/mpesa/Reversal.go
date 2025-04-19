package mpesa

type ReversePayload struct {
	Initiator              string `json:"Initiator"`
	PassKey                string `json:"PassKey"`
	CommandID              string `json:"CommandID"`
	TransactionID          string `json:"TransactionID"`
	Amount                 string `json:"Amount"`
	ReceiverParty          string `json:"ReceiverParty"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	ResultURL              string `json:"ResultURL"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	Remarks                string `json:"Remarks"`
	Occasion               string `json:"Occasion"`
}

type ReversalResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (d *Daraja) ReverseTransaction(transaction ReversePayload) (*ReversalResponse, *ErrorResponse) {
	secureResponse, err := performSecurePostRequest[ReversalResponse](transaction, endpointReversal, d)
	if err != nil {
		return nil, err
	}
	return &secureResponse.Body, nil
}
