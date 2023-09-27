package kafka

import (
	"github.com/hexolan/panels/post-service/internal"
	"github.com/hexolan/panels/post-service/internal/kafka/producer"
	"github.com/hexolan/panels/post-service/internal/kafka/consumers"
)

type eventConsumers struct {
	userEC consumers.UserEventConsumer
	panelEC consumers.PanelEventConsumer
}

func NewEventConsumers(cfg internal.Config, dbRepo internal.PostDBRepository, eventProd producer.PostEventProducer) eventConsumers {
	return eventConsumers{
		userEC: consumers.NewUserEventConsumer(cfg, dbRepo, eventProd),
		panelEC: consumers.NewPanelEventConsumer(cfg, dbRepo, eventProd),
	}
}

func (ecs eventConsumers) Start() {
	go ecs.panelEC.Start()
}