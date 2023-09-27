import { Empty } from "@bufbuild/protobuf"
import { ConnectRouter, ConnectError, Code } from "@connectrpc/connect"

import { createUser, getUserById, getUserByUsername, updateUserById, updateUserByUsername, deleteUserById, deleteUserByUsername } from "../service"
import { userToProtoUser } from "../proto/convert"
import { UserService } from "../proto/user_connect"
import { Health } from "../proto/grpc_health_connect"

import {
  HealthCheckRequest,
  HealthCheckResponse,
  HealthCheckResponse_ServingStatus
} from "../proto/grpc_health_pb"

import {
  CreateUserRequest,
  GetUserByIdRequest,
  GetUserByNameRequest,
  UpdateUserByIdRequest,
  UpdateUserByNameRequest,
  DeleteUserByIdRequest,
  DeleteUserByNameRequest,
  User as ProtoUser
} from "../proto/user_pb"

export default (router: ConnectRouter) => {
  router.service(UserService, {
    async createUser(req: CreateUserRequest): Promise<ProtoUser> {
      // validate inputs
      if (req.data === undefined) {
        throw new ConnectError("no values provided", Code.InvalidArgument);
      }

      if (req.data.username === undefined || req.data.username === "") {
        throw new ConnectError("no username provided", Code.InvalidArgument);
      }

      // attempt to create user
      const user = await createUser(req.data.username);
      return userToProtoUser(user);
    },

    async getUser(req: GetUserByIdRequest): Promise<ProtoUser> {
      const user = await getUserById(req.id);
      return userToProtoUser(user);
    },
    
    async getUserByName(req: GetUserByNameRequest): Promise<ProtoUser> {
      const user = await getUserByUsername(req.username);
      return userToProtoUser(user);
    },
    
    async updateUser(req: UpdateUserByIdRequest): Promise<ProtoUser> {
      // validate inputs
      if (req.id === "") {
        throw new ConnectError("no user id provided", Code.InvalidArgument);
      }

      if (req.data === undefined) {
        throw new ConnectError("no values provided", Code.InvalidArgument);
      }

      if (req.data.username === undefined) {
        throw new ConnectError("no username value provided", Code.InvalidArgument);
      }

      // attempt to update user
      const user = await updateUserById(req.id, req.data.username);
      return userToProtoUser(user);
    },
    
    async updateUserByName(req: UpdateUserByNameRequest): Promise<ProtoUser> {
      // validate inputs
      if (req.username === "") {
        throw new ConnectError("no username provided", Code.InvalidArgument);
      }

      if (req.data === undefined) {
        throw new ConnectError("no values provided", Code.InvalidArgument);
      }

      if (req.data.username === undefined) {
        throw new ConnectError("no username value provided", Code.InvalidArgument);
      }

      // attempt to update user
      const user = await updateUserByUsername(req.username, req.data.username);
      return userToProtoUser(user);
    },
    
    async deleteUser(req: DeleteUserByIdRequest): Promise<Empty> {
      // validate input
      if (req.id === "") {
        throw new ConnectError("no user id provided", Code.InvalidArgument);
      }

      // attempt to delete the user
      await deleteUserById(req.id);
      return new Empty();
    },
    
    async deleteUserByName(req: DeleteUserByNameRequest): Promise<Empty> {
      // validate input
      if (req.username === "") {
        throw new ConnectError("no username provided", Code.InvalidArgument);
      }

      // attempt to delete the user
      await deleteUserByUsername(req.username);
      return new Empty();
    }
  });

  // Health gRPC Service
  router.service(Health, {
    check(): HealthCheckResponse {
      const healthyResponse = new HealthCheckResponse({
        status: HealthCheckResponse_ServingStatus.SERVING,
      });
      return healthyResponse;
    },

    async *watch(req: HealthCheckRequest): AsyncGenerator<HealthCheckResponse> {
      const healthyResponse = new HealthCheckResponse({
        status: HealthCheckResponse_ServingStatus.SERVING,
      });
      
      while (req) {
        yield healthyResponse;
        await new Promise(resolve => setTimeout(resolve, 1000));
      }
    }
  });
}