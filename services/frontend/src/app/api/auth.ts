import { apiSlice } from '../features/api'
import { convertRawAuthData } from '../types/auth'

import type { AuthData } from '../types/common'
import type { LoginRequest, RawAuthResponse } from '../types/auth'

export const authApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    login: builder.mutation<AuthData, LoginRequest>({
      query: data => ({
        url: '/v1/auth/login',
        method: 'POST',
        body: { ...data }
      }),
      transformResponse: (response: RawAuthResponse) => {
        if (response.data === undefined) { throw Error('invalid auth response') }

        return convertRawAuthData(response.data)
      },
    }),
  })
})

export const { useLoginMutation } = authApiSlice