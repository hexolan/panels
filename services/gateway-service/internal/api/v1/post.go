package v1

import (
	"time"
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/rpc"
	"github.com/hexolan/panels/gateway-service/internal/rpc/postv1"
	"github.com/hexolan/panels/gateway-service/internal/api/handlers"
)

func getPostById(postId string) (*postv1.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return rpc.Svcs.GetPostSvc().GetPost(
		ctx,
		&postv1.GetPostRequest{Id: postId},
	)
}

func GetFeedPosts(c *fiber.Ctx) error {
	// Make the request for feed posts
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	posts, err := rpc.Svcs.GetPostSvc().GetFeedPosts(
		ctx,
		&postv1.GetFeedPostsRequest{},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": posts})
}

func GetPanelPostFromId(c *fiber.Ctx) error {
	// Make the request for the panel post
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	post, err := rpc.Svcs.GetPostSvc().GetPanelPost(
		ctx,
		&postv1.GetPanelPostRequest{Id: c.Params("id"), PanelId: c.Params("panel_id")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": post})
}

func GetPanelPostFromName(c *fiber.Ctx) error {
	// Get the panel ID from name.
	panelId, err := getPanelIDFromName(c.Params("panel_name"))
	if err != nil {
		return err
	}

	// Make the request for the panel post
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	post, err := rpc.Svcs.GetPostSvc().GetPanelPost(
		ctx,
		&postv1.GetPanelPostRequest{Id: c.Params("id"), PanelId: panelId},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": post})
}

func GetUserPostsFromId(c *fiber.Ctx) error {
	// Make the request for user posts
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	posts, err := rpc.Svcs.GetPostSvc().GetUserPosts(
		ctx,
		&postv1.GetUserPostsRequest{UserId: c.Params("user_id")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": posts})
}

func GetUserPostsFromUsername(c *fiber.Ctx) error {
	// Get the user ID from username.
	user, err := getUserByUsername(c.Params("username"))
	if err != nil {
		return err
	}

	// Make the request for user posts
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	posts, err := rpc.Svcs.GetPostSvc().GetUserPosts(
		ctx,
		&postv1.GetUserPostsRequest{UserId: user.Id},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": posts})
}

func GetPanelPostsFromId(c *fiber.Ctx) error {
	// Make the request for panel posts
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	posts, err := rpc.Svcs.GetPostSvc().GetPanelPosts(
		ctx,
		&postv1.GetPanelPostsRequest{PanelId: c.Params("panel_id")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": posts})
}

func GetPanelPostsFromName(c *fiber.Ctx) error {
	// Get the panel ID from name.
	panelId, err := getPanelIDFromName(c.Params("panel_name"))
	if err != nil {
		return err
	}

	// Make the request for panel posts
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	posts, err := rpc.Svcs.GetPostSvc().GetPanelPosts(
		ctx,
		&postv1.GetPanelPostsRequest{PanelId: panelId},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": posts})
}

func UpdatePost(c *fiber.Ctx) error {
	// check if user has permissions to update the post
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	post, err := getPostById(c.Params("id"))
	if err != nil {
		return err
	}

	if (post.AuthorId != currentUser.Id) {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to update that post")
	}

	// form patch data for update request
	patchData := new(postv1.PostMutable)
	if err := c.BodyParser(patchData); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	// update the post
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	post, err = rpc.Svcs.GetPostSvc().UpdatePost(
		ctx,
		&postv1.UpdatePostRequest{Id: c.Params("id"), Data: patchData},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": post})
}

func DeletePost(c *fiber.Ctx) error {
	// check if user has permissions to delete the post
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	post, err := getPostById(c.Params("id"))
	if err != nil {
		return err
	}

	if (post.AuthorId != currentUser.Id && !currentUser.IsAdmin) {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to delete that post")
	}
	
	// delete the post
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = rpc.Svcs.GetPostSvc().DeletePost(
		ctx,
		&postv1.DeletePostRequest{Id: c.Params("id")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "post deleted"})
}

func CreatePanelPostFromId(c *fiber.Ctx) error {
	// Parse the body data
	newPost := new(postv1.PostMutable)
	if err := c.BodyParser(newPost); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	// Get the panel ID from provided panel name
	panel, err := getPanelById(c.Params("panel_id"))
	if err != nil {
		return err
	}

	// access token claims
	tokenClaims, err := handlers.GetTokenClaims(c)
	if err != nil {
		return err
	}

	// make rpc call
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	post, err := rpc.Svcs.GetPostSvc().CreatePost(
		ctx,
		&postv1.CreatePostRequest{
			PanelId: panel.Id,
			UserId: tokenClaims.Subject,
			Data: newPost,
		},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": post})
}

func CreatePanelPostFromName(c *fiber.Ctx) error {
	// Parse the body data
	newPost := new(postv1.PostMutable)
	if err := c.BodyParser(newPost); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	// Get the panel ID from provided panel name
	panelId, err := getPanelIDFromName(c.Params("panel_name"))
	if err != nil {
		return err
	}

	// access token claims
	tokenClaims, err := handlers.GetTokenClaims(c)
	if err != nil {
		return err
	}

	// make rpc call
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	post, err := rpc.Svcs.GetPostSvc().CreatePost(
		ctx,
		&postv1.CreatePostRequest{
			PanelId: panelId,
			UserId: tokenClaims.Subject,
			Data: newPost,
		},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": post})
}