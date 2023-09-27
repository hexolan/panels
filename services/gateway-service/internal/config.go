package internal

import (
	"os"
	"crypto/rsa"
	"encoding/base64"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2/log"
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
		PanelSvcAddr: optFromEnvRequire("PANEL_SVC_ADDR"),
		PostSvcAddr: optFromEnvRequire("POST_SVC_ADDR"),
		UserSvcAddr: optFromEnvRequire("USER_SVC_ADDR"),
		AuthSvcAddr: optFromEnvRequire("AUTH_SVC_ADDR"),
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
	PostSvcAddr string
	PanelSvcAddr string
	UserSvcAddr string
	AuthSvcAddr string
	CommentSvcAddr string

	JWTPubKey *rsa.PublicKey
}