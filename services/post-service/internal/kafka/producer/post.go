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

package producer

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	"github.com/hexolan/panels/post-service/internal"
	"github.com/hexolan/panels/post-service/internal/rpc/postv1"
)

type PostEventProducer struct {
	writer *kafka.Writer
}

func NewPostEventProducer(cfg internal.Config) PostEventProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.KafkaBrokers...),
		Topic:    "post",
		Balancer: &kafka.LeastBytes{},
	}

	return PostEventProducer{writer: writer}
}

func (ep PostEventProducer) SendEvent(event *postv1.PostEvent) {
	evtBytes, err := proto.Marshal(event)
	if err != nil {
		log.Panic().Err(err).Msg("failed to marshal event")
	}

	// Write to kafka
	err = ep.writer.WriteMessages(context.Background(), kafka.Message{Value: evtBytes})
	if err != nil {
		log.Panic().Err(err).Msg("failed to dispatch event")
	}
}

func (ep PostEventProducer) DispatchCreatedEvent(post *internal.Post) {
	go ep.SendEvent(&postv1.PostEvent{
		Type: "created",
		Data: postv1.PostToProto(post),
	})
}

func (ep PostEventProducer) DispatchUpdatedEvent(post *internal.Post) {
	go ep.SendEvent(&postv1.PostEvent{
		Type: "updated",
		Data: postv1.PostToProto(post),
	})
}

func (ep PostEventProducer) DispatchDeletedEvent(id internal.PostId) {
	go ep.SendEvent(&postv1.PostEvent{
		Type: "deleted",
		Data: &postv1.Post{Id: id.GetReprId()},
	})
}
