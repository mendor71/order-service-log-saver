package repository

import (
	"database/sql"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"order-service-log-saver/domain"
	"strconv"
)

type RequestRepository struct {
	DbPool *sql.DB
}

func (saver RequestRepository) SaveGet(request domain.GetOrderGatewayRequest) error {
	bodyJson, marshalErr := json.Marshal(request.Body)
	if marshalErr != nil {
		log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
		return marshalErr
	}
	bodyData := string(bodyJson)
	_, err := saver.DbPool.Exec(
		"INSERT INTO request(request_id, request_time, type, body) values ($1, $2, $3, $4)",
		request.RequestId,
		strconv.FormatFloat(request.RequestTime, 'E', -1, 64),
		request.Type,
		bodyData,
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return err
	}
	return nil
}

func (saver RequestRepository) SaveCreate(request domain.CreateOrderGatewayRequest) error {
	bodyJson, marshalErr := json.Marshal(request.Body)
	if marshalErr != nil {
		log.Error().Msgf("Can't marshal body to json: %s", marshalErr.Error())
		return marshalErr
	}
	bodyData := string(bodyJson)
	_, err := saver.DbPool.Exec(
		"INSERT INTO request(request_id, request_time, type, body) values ($1, $2, $3, $4)",
		request.RequestId,
		strconv.FormatFloat(request.RequestTime, 'E', -1, 64),
		request.Type,
		bodyData,
	)

	if err != nil {
		log.Error().Msgf("Can't insert row: %s", err.Error())
		return err
	}
	return nil
}
