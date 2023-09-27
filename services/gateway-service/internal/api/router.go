package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/api/v1"
	"github.com/hexolan/panels/gateway-service/internal/api/handlers"
)

func RegisterRoutes(app *fiber.App) {
	apiV1 := app.Group("/v1")
	panelV1 := apiV1.Group("/panels")
	postV1 := apiV1.Group("/posts")
	userV1 := apiV1.Group("/users")
	authV1 := apiV1.Group("/auth")
	commentV1 := postV1.Group("/:post_id/comments")

	// Panel Service Routes
	panelV1.Post("/", handlers.AuthMiddleware, v1.CreatePanel)

	panelV1.Get("/id/:id", v1.GetPanelById)
	panelV1.Patch("/id/:id", handlers.AuthMiddleware, v1.UpdatePanelById)
	panelV1.Delete("/id/:id", handlers.AuthMiddleware, v1.DeletePanelById)

	panelV1.Get("/name/:name", v1.GetPanelByName)
	panelV1.Patch("/name/:name", handlers.AuthMiddleware, v1.UpdatePanelByName)
	panelV1.Delete("/name/:name", handlers.AuthMiddleware, v1.DeletePanelByName)

	// Post Service Routes
	postV1.Get("/feed", v1.GetFeedPosts)
	postV1.Patch("/:id", handlers.AuthMiddleware, v1.UpdatePost)
	postV1.Delete("/:id", handlers.AuthMiddleware, v1.DeletePost)

	userV1.Get("/id/:user_id/posts", v1.GetUserPostsFromId)
	userV1.Get("/username/:username/posts", v1.GetUserPostsFromUsername)

	panelV1.Get("/id/:panel_id/posts", v1.GetPanelPostsFromId)
	panelV1.Get("/name/:panel_name/posts", v1.GetPanelPostsFromName)
	
	panelV1.Get("/id/:panel_id/posts/:id", v1.GetPanelPostFromId)
	panelV1.Get("/name/:panel_name/posts/:id", v1.GetPanelPostFromName)
	
	panelV1.Post("/id/:panel_id", handlers.AuthMiddleware, v1.CreatePanelPostFromId)
	panelV1.Post("/name/:panel_name", handlers.AuthMiddleware, v1.CreatePanelPostFromName)

	// User Service Routes
	userV1.Post("/", v1.UserSignup)

	userV1.Get("/id/:id", v1.GetUserById)
	userV1.Delete("/id/:id", handlers.AuthMiddleware, v1.DeleteUserById)

	userV1.Get("/username/:username", v1.GetUserByUsername)
	userV1.Delete("/username/:username", handlers.AuthMiddleware, v1.DeleteUserByUsername)
	
	userV1.Get("/me", handlers.AuthMiddleware, v1.GetCurrentUser)
	userV1.Delete("/me", handlers.AuthMiddleware, v1.DeleteCurrentUser)
	
	// Auth Service Routes
	authV1.Post("/login", v1.LoginWithPassword)

	// Comment Service Routes
	commentV1.Get("/", v1.GetPostComments)
	commentV1.Post("/", handlers.AuthMiddleware, v1.CreateComment)
	commentV1.Patch("/:id", handlers.AuthMiddleware, v1.UpdateComment)
	commentV1.Delete("/:id", handlers.AuthMiddleware, v1.DeleteComment)
}