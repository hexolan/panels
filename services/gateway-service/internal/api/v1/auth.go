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

package v1

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/rpc"
	"github.com/hexolan/panels/gateway-service/internal/rpc/authv1"
)

type userLoginForm struct {
	Username string
	Password string
}

func setAuthMethod(userId string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err := rpc.Svcs.GetAuthSvc().SetPasswordAuth(
		ctx,
		&authv1.SetPasswordAuthMethod{
			UserId:   userId,
			Password: password,
		},
	)
	return err
}

func authWithPassword(userId string, password string) (*authv1.AuthToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	token, err := rpc.Svcs.GetAuthSvc().AuthWithPassword(
		ctx,
		&authv1.PasswordAuthRequest{
			UserId:   userId,
			Password: password,
		},
	)
	return token, err
}

func LoginWithPassword(c *fiber.Ctx) error {
	// Parse the body data
	form := new(userLoginForm)
	if err := c.BodyParser(form); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	// username -> user ID
	user, err := getUserByUsername(form.Username)
	if err != nil {
		return err
	}

	// attempt to auth
	token, err := authWithPassword(user.Id, form.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user":  user,
			"token": token,
		},
	})
}
