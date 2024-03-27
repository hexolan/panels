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

function getMongoURI(): string {
  if (process.env.MONGODB_URI === undefined) {
    throw new Error("mongodb_uri configuration not provided");
  } else {
    return process.env.MONGODB_URI
  }
}

function getKafkaBrokers(): string[] {
  if (process.env.KAFKA_BROKERS === undefined) {
    throw new Error("kafka_brokers configuration not provided");
  } else {
    return process.env.KAFKA_BROKERS.split(",")
  }
}

const mongodbUri = getMongoURI();
const kafkaBrokers = getKafkaBrokers();

export { mongodbUri, kafkaBrokers }