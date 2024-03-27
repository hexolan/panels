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

package handlers

import (
	"crypto/rsa"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/golang-jwt/jwt/v5"

	"github.com/hexolan/panels/gateway-service/internal"
)

var AuthMiddleware fiber.Handler

type TokenClaims struct {
	jwt.RegisteredClaims
}

type tokenValidator struct {
	pubKey *rsa.PublicKey
}

func NewAuthMiddleware(cfg internal.Config) {
	tokenValidator := tokenValidator{pubKey: cfg.JWTPubKey}
	AuthMiddleware = keyauth.New(keyauth.Config{
		AuthScheme: "Bearer",
		Validator:  tokenValidator.ValidateToken,
	})
}

func GetTokenClaims(c *fiber.Ctx) (TokenClaims, error) {
	var tokenClaims TokenClaims
	tokenClaims, ok := c.Locals("tokenClaims").(TokenClaims)
	if !ok {
		return TokenClaims{}, fiber.NewError(fiber.StatusUnauthorized, "unable to access token claims")
	}
	return tokenClaims, nil
}

func (tv tokenValidator) validateToken(token *jwt.Token) (interface{}, error) {
	// Ensure token is signed with RSA
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, keyauth.ErrMissingOrMalformedAPIKey
	}

	// Validate token with public key
	return tv.pubKey, nil
}

func (tv tokenValidator) ValidateToken(c *fiber.Ctx, userToken string) (bool, error) {
	claims := TokenClaims{}
	_, err := jwt.ParseWithClaims(userToken, &claims, tv.validateToken)
	if err != nil {
		return false, err
	}

	c.Locals("tokenClaims", claims)
	return true, nil
}
