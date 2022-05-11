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

	saver := LogSaver{
		connection: db,
	}
	_, err := saver.saveCreateRequest(request)

	assert.NoError(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLogSaver_saveCreateResponse(t *testing.T) {
	type fields struct {
		connection *sql.DB
	}
	type args struct {
		response domain.CreateOrderGatewayResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := LogSaver{
				connection: tt.fields.connection,
			}
			saver.saveCreateResponse(tt.args.response)
		})
	}
}

func TestLogSaver_saveGetRequest(t *testing.T) {
	type fields struct {
		connection *sql.DB
	}
	type args struct {
		request domain.GetOrderGatewayRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := LogSaver{
				connection: tt.fields.connection,
			}
			saver.saveGetRequest(tt.args.request)
		})
	}
}

func TestLogSaver_saveGetResponse(t *testing.T) {
	type fields struct {
		connection *sql.DB
	}
	type args struct {
		response domain.GetOrderGatewayResponse
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := LogSaver{
				connection: tt.fields.connection,
			}
			saver.saveGetResponse(tt.args.response)
		})
	}
}
