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

package rpc

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/hexolan/panels/post-service/internal"
	"github.com/hexolan/panels/post-service/internal/rpc/postv1"
)

type RPCServer struct {
	grpcSvr *grpc.Server
}

func NewRPCServer(service internal.PostService) *RPCServer {
	logger := log.Logger.With().Timestamp().Str("src", "rpc").Logger()

	svr := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(loggingInterceptor(logger)),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(loggingInterceptor(logger)),
		),
	)

	// Post Service Server
	postSvr := NewPostServer(service)
	postv1.RegisterPostServiceServer(svr, &postSvr)

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
	// Prepare the listening port
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Panic().Err(err).Caller().Msg("failed to listen on RPC port (:9090)")
	}

	// Begin serving RPC
	log.Info().Str("address", lis.Addr().String()).Msg("Attempting to serve RPC...")
	err = r.grpcSvr.Serve(lis)
	if err != nil {
		log.Panic().Err(err).Caller().Msg("failed to serve RPC")
	}
}
