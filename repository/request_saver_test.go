package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"order-service-log-saver/domain"
	"testing"
)

func TestLogSaver_saveCreateRequest(t *testing.T) {
	db, mock, _ := sqlmock.New()

	query := "INSERT INTO request\\(request_id, request_time, type, body\\) values \\(\\$1, \\$2, \\$3, \\$4\\)"

	mock.ExpectExec(query).WithArgs(
		"5fbfd3c8-c079-4795-b616-7007284dc678",
		"1.651843040069338E+09",
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
		saver := RequestRepository{
			DbPool: db,
		}
		err := saver.SaveCreate(request)
		assert.NoError(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestLogSaver_saveGetRequest(t *testing.T) {
	db, mock, _ := sqlmock.New()

	query := "INSERT INTO request\\(request_id, request_time, type, body\\) values \\(\\$1, \\$2, \\$3, \\$4\\)"

	request := domain.GetOrderGatewayRequest{
		RequestId:   "1c0a8eec-7181-4f7c-8866-ccafdc46780e",
		RequestTime: 1651843040.069356000,
		Type:        "GET_ORDER",
		Body:        uint64(1),
	}

	mock.ExpectExec(query).WithArgs(
		"1c0a8eec-7181-4f7c-8866-ccafdc46780e",
		"1.651843040069356E+09",
		"GET_ORDER",
		"1",
	).WillReturnResult(sqlmock.NewResult(0, 1))

	t.Run("OK_save_get_request", func(t *testing.T) {
		saver := RequestRepository{DbPool: db}
		err := saver.SaveGet(request)
		assert.NoError(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
