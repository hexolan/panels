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
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hexolan/panels/gateway-service/internal"
	"github.com/hexolan/panels/gateway-service/internal/rpc/authv1"
	"github.com/hexolan/panels/gateway-service/internal/rpc/commentv1"
	"github.com/hexolan/panels/gateway-service/internal/rpc/panelv1"
	"github.com/hexolan/panels/gateway-service/internal/rpc/postv1"
	"github.com/hexolan/panels/gateway-service/internal/rpc/userv1"
)

var Svcs RPCServices

type RPCServices struct {
	panelSvcConn   *grpc.ClientConn
	postSvcConn    *grpc.ClientConn
	userSvcConn    *grpc.ClientConn
	authSvcConn    *grpc.ClientConn
	commentSvcConn *grpc.ClientConn
}

func (rpcSvcs RPCServices) GetPanelSvc() panelv1.PanelServiceClient {
	return panelv1.NewPanelServiceClient(rpcSvcs.panelSvcConn)
}

func (rpcSvcs RPCServices) GetPostSvc() postv1.PostServiceClient {
	return postv1.NewPostServiceClient(rpcSvcs.postSvcConn)
}

func (rpcSvcs RPCServices) GetUserSvc() userv1.UserServiceClient {
	return userv1.NewUserServiceClient(rpcSvcs.userSvcConn)
}

func (rpcSvcs RPCServices) GetAuthSvc() authv1.AuthServiceClient {
	return authv1.NewAuthServiceClient(rpcSvcs.authSvcConn)
}

func (rpcSvcs RPCServices) GetCommentSvc() commentv1.CommentServiceClient {
	return commentv1.NewCommentServiceClient(rpcSvcs.commentSvcConn)
}

func DialRPCServices(cfg internal.Config) {
	Svcs = RPCServices{
		panelSvcConn:   dialRPC(cfg.PanelSvcAddr),
		postSvcConn:    dialRPC(cfg.PostSvcAddr),
		userSvcConn:    dialRPC(cfg.UserSvcAddr),
		authSvcConn:    dialRPC(cfg.AuthSvcAddr),
		commentSvcConn: dialRPC(cfg.CommentSvcAddr),
	}
}

func dialRPC(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to rpc: %s", addr))
	}

	return conn
}
