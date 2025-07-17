package exercises

import (
	middleware "admin/internal/middlewares"
	"admin/internal/routes/exercises/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupExerciseRoutes(app *fiber.App) {
	app.Post("/exercises/", middleware.JWTMiddleware, handlers.AddExercise)
	app.Get("/exercises/", middleware.JWTMiddleware, handlers.GetExercisesPaginate)
	app.Get("/exercises/photo/:id", middleware.JWTMiddleware, handlers.GetExercisePhoto)
	app.Get("/exercises/photo/", middleware.JWTMiddleware, handlers.GetAllExercisePhotos)
	app.Get("/exercises/:id", middleware.JWTMiddleware, handlers.GetExerciseByID)
	app.Patch("/exercises/:id", middleware.JWTMiddleware, handlers.UpdateExercise)
	app.Delete("/exercises/:id", middleware.JWTMiddleware, handlers.DeleteExercise)
}
