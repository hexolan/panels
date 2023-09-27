package kafka

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/hexolan/panels/panel-service/internal"
	"github.com/hexolan/panels/panel-service/internal/rpc/panelv1"
)

type PanelEventProducer struct {
	writer *kafka.Writer
}

func NewPanelEventProducer(cfg internal.Config) PanelEventProducer {
	writer := &kafka.Writer{
		Addr: kafka.TCP(cfg.KafkaBrokers...),
		Topic: "panel",
		Balancer: &kafka.LeastBytes{},
	}

	return PanelEventProducer{writer: writer}
}

func (ep PanelEventProducer) SendEvent(event *panelv1.PanelEvent) {
	// Encode the protobuf event
	evtBytes, err := proto.Marshal(event)
	if err != nil {
		log.Panic().Err(err).Msg("failed to marshal event")
	}

	// Write to kafka
	err = ep.writer.WriteMessages(context.Background(), kafka.Message{Value: evtBytes})
	if err != nil {
		// todo: implement recovery method e.g. storing failed event dispatches on DB to send on recovery (such as from Kafka going offline)
		log.Panic().Err(err).Msg("failed to dispatch event")
	}
}

func (ep PanelEventProducer) DispatchCreatedEvent(panel internal.Panel) {
	go ep.SendEvent(&panelv1.PanelEvent{
		Type: "created",
		Data: panelv1.PanelToProto(&panel),
	})
}

func (ep PanelEventProducer) DispatchUpdatedEvent(panel internal.Panel) {
	go ep.SendEvent(&panelv1.PanelEvent{
		Type: "updated",
		Data: panelv1.PanelToProto(&panel),
	})
}

func (ep PanelEventProducer) DispatchDeletedEvent(id int64) {
	go ep.SendEvent(&panelv1.PanelEvent{
		Type: "deleted",
		Data: &panelv1.Panel{Id: internal.StringifyPanelId(id)},
	})
}