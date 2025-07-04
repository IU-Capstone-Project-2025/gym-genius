package workouts

import (
	"admin/internal/workouts/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpWorkoutRoutes(app *fiber.App) {
	app.Post("/workouts/", handlers.CreateWorkout)
	app.Get("/workouts/:id", handlers.GetWorkout)
	app.Patch("/workouts/:id", handlers.UpdateWorkout)
	app.Delete("/workouts/:id", handlers.DeleteWorkout)
}