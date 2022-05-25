package request

import (
	"order-service-log-saver/domain"
	"reflect"
	"testing"
)

func TestUnmarshalRequests(t *testing.T) {
	tests := []struct {
		name           string
		messages       [][]byte
		wantGetReqs    []domain.GetOrderGatewayRequest
		wantCreateReqs []domain.CreateOrderGatewayRequest
	}{
		{
			name: "OK",
			messages: [][]byte{
				[]byte(`
					{
					  "requestId": "1c0a8eec-7181-4f7c-8866-ccafdc46780e",
					  "requestTime": 1651843040.069356000,
					  "type": "GET_ORDER",
					  "body": 1
					}`,
				),
				[]byte(`
					{
					  "requestId": "5fbfd3c8-c079-4795-b616-7007284dc678",
					  "requestTime": 1651843040.069338000,
					  "type": "CREATE_ORDER",
					  "body": {
						"salePoint": {
						  "id": 1,
						  "spStatusId": 1,
						  "spName": "salePoint"
						},
						"status": {
						  "id": 1,
						  "stEntityType": "Order",
						  "stName": "NEW"
						},
						"positions": [
						  {
							"id": 1,
							"posCategoryId": 1,
							"posName": "iPhoneXS",
							"posStatusId": 1
						  }
						]
					  }
					} `,
				),
			},
			wantGetReqs: []domain.GetOrderGatewayRequest{
				{
					RequestId:   "1c0a8eec-7181-4f7c-8866-ccafdc46780e",
					RequestTime: 1651843040.069356000,
					Type:        "GET_ORDER",
					Body:        uint64(1),
				},
			},
			wantCreateReqs: []domain.CreateOrderGatewayRequest{
				{
					RequestId:   "5fbfd3c8-c079-4795-b616-7007284dc678",
					RequestTime: 1651843040.069338000,
					Type:        "CREATE_ORDER",
					Body: domain.TransferOrder{
						SalePoint: domain.SalePoint{
							Id:         1,
							SpStatusId: 1,
							SpName:     "salePoint",
						},
						Status: domain.Status{
							Id:           1,
							StEntityType: "Order",
							StName:       "NEW",
						},
						Positions: []domain.Position{
							{
								Id:            1,
								PosCategoryId: 1,
								PosName:       "iPhoneXS",
								PosStatusId:   1,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actGetReqs, actCreateReqs := UnmarshalRequests(tt.messages)
			if !reflect.DeepEqual(actGetReqs, tt.wantGetReqs) {
				t.Errorf("UnmarshalRequests() got = %v, want %v", actGetReqs, tt.wantGetReqs)
			}
			if !reflect.DeepEqual(actCreateReqs, tt.wantCreateReqs) {
				t.Errorf("UnmarshalRequests() got1 = %v, want %v", actCreateReqs, tt.wantCreateReqs)
			}
		})
	}
}
