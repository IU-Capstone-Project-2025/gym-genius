package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/middlewares"
	"admin/internal/models"

	"github.com/gofiber/fiber/v2"
)

// GetExercisePhoto
// @Description Get an exercise photo by its ID
// @Security BearerAuth
// @Summary Get an exercise photo by ID
// @Tags exercises
// @Accept json
// @Produce json
// @Param id path int true "Exercise ID"
// @Success 200 {object} models.MessageResponse "Exercise photo retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Exercise Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /exercises/photo/{id} [get]
func GetExercisePhoto(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "required id parameter is malformed; should be > 0",
		})
	}

	roleRaw := c.Locals(middleware.RoleKey)

	role, ok := roleRaw.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Unauthorized or invalid token (role)",
		})
	}

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error: "This endpoint is restricted to admin users only",
		})
	}

	var exercise schemas.Exercise
	if err := database.DB.First(&exercise, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
			Error: "Exercise not found",
		})
	}

	return c.JSON(models.PhotoResponse{
		Message: "Exercise photo path returned successfully",
		Data:    exercise.ImagePath,
	})
}

// GetAllExercisePhoto
// @Description Get all exercise photos
// @Security BearerAuth
// @Summary Get all exercise photos
// @Tags exercises
// @Accept json
// @Produce json
// @Success 200 {object} models.MessageResponse "Exercise photos retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "Exercise Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /exercises/photo [get]
func GetAllExercisePhotos(c *fiber.Ctx) error {
	roleRaw := c.Locals(middleware.RoleKey)

	role, ok := roleRaw.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Unauthorized or invalid token (role)",
		})
	}

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(models.ErrorResponse{
			Error: "This endpoint is restricted to admin users only",
		})
	}

	var exercises []schemas.Exercise
	if err := database.DB.Find(&exercises).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "Failed to fetch exercises",
		})
	}

	photos := make([]string, len(exercises))
	for i, ex := range exercises {
		photos[i] = ex.ImagePath
	}

	return c.JSON(fiber.Map{
		"photos": photos,
	})
}