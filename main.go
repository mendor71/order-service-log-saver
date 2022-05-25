package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"order-service-log-saver/repository"
	requestUnmarshaler "order-service-log-saver/unmarshaler/request"
	responseUnmarshaler "order-service-log-saver/unmarshaler/response"
	"sync"
	"time"
)

type Buffer struct {
	container [][]byte
	size      uint8
}

func (bfr *Buffer) append(element []byte) {
	bfr.container = append(bfr.container, element)
	bfr.size++
}

func (bfr *Buffer) clean() {
	bfr.container = [][]byte{}
	bfr.size = 0
}

func readAndSaveRequests(
	reader *kafka.Reader,
	batchSize uint8,
	requestRepository repository.RequestRepository,
) {
	var buffer = Buffer{
		container: [][]byte{},
		size:      0,
	}
	var ctx = context.Background()
	for {
		message, err := reader.FetchMessage(ctx)
		if err != nil {
			log.Error().Msgf("message receive error: %s", err.Error())
			continue
		}

		buffer.append(message.Value)

		if buffer.size >= batchSize {
			get, create := requestUnmarshaler.UnmarshalRequests(buffer.container)

			for _, getElement := range get {
				err := requestRepository.SaveGet(getElement)
				if err != nil {
					log.Error().Msgf("get request save error: %s", err.Error())
					continue
				}
				log.Info().Msgf("Save get request: %s", getElement.RequestId)
			}

			for _, createElement := range create {
				err := requestRepository.SaveCreate(createElement)
				if err != nil {
					log.Error().Msgf("create request save error: %s", err.Error())
					continue
				}
				log.Info().Msgf("Save create request: %s", createElement.RequestId)
			}

			buffer.clean()

			if err := reader.CommitMessages(ctx, message); err != nil {
				log.Error().Msgf("failed to commit messages: %s", err.Error())
			}
		}
	}
}

func readAndSaveResponses(
	reader *kafka.Reader,
	batchSize uint8,
	responseRepository repository.ResponseRepository,
) {
	var buffer = Buffer{
		container: [][]byte{},
		size:      0,
	}
	var ctx = context.Background()
	for {
		message, err := reader.FetchMessage(ctx)
		if err != nil {
			log.Error().Msgf("message receive error: %s", err.Error())
			continue
		}

		buffer.append(message.Value)

		if buffer.size >= batchSize {
			get, create := responseUnmarshaler.UnmarshalResponses(buffer.container)

			for _, getElement := range get {
				err := responseRepository.SaveGet(getElement)
				if err != nil {
					log.Error().Msgf("get response save error: %s", err.Error())
					continue
				}
				log.Info().Msgf("Save get response: %s", getElement.RequestId)
			}

			for _, createElement := range create {
				err := responseRepository.SaveCreate(createElement)
				if err != nil {
					log.Error().Msgf("create response save error: %s", err.Error())
					continue
				}
				log.Info().Msgf("Save create response: %s", createElement.RequestId)
			}

			buffer.clean()

			if err := reader.CommitMessages(ctx, message); err != nil {
				log.Error().Msgf("failed to commit messages: %s", err.Error())
			}
		}
	}
}

func main() {
	brokers := []string{"127.0.0.1:9092"}
	groupId := "order-service-log-repository-1"

	requestTopicConfig := kafka.ReaderConfig{
		Topic:           "order-service-request-log",
		Brokers:         brokers,
		GroupID:         groupId,
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         10 * time.Second,
		ReadLagInterval: -1,
	}

	responseTopicConfig := kafka.ReaderConfig{
		Topic:           "order-service-response-log",
		Brokers:         brokers,
		GroupID:         groupId,
		MinBytes:        10e3,
		MaxBytes:        10e6,
		MaxWait:         10 * time.Second,
		ReadLagInterval: -1,
	}

	requestReader := kafka.NewReader(requestTopicConfig)
	responseReader := kafka.NewReader(responseTopicConfig)

	connectionPool := repository.NewDbConnectionPool()

	requestRepository := repository.RequestRepository{DbPool: connectionPool}
	responseRepository := repository.ResponseRepository{DbPool: connectionPool}

	var wg sync.WaitGroup
	wg.Add(1)

	go readAndSaveRequests(requestReader, 10, requestRepository)
	go readAndSaveResponses(responseReader, 10, responseRepository)

	wg.Wait()
}
