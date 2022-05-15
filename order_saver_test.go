package main

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"order-service-log-saver/domain"
	"testing"
)

func TestLogSaver_saveCreateRequest(t *testing.T) {
	db, mock, _ := sqlmock.New()

	query := "INSERT INTO request\\(request_id, request_time, type, body\\) values \\(\\?, \\?, \\?, \\?\\)"

	mock.ExpectExec(query).WithArgs(
		"5fbfd3c8-c079-4795-b616-7007284dc678",
		1651843040.069338000,
		"CREATE_ORDER",
		"{\"salePoint\":{\"id\":1,\"spStatusId\":1,\"spName\":\"salePoint\"},\"status\":{\"id\":1,\"stEntityType\":\"Order\",\"stName\":\"NEW\"},\"positions\":[{\"id\":1,\"posCategoryId\":1,\"posName\":\"iPhoneXS\",\"posStatusId\":1}]}",
	).WillReturnResult(sqlmock.NewResult(0, 1))

	request := domain.CreateOrderGatewayRequest{
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
	}

	t.Run("OK_save_create_request", func(t *testing.T) {
		saver := LogSaver{
			connection: db,
		}
		_, err := saver.saveCreateRequest(request)
		assert.NoError(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestLogSaver_saveCreateResponse(t *testing.T) {
	db, mock, _ := sqlmock.New()

	query := "INSERT INTO response\\(request_id, response_time, type, status, body, error_message\\) " +
		"values \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"

	mock.ExpectExec(query).WithArgs(
		"22746a58-09c9-43bf-a326-949fd61b4430",
		1651844353.672617000,
		"CREATE_ORDER",
		"OK",
		"{\"ordId\":1}",
		sql.NullString{
			String: "",
			Valid:  true,
		},
	).WillReturnResult(sqlmock.NewResult(0, 1))

	response := domain.CreateOrderGatewayResponse{
		RequestId:    "22746a58-09c9-43bf-a326-949fd61b4430",
		ResponseTime: 1651844353.672617000,
		Type:         "CREATE_ORDER",
		Status:       "OK",
		Body: domain.CreateOrderResponse{
			OrdId: uint64(1),
		},
	}

	t.Run("OK_save_create_response", func(t *testing.T) {
		saver := LogSaver{
			connection: db,
		}
		_, err := saver.saveCreateResponse(response)

		assert.NoError(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestLogSaver_saveGetRequest(t *testing.T) {
	db, mock, _ := sqlmock.New()

	query := "INSERT INTO request\\(request_id, request_time, type, body\\) values \\(\\?, \\?, \\?, \\?\\)"

	request := domain.GetOrderGatewayRequest{
		RequestId:   "1c0a8eec-7181-4f7c-8866-ccafdc46780e",
		RequestTime: 1651843040.069356000,
		Type:        "GET_ORDER",
		Body:        uint64(1),
	}

	mock.ExpectExec(query).WithArgs(
		"1c0a8eec-7181-4f7c-8866-ccafdc46780e",
		1651843040.069356000,
		"GET_ORDER",
		"1",
	).WillReturnResult(sqlmock.NewResult(0, 1))

	t.Run("OK_save_get_request", func(t *testing.T) {
		saver := LogSaver{connection: db}
		_, err := saver.saveGetRequest(request)
		assert.NoError(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestLogSaver_saveGetResponse(t *testing.T) {
	db, mock, _ := sqlmock.New()

	request := "INSERT INTO response\\(request_id, response_time, type, status, body, error_message\\) " +
		"values \\(\\?, \\?, \\?, \\?, \\?, \\?\\)"

	response := domain.GetOrderGatewayResponse{
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
	}

	mock.ExpectExec(request).WithArgs(
		"f91ece25-9bfc-41ed-88b9-1f30d87aa0ad",
		1651843040.069418000,
		"GET_ORDER",
		"OK",
		"{\"order\":{\"salePoint\":{\"id\":1,\"spStatusId\":1,\"spName\":\"salePoint\"},\"status\":{\"id\":1,\"stEntityType\":\"Order\",\"stName\":\"NEW\"},\"positions\":[{\"id\":1,\"posCategoryId\":1,\"posName\":\"iPhoneXS\",\"posStatusId\":1}]}}",
		sql.NullString{
			String: "",
			Valid:  true,
		},
	).WillReturnResult(sqlmock.NewResult(0, 1))

	t.Run("OK_save_get_response", func(t *testing.T) {
		saver := LogSaver{
			connection: db,
		}
		_, err := saver.saveGetResponse(response)
		assert.NoError(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
