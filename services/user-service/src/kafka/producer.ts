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

import { Kafka, Message, Producer } from "kafkajs";

import { kafkaBrokers } from "../config";
import { IUser } from "../mongo/User";
import { UserEvent } from "../proto/user_pb";
import { userToProtoUser } from "../proto/convert";

class UserProducer {
  private producer: Producer

  constructor() {
    const kafka = new Kafka({
      clientId: "user-service",
      brokers: kafkaBrokers,
    })

    this.producer = kafka.producer()
  }

  public async connect(): Promise<void> {
    try {
      await this.producer.connect()
    } catch (error) {
      console.log("error connecting to kafka producer ", error);
      throw error
    }
  }

  public async disconnect(): Promise<void> {
    await this.producer.disconnect()
  }

  public async sendCreatedEvent(user: IUser): Promise<void> {
    const event = new UserEvent({
      type: "created",
      data: userToProtoUser(user)
    })

    this.sendEvent(event);
  }

  public async sendUpdatedEvent(user: IUser): Promise<void> {
    const event = new UserEvent({
      type: "updated",
      data: userToProtoUser(user)
    })

    this.sendEvent(event);
  }

  public async sendDeletedEvent(user: IUser): Promise<void> {
    const event = new UserEvent({
      type: "deleted",
      data: userToProtoUser(user)
    })

    this.sendEvent(event);
  }

  private async sendEvent(event: UserEvent): Promise<void> {
    const msg: Message = {
      value: Buffer.from(event.toBinary())
    }
    
    await this.producer.send({
      topic: "user",
      messages: [msg]
    });
  }
}

const userProducer = new UserProducer();
export default userProducer;