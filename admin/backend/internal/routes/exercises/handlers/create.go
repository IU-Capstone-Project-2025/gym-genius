package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"errors"
	middleware "admin/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateExercise
// @Security BearerAuth
// @Summary Create a new exercise
// @Tags exercises
// @Accept json
// @Produce json
// @Param exercise body models.ExerciseCreate true "Exercise create payload"
// @Success 200 {object} models.CreatedResponse
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /exercises [post]
func AddExercise(c *fiber.Ctx) error {
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

	exerciseCreate := &models.ExerciseCreate{}

	if err := c.BodyParser(exerciseCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: err.Error(),
		})
	}

	if exerciseCreate.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "name must be provided",
		})
	}

	exercise := &schemas.Exercise{
		Name:         exerciseCreate.Name,
		Description:  exerciseCreate.Description,
		MuscleGroups: exerciseCreate.MuscleGroups,
		ImagePath:          exerciseCreate.ImagePath,
	}

	if err := database.DB.Create(exercise).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Error: "exercise with this name already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to create exercise",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.CreatedResponse{
		Message: "exercise created successfully",
		ID:      exercise.ID,
	})
}
