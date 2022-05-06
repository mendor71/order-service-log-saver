package domain

type GetOrderResponse struct {
	Order TransferOrder `json:"order"`
}

type CreateOrderResponse struct {
	OrdId uint64 `json:"ordId"`
}

type GetOrderGatewayResponse struct {
	RequestId    string           `json:"requestId"`
	ResponseTime float64          `json:"responseTime"`
	Type         string           `json:"type"`
	Status       string           `json:"status"`
	Body         GetOrderResponse `json:"body"`
	ErrorMessage string           `json:"errorMessage"`
}

type CreateOrderGatewayResponse struct {
	RequestId    string              `json:"requestId"`
	ResponseTime float64             `json:"responseTime"`
	Type         string              `json:"type"`
	Status       string              `json:"status"`
	Body         CreateOrderResponse `json:"body"`
	ErrorMessage string              `json:"errorMessage"`
}
