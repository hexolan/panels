// Copyright 2023 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package consumers

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/hexolan/panels/post-service/internal"
	"github.com/hexolan/panels/post-service/internal/kafka/consumers/panelv1"
	"github.com/hexolan/panels/post-service/internal/kafka/producer"
)

type PanelEventConsumer struct {
	reader *kafka.Reader

	dbRepo    internal.PostDBRepository
	eventProd producer.PostEventProducer
}

func NewPanelEventConsumer(cfg internal.Config, dbRepo internal.PostDBRepository, eventProd producer.PostEventProducer) PanelEventConsumer {
	return PanelEventConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: cfg.KafkaBrokers,
			GroupID: "post-service",
			Topic:   "panel",
		}),
		dbRepo:    dbRepo,
		eventProd: eventProd,
	}
}

func (ec PanelEventConsumer) ProcessEvent(evt *panelv1.PanelEvent) {
	if evt.GetData() == nil {
		log.Error().Str("src", "panel-event-consumer").Any("event", evt).Msg("no event data provided")
		return
	}

	if evt.Type == "deleted" {
		postIds, err := ec.dbRepo.DeletePostsOnPanel(context.Background(), evt.GetData().GetId())
		if err == nil {
			for _, postId := range postIds {
				ec.eventProd.DispatchDeletedEvent(postId)
			}
		}
		log.Debug().Str("src", "panel-event-consumer").Any("event", evt).Msg("processed panel deleted event")
	}
}

func (ec PanelEventConsumer) ProcessMessage(msg kafka.Message) {
	var evt panelv1.PanelEvent
	err := proto.Unmarshal(msg.Value, &evt)
	if err != nil {
		log.Error().Err(err).Str("src", "panel-event-consumer").Msg("failed to unmarshal event")
		return
	}

	ec.ProcessEvent(&evt)
}

func (ec PanelEventConsumer) Start() {
	for {
		msg, err := ec.reader.FetchMessage(context.Background())
		if err != nil {
			log.Error().Err(err).Str("src", "panel-event-consumer").Msg("failed to fetch msg from Kafka")
			break
		}
		ec.ProcessMessage(msg)
	}

	// Attempt to close the reader connection (after experiencing a Kafka error)
	if err := ec.reader.Close(); err != nil {
		log.Panic().Err(err).Str("src", "panel-event-consumer").Msg("failed to close Kafka reader")
	}
}
