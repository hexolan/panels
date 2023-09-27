import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import type { BaseQueryFn } from '@reduxjs/toolkit/query'

import { setUnauthed } from './auth'
import type { RootState } from '../store'
import type { QueryError } from '../types/api';

const baseQuery = fetchBaseQuery({
  baseUrl: import.meta.env.VITE_API_URL,
  prepareHeaders: (headers, { getState }) => {
    const state = getState() as RootState

    const token = state.auth.accessToken
    if (token) {
      headers.set('Authorization', `Bearer ${token}`)
    }

    return headers
  }
})

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const wrappedBaseQuery: BaseQueryFn<any, unknown, QueryError> = async (args, api, extraOptions) => {
  const result = await baseQuery(args, api, extraOptions)
  if ((api.getState() as RootState).auth.accessToken && result?.error?.status === 403) {
    api.dispatch(setUnauthed())
  }

  return result
}

export const apiSlice = createApi({
  baseQuery: wrappedBaseQuery,
  endpoints: () => ({}),
})