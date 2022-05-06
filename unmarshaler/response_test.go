package unmarshaler

import (
	"order-service-log-saver/domain"
	"reflect"
	"testing"
)

func TestUnmarshalResponses(t *testing.T) {
	tests := []struct {
		name            string
		messages        [][]byte
		wantGetResps    []domain.GetOrderGatewayResponse
		wantCreateResps []domain.CreateOrderGatewayResponse
	}{
		{
			name: "OK",
			messages: [][]byte{
				[]byte(`
					{
					  "requestId": "f91ece25-9bfc-41ed-88b9-1f30d87aa0ad",
					  "responseTime": 1651843040.069418000,
					  "type": "GET_ORDER",
					  "status": "OK",
					  "body": {
						"order": {
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
					  },
					  "errorMessage": null
					}`,
				),
				[]byte(`
					{
					  "requestId": "22746a58-09c9-43bf-a326-949fd61b4430",
					  "responseTime": 1651844353.672617000,
					  "type": "CREATE_ORDER",
					  "status": "OK",
					  "body": {
						"ordId": 1
					  },
					  "errorMessage": null
					} 
				`),
			},
			wantGetResps: []domain.GetOrderGatewayResponse{
				{
					RequestId:    "f91ece25-9bfc-41ed-88b9-1f30d87aa0ad",
					ResponseTime: 1651843040.069418000,
					Type:         "GET_ORDER",
					Status:       "OK",
					Body: domain.GetOrderResponse{
						Order: domain.TransferOrder{
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
			wantCreateResps: []domain.CreateOrderGatewayResponse{
				{
					RequestId:    "22746a58-09c9-43bf-a326-949fd61b4430",
					ResponseTime: 1651844353.672617000,
					Type:         "CREATE_ORDER",
					Status:       "OK",
					Body: domain.CreateOrderResponse{
						OrdId: uint64(1),
					},
				},
			},
		},
		{
			name: "Empty",
			messages: [][]byte{
				[]byte(`
					{
					  "requestId": "f91ece25-9bfc-41ed-88b9-1f30d87aa0ad",
					  "responseTime": 1651843040.069418000,
					  "type": "GET_ORDER",
					  "status": "OK",
					  "body": {},
					  "errorMessage": null
					}`,
				),
				[]byte(`
					{
					  "requestId": "22746a58-09c9-43bf-a326-949fd61b4430",
					  "responseTime": 1651844353.672617000,
					  "type": "CREATE_ORDER",
					  "status": "OK",
					  "body": {},
					  "errorMessage": null
					} 
				`),
			},
			wantGetResps: []domain.GetOrderGatewayResponse{
				{
					RequestId:    "f91ece25-9bfc-41ed-88b9-1f30d87aa0ad",
					ResponseTime: 1651843040.069418000,
					Type:         "GET_ORDER",
					Status:       "OK",
				},
			},
			wantCreateResps: []domain.CreateOrderGatewayResponse{
				{
					RequestId:    "22746a58-09c9-43bf-a326-949fd61b4430",
					ResponseTime: 1651844353.672617000,
					Type:         "CREATE_ORDER",
					Status:       "OK",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actGetResps, actCreateResps := UnmarshalResponses(tt.messages)
			if !reflect.DeepEqual(actGetResps, tt.wantGetResps) {
				t.Errorf("UnmarshalResponses() actGetResps = %v, want %v", actGetResps, tt.wantGetResps)
			}
			if !reflect.DeepEqual(actCreateResps, tt.wantCreateResps) {
				t.Errorf("UnmarshalResponses() actCreateResps = %v, want %v", actCreateResps, tt.wantCreateResps)
			}
		})
	}
}
