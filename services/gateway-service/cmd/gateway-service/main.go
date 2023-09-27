package main

import (
	"github.com/hexolan/panels/gateway-service/internal"
	"github.com/hexolan/panels/gateway-service/internal/api"
	"github.com/hexolan/panels/gateway-service/internal/rpc"
)

func main() {
	// Load the configuration
	cfg := internal.NewConfig()

	// Connect to the RPC services
	rpc.DialRPCServices(cfg)

	// Serve the api.
	app := api.NewAPIApp(cfg)
	api.ServeAPIApp(app)
}