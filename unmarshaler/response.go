package unmarshaler

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"order-service-log-saver/domain"
	"strings"
)

func UnmarshalResponses(messages [][]byte) ([]domain.GetOrderGatewayResponse, []domain.CreateOrderGatewayResponse) {
	var getOrderResponses []domain.GetOrderGatewayResponse
	var createOrderResponses []domain.CreateOrderGatewayResponse

	for _, msg := range messages {
		if strings.Contains(string(msg), GetOrder) {
			resp, err := unmarshallGetOrderResponse(msg)
			if err != nil {
				continue
			}
			getOrderResponses = append(getOrderResponses, resp)
		} else if strings.Contains(string(msg), CreateOrder) {
			resp, err := unmarshallCreateOrderResponse(msg)
			if err != nil {
				continue
			}
			createOrderResponses = append(createOrderResponses, resp)
		}
	}
	return getOrderResponses, createOrderResponses
}

func unmarshallGetOrderResponse(msg []byte) (domain.GetOrderGatewayResponse, error) {
	var response domain.GetOrderGatewayResponse
	err := json.Unmarshal(msg, &response)
	if err != nil {
		log.Error().Msgf("Can't deserialize GatewayResponse: %s", string(msg))
	}
	return response, err
}

func unmarshallCreateOrderResponse(msg []byte) (domain.CreateOrderGatewayResponse, error) {
	var response domain.CreateOrderGatewayResponse
	err := json.Unmarshal(msg, &response)
	if err != nil {
		log.Error().Msgf("Can't deserialize GatewayResponse: %s", string(msg))
	}
	return response, err
}
