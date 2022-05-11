package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"order-service-log-saver/domain"
	"reflect"
)

type LogSaver struct {
	connection *sql.DB
}

func NewLogSaver() *LogSaver {
	return &LogSaver{createDbConnection()}
}

func (saver LogSaver) saveGetRequest(request domain.GetOrderGatewayRequest) {
	bodyJson, marshalErr := json.Marshal(request.Body)
	if marshalErr != nil {
		log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
		return
	}
	_, err := saver.connection.Exec(
		"INSERT INTO request(request_id, request_time, type, body) values ($1, $2, $3, $4)",
		request.RequestId, request.RequestTime, request.Type, string(bodyJson),
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return
	}
}

func (saver LogSaver) saveCreateRequest(request domain.CreateOrderGatewayRequest) (int64, error) {
	bodyJson, marshalErr := json.Marshal(request.Body)
	if marshalErr != nil {
		log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
		return -1, marshalErr
	}

	res, err := saver.connection.Exec(
		"INSERT INTO request(request_id, request_time, type, body) values (?, ?, ?, ?)",
		request.RequestId, request.RequestTime, request.Type, string(bodyJson),
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return -1, marshalErr
	}

	return res.LastInsertId()
}

func (saver LogSaver) saveGetResponse(response domain.GetOrderGatewayResponse) {
	var bodyJson []byte
	var marshalErr error

	if !reflect.DeepEqual(response.Body, domain.GetOrderResponse{}) {
		bodyJson, marshalErr = json.Marshal(response.Body)
		if marshalErr != nil {
			log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
			return
		}
	}

	_, err := saver.connection.Exec(
		"INSERT INTO response(request_id, response_time, type, status, body, error_message) "+
			"values ($1, $2, $3, $4, $5, $6)",
		response.RequestId, response.ResponseTime, response.Type, response.Status, string(bodyJson), response.ErrorMessage,
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return
	}
}

func (saver LogSaver) saveCreateResponse(response domain.CreateOrderGatewayResponse) {
	var bodyJson []byte
	var marshalErr error

	if !reflect.DeepEqual(response.Body, domain.CreateOrderResponse{}) {
		bodyJson, marshalErr = json.Marshal(response.Body)
		if marshalErr != nil {
			log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
			return
		}
	}

	_, err := saver.connection.Exec(
		"INSERT INTO response(request_id, response_time, type, status, body, error_message) "+
			"values ($1, $2, $3, $4, $5, $6)",
		response.RequestId, response.ResponseTime, response.Type, response.Status, string(bodyJson), response.ErrorMessage,
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return
	}
}

const (
	host     = "localhost"
	port     = 5433
	user     = "order-service-log"
	password = "order-service-log"
	dbName   = "order-service-log"
)

func createDbConnection() *sql.DB {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", user, password, host, port, dbName)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return db
}
