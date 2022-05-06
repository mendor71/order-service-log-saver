package domain

type CreateOrderGatewayRequest struct {
	RequestId   string        `json:"requestId"`
	RequestTime float64       `json:"requestTime"`
	Type        string        `json:"type"`
	Body        TransferOrder `json:"body"`
}

type GetOrderGatewayRequest struct {
	RequestId   string  `json:"requestId"`
	RequestTime float64 `json:"requestTime"`
	Type        string  `json:"type"`
	Body        uint64  `json:"body"`
}
