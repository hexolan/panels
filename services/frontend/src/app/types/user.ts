import { convertRawTimestamp } from './api';

import type { User } from './common';
import type { RawResponse, RawTimestamp } from './api';

// Request Data
type RegisterUserData = {
  username: string;
  password: string;
}

// API Request Paramaters
type UserByIdBase = {
  id: string;
}

type UserByNameBase = {
  username: string;
}

export type GetUserByIdRequest = UserByIdBase
export type GetUserByNameRequest = UserByNameBase

export type DeleteUserByIdRequest = UserByIdBase
export type DeleteUserByNameRequest = UserByNameBase

export type RegisterUserRequest = RegisterUserData

// API Responses
export type RawUser = {
  id: string;
  username: string;
  is_admin?: boolean;
  created_at?: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type RawUserResponse = RawResponse & {
  data?: RawUser;
}

// API Response Conversion
export const convertRawUser = (rawUser: RawUser): User => ({
  id: rawUser.id,
  username: rawUser.username,
  isAdmin: (rawUser.is_admin ? rawUser.is_admin : false),
  createdAt: (rawUser.created_at ? convertRawTimestamp(rawUser.created_at) : undefined),
  updatedAt: (rawUser.updated_at ? convertRawTimestamp(rawUser.updated_at) : undefined),
})