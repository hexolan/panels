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