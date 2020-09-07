package apiserver

import "github.com/gofiber/fiber"

// Routes list of the available routes for project
func Routes(app *fiber.App) {
	// Create group for API routes
	APIGroup := app.Group("/v1/api")

	// API routes
	APIGroup.Get("/docs", GetDocs)

}

// GetDocs will send the documents list
func GetDocs(ctx *fiber.Ctx) {
	// Set JSON data
	data := fiber.Map{
		"message": "ok",
		"results": []fiber.Map{
			{
				"name": "Documentation",
				"url":  "https://create-go.app/",
			},
			{
				"name": "Detailed guides",
				"url":  "https://create-go.app/detailed-guides/",
			},
			{
				"name": "GitHub",
				"url":  "https://github.com/create-go-app/cli",
			},
		},
	}

	// Set 200 OK status and return JSON, or skip route
	if errJSON := ctx.Status(200).JSON(data); errJSON != nil {
		// Send 500 status and error to Fiber
		ctx.Status(500).Next(errJSON)
	}
}
