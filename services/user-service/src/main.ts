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