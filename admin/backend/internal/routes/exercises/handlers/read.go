package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetExerciseByID
// @Security BearerAuth
// @Summary Get an exercise by id
// @Tags exercises
// @Accept json
// @Produce json
// @Param id path int true "Exercise ID"
// @Success 200 {object} models.ExerciseRead
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /exercises/{id} [get]
func GetExerciseByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "invalid exercise ID",
		})
	}

	exercise := &schemas.Exercise{}

	if err := database.DB.First(exercise, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
				Error: "exercise not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to retrieve exercise",
		})
	}

	exerciseRead := models.ExerciseRead{
		ID:           exercise.ID,
		Name:         exercise.Name,
		Description:  exercise.Description,
		MuscleGroups: exercise.MuscleGroups,
		URL:          exercise.ImagePath,
	}

	return c.Status(fiber.StatusOK).JSON(exerciseRead)
}

func GetExercisesPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}
	limit := c.QueryInt("limit", 10)
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var exercises []schemas.Exercise
	if err := database.DB.Limit(limit).Offset(offset).Find(&exercises).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: "failed to retrieve users",
		})
	}

	exerciseReads := make([]models.ExerciseRead, len(exercises))
	for i, exercise := range exercises {
		exerciseReads[i] = models.ExerciseRead{
			ID:           exercise.ID,
			Name:         exercise.Name,
			Description:  exercise.Description,
			MuscleGroups: exercise.MuscleGroups,
			URL:          exercise.ImagePath,
		}
	}

	return c.Status(fiber.StatusOK).JSON(exerciseReads)
}
