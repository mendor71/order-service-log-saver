package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
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

func main() {
	brokers := []string{"127.0.0.1:9092"}
	groupId := "order-service-log-saver-1"

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

	var wg sync.WaitGroup
	wg.Add(1)

	go readAndSave(requestReader, 10)
	go readAndSave(responseReader, 10)

	wg.Wait()
}

func readAndSave(reader *kafka.Reader, batchSize uint8) {
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
			for _, msg := range buffer.container {
				log.Info().Msg(string(msg))
			}

			buffer.clean()

			if err := reader.CommitMessages(ctx, message); err != nil {
				log.Error().Msgf("failed to commit messages: %s", err.Error())
			}
		}
	}
}

//func flushMessages(buffer *Buffer) {
//	log.Info().Msgf("Flush %d messages", buffer.size)
//
//	requests, responses := Unmarshal(buffer.container)
//	saver := NewLogSaver()
//
//	for _, req := range requests {
//		saver.saveRequest(req)
//	}
//
//	for _, resp := range responses {
//		saver.saveResponse(resp)
//	}
//
//	buffer.clean()
//	log.Info().Msgf("Buffer clear, size: %d", buffer.size)
//}
