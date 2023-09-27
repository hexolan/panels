import { createSlice } from '@reduxjs/toolkit'

import { authApiSlice } from '../api/auth'
import { usersApiSlice } from '../api/users'
import type { User } from '../types/common'

export interface AuthState {
  accessToken: string | null;
  currentUser: User | null;
}

const initialState: AuthState = {
  accessToken: null,
  currentUser: null
}

export const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setUnauthed: state => {
      state.accessToken = null
      state.currentUser = null
    }
  },
  extraReducers: (builder) => {
    builder.addMatcher(
      authApiSlice.endpoints.login.matchFulfilled,
      (state, { payload }) => {
        state.accessToken = payload.token.access_token
        state.currentUser = payload.user
      }
    ).addMatcher(
      usersApiSlice.endpoints.registerUser.matchFulfilled,
      (state, { payload }) => {
        state.accessToken = payload.token.access_token
        state.currentUser = payload.user
      }
    )
  },
})

export const { setUnauthed } = authSlice.actions
export default authSlice.reducer