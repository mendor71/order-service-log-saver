package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"order-service-log-saver/domain"
	"os"
	"reflect"
)

type LogSaver struct {
	connection *sql.DB
}

func NewLogSaver() *LogSaver {
	host := os.Getenv("log_saver_db_host")
	port := os.Getenv("log_saver_db_port")
	user := os.Getenv("log_saver_db_user")
	password := os.Getenv("log_saver_db_password")
	dbName := os.Getenv("log_saver_db_name")

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbName)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return &LogSaver{db}
}

func (saver LogSaver) saveGetRequest(request domain.GetOrderGatewayRequest) (int64, error) {
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
		return -1, err
	}

	return res.LastInsertId()
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
		return -1, err
	}

	return res.LastInsertId()
}

func (saver LogSaver) saveGetResponse(response domain.GetOrderGatewayResponse) (int64, error) {
	var bodyJson []byte
	var marshalErr error

	if !reflect.DeepEqual(response.Body, domain.GetOrderResponse{}) {
		bodyJson, marshalErr = json.Marshal(response.Body)
		if marshalErr != nil {
			log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
			return -1, marshalErr
		}
	}

	res, err := saver.connection.Exec(
		"INSERT INTO response(request_id, response_time, type, status, body, error_message) "+
			"values (?, ?, ?, ?, ?, ?)",
		response.RequestId, response.ResponseTime, response.Type, response.Status, string(bodyJson), response.ErrorMessage,
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return -1, err
	}

	return res.LastInsertId()
}

func (saver LogSaver) saveCreateResponse(response domain.CreateOrderGatewayResponse) (int64, error) {
	var bodyJson []byte
	var marshalErr error

	if !reflect.DeepEqual(response.Body, domain.CreateOrderResponse{}) {
		bodyJson, marshalErr = json.Marshal(response.Body)
		if marshalErr != nil {
			log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
			return -1, marshalErr
		}
	}

	res, err := saver.connection.Exec(
		"INSERT INTO response(request_id, response_time, type, status, body, error_message) "+
			"values (?, ?, ?, ?, ?, ?)",
		response.RequestId, response.ResponseTime, response.Type, response.Status, string(bodyJson), response.ErrorMessage,
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return -1, err
	}

	return res.LastInsertId()
}
