package repository

import (
	"database/sql"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"order-service-log-saver/domain"
	"reflect"
	"strconv"
)

type ResponseRepository struct {
	DbPool *sql.DB
}

func (saver ResponseRepository) SaveGet(response domain.GetOrderGatewayResponse) error {
	var bodyJson []byte
	var marshalErr error

	if !reflect.DeepEqual(response.Body, domain.GetOrderResponse{}) {
		bodyJson, marshalErr = json.Marshal(response.Body)
		if marshalErr != nil {
			log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
			return marshalErr
		}
	}

	bodyData := string(bodyJson)
	_, err := saver.DbPool.Exec(
		"INSERT INTO response(request_id, response_time, type, status, body, error_message) "+
			"values ($1, $2, $3, $4, $5, $6)",
		response.RequestId,
		strconv.FormatFloat(response.ResponseTime, 'E', -1, 64),
		response.Type,
		response.Status,
		bodyData,
		response.ErrorMessage,
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return err
	}

	return nil
}

func (saver ResponseRepository) SaveCreate(response domain.CreateOrderGatewayResponse) error {
	var bodyJson []byte
	var marshalErr error

	if !reflect.DeepEqual(response.Body, domain.CreateOrderResponse{}) {
		bodyJson, marshalErr = json.Marshal(response.Body)
		if marshalErr != nil {
			log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
			return marshalErr
		}
	}

	bodyData := string(bodyJson)
	_, err := saver.DbPool.Exec(
		"INSERT INTO response(request_id, response_time, type, status, body, error_message) "+
			"values ($1, $2, $3, $4, $5, $6)",
		response.RequestId,
		strconv.FormatFloat(response.ResponseTime, 'E', -1, 64),
		response.Type,
		response.Status,
		bodyData,
		response.ErrorMessage,
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return err
	}

	return nil
}
