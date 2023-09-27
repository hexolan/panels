package consumers

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/hexolan/panels/post-service/internal"
	"github.com/hexolan/panels/post-service/internal/kafka/producer"
	"github.com/hexolan/panels/post-service/internal/kafka/consumers/panelv1"
)

type PanelEventConsumer struct {
	reader *kafka.Reader

	dbRepo internal.PostDBRepository
	eventProd producer.PostEventProducer
}

func NewPanelEventConsumer(cfg internal.Config, dbRepo internal.PostDBRepository, eventProd producer.PostEventProducer) PanelEventConsumer {
	return PanelEventConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: cfg.KafkaBrokers,
			GroupID: "post-service",
			Topic: "panel",
		}),
		dbRepo: dbRepo,
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