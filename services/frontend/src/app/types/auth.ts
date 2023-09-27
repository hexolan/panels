import { convertRawUser } from './user'

import type { RawUser } from './user'
import type { RawResponse } from './api'
import type { AuthData, AuthToken } from './common'

// API Request Paramaters
export type LoginRequest = {
  username: string;
  password: string;
}

// API Responses
type RawAuthData = {
  token: AuthToken,
  user: RawUser
}

export type RawAuthResponse = RawResponse & {
  data?: RawAuthData;
}

// API Response Conversion
export const convertRawAuthData = (data: RawAuthData): AuthData => ({
  token: data.token,
  user: convertRawUser(data.user)
})