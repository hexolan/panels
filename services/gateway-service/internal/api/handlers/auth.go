package handlers

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"

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
		Validator: tokenValidator.ValidateToken,
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