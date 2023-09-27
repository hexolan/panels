package v1

import (
	"time"
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/rpc"
	"github.com/hexolan/panels/gateway-service/internal/rpc/userv1"
	"github.com/hexolan/panels/gateway-service/internal/api/handlers"
)

type userSignupForm struct {
	Username string
	Password string
}

func getCurrentUser(c *fiber.Ctx) (*userv1.User, error) {
	// access token claims
	tokenClaims, err := handlers.GetTokenClaims(c)
	if err != nil {
		return nil, err
	}

	// fetch current user
	return getUserById(tokenClaims.Subject)
}

func getUserById(id string) (*userv1.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	user, err := rpc.Svcs.GetUserSvc().GetUser(
		ctx,
		&userv1.GetUserByIdRequest{Id: id},
	)

	return user, err
}

func getUserByUsername(username string) (*userv1.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	user, err := rpc.Svcs.GetUserSvc().GetUserByName(
		ctx,
		&userv1.GetUserByNameRequest{Username: username},
	)

	return user, err
}

func GetUserById(c *fiber.Ctx) error {
	user, err := getUserById(c.Params("id"))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

func GetUserByUsername(c *fiber.Ctx) error {
	user, err := getUserByUsername(c.Params("username"))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

func GetCurrentUser(c *fiber.Ctx) error {
	user, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

func DeleteUserById(c *fiber.Ctx) error {
	// get current user info
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	// check current user has permission to delete user
	if currentUser.Id != c.Params("id") && !currentUser.IsAdmin {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to delete that user")
	}
	
	// delete the user
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = rpc.Svcs.GetUserSvc().DeleteUser(
		ctx,
		&userv1.DeleteUserByIdRequest{Id: c.Params("id")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "user deleted"})
}

func DeleteUserByUsername(c *fiber.Ctx) error {
	// get current user info
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	// check current user has permission to delete user
	if currentUser.Id != c.Params("id") && !currentUser.IsAdmin {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to delete that user")
	}
	
	// delete the user
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = rpc.Svcs.GetUserSvc().DeleteUserByName(
		ctx,
		&userv1.DeleteUserByNameRequest{Username: c.Params("username")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "user deleted"})
}

func DeleteCurrentUser(c *fiber.Ctx) error {
	// View access token claims
	tokenClaims, err := handlers.GetTokenClaims(c)
	if err != nil {
		return err
	}

	// RPC call
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = rpc.Svcs.GetUserSvc().DeleteUser(
		ctx,
		&userv1.DeleteUserByIdRequest{Id: tokenClaims.Subject},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "user deleted"})
}

func UserSignup(c *fiber.Ctx) error {
	// Parse the body data
	form := new(userSignupForm)
	if err := c.BodyParser(form); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}
	
	// Attempt to create the user
	// todo: defer this logic away from gateway-service in future (potentially into seperate registration-service)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	user, err := rpc.Svcs.GetUserSvc().CreateUser(
		ctx,
		&userv1.CreateUserRequest{
			Data: &userv1.UserMutable{
				Username: &form.Username,
			},
		},
	)
	if err != nil {
		return err
	}

	// Attempt to set password auth method for the user
	err = setAuthMethod(user.Id, form.Password)
	if err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		_, _ = rpc.Svcs.GetUserSvc().DeleteUser(
			ctx,
			&userv1.DeleteUserByIdRequest{Id: user.Id},
		)
		return err
	}

	// Signup success - attempt to get auth token
	token, _ := authWithPassword(user.Id, form.Password)
	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": user,
			"token": token,
		},
	})
}