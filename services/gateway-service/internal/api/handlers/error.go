package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
    c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
    code := fiber.StatusInternalServerError
	msg := err.Error()

    // Retrieval of codes from fiber.Errors
    var e *fiber.Error
    if errors.As(err, &e) {
        code = e.Code
    } else {
		// Retrival of codes from gRPC errors.
		status, ok := status.FromError(err)
		if ok {
			msg = status.Message()

			switch status.Code() {
			case codes.NotFound:
				code = fiber.StatusNotFound
			case codes.InvalidArgument:
				code = fiber.StatusUnprocessableEntity
			case codes.AlreadyExists:
				code = fiber.StatusConflict
			case codes.PermissionDenied:
				code = fiber.StatusForbidden
			case codes.Unauthenticated:
				code = fiber.StatusUnauthorized
			case codes.Internal:
				code = fiber.StatusInternalServerError
			case codes.Unavailable:
				code = fiber.StatusBadGateway
				msg = "Service unavaliable for request."
			default:
				code = fiber.StatusInternalServerError
				msg = "Something went wrong."
				log.Error(err)
			}
		} else {
			msg = "Something unexpected went wrong."
			log.Error(err)
		}
	}

    return c.Status(code).JSON(fiber.Map{"status": "failure", "msg": msg})
}