import { apiSlice } from '../features/api'
import { convertRawPanel } from '../types/panels'

import type { Panel } from '../types/common'
import type {
  RawPanelResponse,
  GetPanelByIdRequest, GetPanelByNameRequest,
  UpdatePanelByIdRequest, UpdatePanelByNameRequest,
  DeletePanelByIdRequest, DeletePanelByNameRequest,
  CreatePanelRequest
} from '../types/panels'

export const panelsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPanelById: builder.query<Panel, GetPanelByIdRequest>({
      query: req => ({ url: `/v1/panels/id/${req.id}` }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return convertRawPanel(response.data)
      }
    }),

    getPanelByName: builder.query<Panel, GetPanelByNameRequest>({
      query: req => ({ url: `/v1/panels/name/${req.name}` }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return convertRawPanel(response.data)
      }
    }),

    updatePanelById: builder.mutation<Panel, UpdatePanelByIdRequest>({
      query: req => ({
        url: `/v1/panels/id/${req.id}`,
        method: 'PATCH',
        body: { ...req.data }
      }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return convertRawPanel(response.data)
      }
    }),

    updatePanelByName: builder.mutation<Panel, UpdatePanelByNameRequest>({
      query: req => ({
        url: `/v1/panels/name/${req.name}`,
        method: 'PATCH',
        body: { ...req.data }
      }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return convertRawPanel(response.data)
      }
    }),

    deletePanelById: builder.mutation<void, DeletePanelByIdRequest>({
      query: req => ({
        url: `/v1/panels/id/${req.id}`,
        method: 'DELETE'
      })
    }),

    deletePanelByName: builder.mutation<void, DeletePanelByNameRequest>({
      query: req => ({
        url: `/v1/panels/id/${req.name}`,
        method: 'DELETE'
      })
    }),

    createPanel: builder.mutation<Panel, CreatePanelRequest>({
      query: req => ({
        url: '/v1/panels',
        method: 'POST',
        body: { ...req }
      }),
      transformResponse: (response: RawPanelResponse) => {
        if (response.data === undefined) { throw Error('invalid panel response') }

        return convertRawPanel(response.data)
      }
    }),
  })
})

export const { 
  useGetPanelByIdQuery, useGetPanelByNameQuery,
  useUpdatePanelByIdMutation, useUpdatePanelByNameMutation,
  useDeletePanelByIdMutation, useDeletePanelByNameMutation,
  useCreatePanelMutation 
} = panelsApiSlice