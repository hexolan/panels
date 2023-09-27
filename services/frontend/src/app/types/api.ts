export type RawResponse = {
  status: string;
  msg?: string;
  data?: object;
}

export type ErrorResponse = RawResponse & {
  msg: string;
  data?: null;
}

export type RawTimestamp = {
  seconds: number;
  nanos?: number;
}

export const convertRawTimestamp = (timestamp: RawTimestamp): string => {
  return new Date(timestamp.seconds * 1000).toISOString()
}