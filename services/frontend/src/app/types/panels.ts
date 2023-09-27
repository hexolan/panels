import { convertRawTimestamp } from './api';

import type { Panel } from './common';
import type { RawResponse, RawTimestamp } from './api';

// Request Data
export type CreatePanelData = {
  name: string;
  description: string;
}

export type UpdatePanelData = Partial<CreatePanelData>

// API Request Paramaters
type PanelByIdBase = {
  id: string;
}

type PanelByNameBase = {
  name: string;
}

export type GetPanelByIdRequest = PanelByIdBase;
export type GetPanelByNameRequest = PanelByNameBase;

export type UpdatePanelByIdRequest = PanelByIdBase & {
  data: UpdatePanelData;
}

export type UpdatePanelByNameRequest = PanelByNameBase & {
  data: UpdatePanelData;
}

export type DeletePanelByIdRequest = PanelByIdBase;
export type DeletePanelByNameRequest = PanelByNameBase;

export type CreatePanelRequest = CreatePanelData;

// API Responses
export type RawPanel = {
  id: string;
  name: string;
  description: string;
  created_at: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type RawPanelResponse = RawResponse & {
  data?: RawPanel;
}

// API Response Conversion
export const convertRawPanel = (rawPanel: RawPanel): Panel => ({
  id: rawPanel.id,
  name: rawPanel.name,
  description: rawPanel.description,
  createdAt: convertRawTimestamp(rawPanel.created_at),
  updatedAt: (rawPanel.updated_at ? convertRawTimestamp(rawPanel.updated_at) : undefined),
})