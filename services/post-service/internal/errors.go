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

package internal

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewServiceError(code ErrorCode, msg string) error {
	return &ServiceError{
		code: code,
		msg:  msg,
	}
}

func NewServiceErrorf(code ErrorCode, msg string, args ...interface{}) error {
	return NewServiceError(code, fmt.Sprintf(msg, args...))
}

func WrapServiceError(original_err error, code ErrorCode, msg string) error {
	return &ServiceError{
		code:         code,
		msg:          msg,
		original_err: original_err,
	}
}

type ErrorCode int32

const (
	UnknownErrorCode ErrorCode = iota
	NotFoundErrorCode
	ConflictErrorCode
	ForbiddenErrorCode
	InvalidArgumentErrorCode
	ConnectionErrorCode
)

func (c ErrorCode) GRPCCode() codes.Code {
	codeMap := map[ErrorCode]codes.Code{
		UnknownErrorCode:         codes.Unknown,
		NotFoundErrorCode:        codes.NotFound,
		ConflictErrorCode:        codes.AlreadyExists,
		ForbiddenErrorCode:       codes.PermissionDenied,
		InvalidArgumentErrorCode: codes.InvalidArgument,
		ConnectionErrorCode:      codes.Unavailable,
	}

	grpcCode, mapped := codeMap[c]
	if mapped {
		return grpcCode
	}
	return codes.Unknown
}

type ServiceError struct {
	code         ErrorCode
	msg          string
	original_err error
}

func (e ServiceError) Error() string {
	if e.original_err != nil {
		return fmt.Sprintf("%s: %s", e.msg, e.original_err.Error())
	}
	return e.msg
}

func (e ServiceError) Code() ErrorCode {
	return e.code
}

func (e ServiceError) GRPCStatus() *status.Status {
	return status.New(e.Code().GRPCCode(), e.msg)
}

func (e ServiceError) Unwrap() error {
	return e.original_err
}
