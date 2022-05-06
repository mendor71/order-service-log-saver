package unmarshaler

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"order-service-log-saver/domain"
	"strings"
)

func UnmarshalRequests(messages [][]byte) ([]domain.GetOrderGatewayRequest, []domain.CreateOrderGatewayRequest) {
	var getOrderRequests []domain.GetOrderGatewayRequest
	var createOrderRequests []domain.CreateOrderGatewayRequest

	for _, msg := range messages {
		if strings.Contains(string(msg), GetOrder) {
			req, err := unmarshalGetOrderRequest(msg)
			if err != nil {
				continue
			}
			getOrderRequests = append(getOrderRequests, req)
		} else if strings.Contains(string(msg), CreateOrder) {
			req, err := unmarshalCreateOrderRequest(msg)
			if err != nil {
				continue
			}
			createOrderRequests = append(createOrderRequests, req)
		}
	}
	return getOrderRequests, createOrderRequests
}

func unmarshalGetOrderRequest(msg []byte) (domain.GetOrderGatewayRequest, error) {
	var request domain.GetOrderGatewayRequest
	err := json.Unmarshal(msg, &request)
	if err != nil {
		log.Error().Msgf("Can't deserialize GatewayRequest: %s", string(msg))
	}
	return request, err
}

func unmarshalCreateOrderRequest(msg []byte) (domain.CreateOrderGatewayRequest, error) {
	var request domain.CreateOrderGatewayRequest
	err := json.Unmarshal(msg, &request)
	if err != nil {
		log.Error().Msgf("Can't deserialize GatewayRequest: %s", string(msg))
	}
	return request, err
}
