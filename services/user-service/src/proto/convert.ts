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

import { Timestamp } from "@bufbuild/protobuf";

import { User as ProtoUser } from "../proto/user_pb";
import { IUser } from "../mongo/User";

function timeToProtoTimestamp(time: Date | undefined): Timestamp | undefined {
  if (!time) return undefined;
  return Timestamp.fromDate(time);
}

export function userToProtoUser(user: IUser): ProtoUser {
  return new ProtoUser({
    id: user._id.toString(),
    username: user.username,
    isAdmin: (user.isAdmin ? user.isAdmin : false),
    createdAt: timeToProtoTimestamp(user.createdAt),
    updatedAt: timeToProtoTimestamp(user.updatedAt)
  });
}