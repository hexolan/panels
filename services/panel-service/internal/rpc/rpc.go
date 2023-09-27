package rpc

import (
	"net"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

	"github.com/hexolan/panels/panel-service/internal"
	"github.com/hexolan/panels/panel-service/internal/rpc/panelv1"
)

type RPCServer struct {
	grpcSvr *grpc.Server
}

func NewRPCServer(panelService internal.PanelService) *RPCServer {
	logger := log.Logger.With().Timestamp().Str("src", "rpc").Logger()

	svr := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(loggingInterceptor(logger)),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(loggingInterceptor(logger)),
		),
	)

	// Panels Service Server
	panelSvr := NewPanelServer(panelService)
	panelv1.RegisterPanelServiceServer(svr, &panelSvr)

	// Health Check Server
	healthSvr := health.NewServer()
	grpc_health_v1.RegisterHealthServer(svr, healthSvr)

	return &RPCServer{grpcSvr: svr}
}

func loggingInterceptor(logger zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		logger := logger.With().Fields(fields).Logger()

		switch lvl {
			case logging.LevelError:
				logger.Error().Msg(msg)
			case logging.LevelWarn:
				logger.Warn().Msg(msg)
			case logging.LevelInfo:
				logger.Info().Msg(msg)
			case logging.LevelDebug:
				logger.Debug().Msg(msg)
			default:
				logger.Debug().Interface("unknown-log-level", lvl).Msg(msg)
		}
	})
}

func (r *RPCServer) Serve() {
	// Prepare the listening port.
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Panic().Err(err).Caller().Msg("failed to listen on RPC port (:9090)")
	}
	
	// Begin serving gRPC.
	log.Info().Str("address", lis.Addr().String()).Msg("attempting to serve RPC...")
	err = r.grpcSvr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Caller().Msg("failed to serve RPC")
	}
}