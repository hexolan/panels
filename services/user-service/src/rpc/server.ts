import { fastify } from "fastify";
import { fastifyConnectPlugin } from "@connectrpc/connect-fastify";

import routes from "./connect";

export default async function serveRPC(): Promise<void> {
  const server = fastify({ http2: true, logger: true });
  await server.register(fastifyConnectPlugin, { routes: routes });

  await server.listen({ host: "0.0.0.0", port: 9090 });
  console.log("rpc server is listening at", server.addresses());
}