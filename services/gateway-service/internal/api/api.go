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

package api

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/hexolan/panels/gateway-service/internal"
	"github.com/hexolan/panels/gateway-service/internal/api/handlers"
)

func NewAPIApp(cfg internal.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "Panels REST Gateway",
		ErrorHandler: handlers.ErrorHandler,

		// Swap out the JSON encoder for faster marshaling
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Middleware
	handlers.NewAuthMiddleware(cfg)
	app.Use(cors.New())
	app.Use(logger.New())

	// Register the routes
	RegisterRoutes(app)

	return app
}

func ServeAPIApp(app *fiber.App) {
	err := app.Listen(":3000")
	if err != nil {
		panic(fmt.Sprintf("failed to serve API: %v", err))
	}
}
