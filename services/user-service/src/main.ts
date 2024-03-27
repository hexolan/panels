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

import mongoose from "mongoose";

import { mongodbUri } from "./config";
import userProducer from "./kafka/producer";
import serveRPC from "./rpc/server";

async function main(): Promise<void> {
  await mongoose.connect(mongodbUri);
  await userProducer.connect()

  serveRPC();
}

main();