package v1

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/api/handlers"
	"github.com/hexolan/panels/gateway-service/internal/rpc"
	"github.com/hexolan/panels/gateway-service/internal/rpc/commentv1"
)

func getComment(id string) (*commentv1.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	comment, err := rpc.Svcs.GetCommentSvc().GetComment(
		ctx,
		&commentv1.GetCommentRequest{Id: id},
	)

	return comment, err
}

func GetPostComments(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	comments, err := rpc.Svcs.GetCommentSvc().GetPostComments(
		ctx,
		&commentv1.GetPostCommentsRequest{PostId: c.Params("post_id")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": comments})
}

func UpdateComment(c *fiber.Ctx) error {
	// check if user has permissions to update the comment
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	comment, err := getComment(c.Params("id"))
	if err != nil {
		return err
	}

	if (comment.AuthorId != currentUser.Id) {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to update that comment")
	}

	// Parse the body data
	updatedComment := new(commentv1.CommentMutable)
	if err := c.BodyParser(updatedComment); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	// Update the comment
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	comment, err = rpc.Svcs.GetCommentSvc().UpdateComment(
		ctx,
		&commentv1.UpdateCommentRequest{
			Id: c.Params("id"),
			Data: updatedComment,
		},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": comment})
}

func DeleteComment(c *fiber.Ctx) error {
	// check if user has permissions to delete the comment
	currentUser, err := getCurrentUser(c)
	if err != nil {
		return err
	}

	comment, err := getComment(c.Params("id"))
	if err != nil {
		return err
	}

	if (comment.AuthorId != currentUser.Id && !currentUser.IsAdmin) {
		return fiber.NewError(fiber.StatusForbidden, "no permissions to delete that comment")
	}

	// Delete the comment
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = rpc.Svcs.GetCommentSvc().DeleteComment(
		ctx,
		&commentv1.DeleteCommentRequest{Id: c.Params("id")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "comment deleted"})
}

func CreateComment(c *fiber.Ctx) error {
	// Parse the body data
	newComment := new(commentv1.CommentMutable)
	if err := c.BodyParser(newComment); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	// check post is real
	post, err := getPostById(c.Params("post_id"))
	if err != nil {
		return err
	}

	// access token claims
	tokenClaims, err := handlers.GetTokenClaims(c)
	if err != nil {
		return err
	}
	
	// Create the comment
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	comment, err := rpc.Svcs.GetCommentSvc().CreateComment(
		ctx,
		&commentv1.CreateCommentRequest{
			PostId: post.Id,
			AuthorId: tokenClaims.Subject,
			Data: newComment,
		},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": comment})
}