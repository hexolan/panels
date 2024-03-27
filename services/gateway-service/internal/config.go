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
	"crypto/rsa"
	"encoding/base64"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func NewConfig() Config {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(optFromEnvRequire("JWT_PUBLIC_KEY"))
	if err != nil {
		log.Fatal("jwt public key not in base64 format")
	}

	jwtPubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		log.Panic("invalid jwt public key provided (must be RSA)")
	}

	return Config{
		PanelSvcAddr:   optFromEnvRequire("PANEL_SVC_ADDR"),
		PostSvcAddr:    optFromEnvRequire("POST_SVC_ADDR"),
		UserSvcAddr:    optFromEnvRequire("USER_SVC_ADDR"),
		AuthSvcAddr:    optFromEnvRequire("AUTH_SVC_ADDR"),
		CommentSvcAddr: optFromEnvRequire("COMMENT_SVC_ADDR"),

		JWTPubKey: jwtPubKey,
	}
}

func optFromEnvRequire(opt string) string {
	optValue, exists := os.LookupEnv(opt)
	if !exists || optValue == "" {
		log.Fatalf("failed to load required config option ('%s')", opt)
	}
	return optValue
}

type Config struct {
	PostSvcAddr    string
	PanelSvcAddr   string
	UserSvcAddr    string
	AuthSvcAddr    string
	CommentSvcAddr string

	JWTPubKey *rsa.PublicKey
}
