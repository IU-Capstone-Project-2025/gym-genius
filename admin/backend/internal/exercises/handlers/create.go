package handlers

import (
	"admin/internal/database/models"
	"admin/internal/database/schemas"
	"admin/internal/database"
	"github.com/gofiber/fiber/v2"
	"errors"
	"gorm.io/gorm"
)

// CreateExercise
// @Summary Create a new exercise
// @Description Create a new exercise
// @Tags exercises
// @Accept json
// @Produce json
// @Param exercise body models.ExerciseCreate true "Exercise create payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /exercises [post]
func AddExercise(c *fiber.Ctx) error {
	exerciseCreate := &models.AddExercise{}

	if err := c.BodyParser(exerciseCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	
	exercise := &schemas.Exercise{
		Name: exerciseCreate.Name,
		URL: exerciseCreate.URL,
	}
	
	if err := database.DB.Create(exercise).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "exercise with this name already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create exercise",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Exercise created successfully",
		"id":      exercise.ID,
	})
}
