package handlers

import (
	"admin/internal/database"
	"admin/internal/database/schemas"
	middleware "admin/internal/middlewares"
	"admin/internal/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetUserByID
// @Security BearerAuth
// @Summary Get any user by ID (Admin only)
// @Description Retrieve any user by their unique ID (Admin privileges required)
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserRead
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 403 {object} models.ErrorResponse "Forbidden"
// @Failure 404 {object} models.ErrorResponse "User Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil || id <= 0 {
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "'id' parameter is malformed; should be > 0",
        })
    }

    return getUserByID(uint(id), c)
}

// GetCurrentUser
// @Security BearerAuth
// @Summary Get current user
// @Description Retrieve the currently authenticated user's data
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.UserRead
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/me [get]
func GetCurrentUser(c *fiber.Ctx) error {
	userIDRaw := c.Locals(middleware.IDKey)

	userID, ok := userIDRaw.(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Error: "Unauthorized or invalid token (user ID)",
		})
	}

	return getUserByID(uint(userID), c)
}

func getUserByID(id uint, c *fiber.Ctx) error {
    user := &schemas.User{}

    if err := database.DB.First(user, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{
                Error: "user not found",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
            Error: "failed to retrieve user",
        })
    }

    userRead := models.UserRead{
        ID:                       user.ID,
        Login:                    user.Login,
        Name:                     user.Name,
        Surname:                  user.Surname,
        Email:                    user.Email,
        SubscriptionType:         user.SubscriptionPlan,
        Status:                   user.Status,
        LastActivity:             user.LastActivity,
        NumberOfWorkouts:         user.NumberOfWorkouts,
        TotalTimeSpentNS:         int64(user.TotalTimeSpent.Seconds()),
        StreakCount:              user.StreakCount,
        AverageWorkoutDurationNS: int64(user.AverageWorkoutDuration.Seconds()),
    }

    return c.Status(fiber.StatusOK).JSON(userRead)
}

type userQueryParams struct {
	Page               int    `query:"page"`
	Limit              int    `query:"limit"`
	UserStatus         string `query:"user_status"`
	SubscriptionPlan   string `query:"subscription_plan"`
	SubscriptionStatus string `query:"subscription_status"`
}

// GetUserPaginate
// @Security BearerAuth
// @Summary Get paginated list of users
// @Description Retrieve a paginated list of users with optional page and limit query parameters
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of users per page" default(10)
// @Param user_status query string false "User status to fliter by" default(active)
// @Param subscription_plan query string false "Subscription plan to filter by" default(basic)
// @Param subscription_status query string false "Subscription status to filter by" default(active)
// @Success 200 {object} []models.UserRead
// @Failure 400 {object} models.ErrorResponse "Malformed query parameters"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users [get]
func GetUsersPaginate(c *fiber.Ctx) error {
	params := &userQueryParams{}

	if err := c.QueryParser(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "malformed query parameters",
		})
	}

	errMessage := ""
	switch {
	case params.UserStatus != "" && !schemas.UserStatusExists[params.UserStatus]:
		errMessage = "invalid user_status query param"
	case params.SubscriptionStatus != "" && !schemas.SubscriptionStatusExists[params.SubscriptionStatus]:
		errMessage = "invalid subscription_status query param"
	case params.SubscriptionPlan != "" && !schemas.SubscriptionPlanExists[params.SubscriptionPlan]:
		errMessage = "invalid subscription_plan query param"
	}

	if errMessage != "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: errMessage,
		})
	}

	offset := (params.Page - 1) * params.Limit

	var users []schemas.User

	err :=
		database.DB.
			Limit(params.Limit).
			Offset(offset).
			Where(&schemas.User{
				SubscriptionPlan:   params.SubscriptionPlan,
				SubscriptionStatus: params.SubscriptionStatus,
				Status:             params.UserStatus,
			}).
			Find(&users).
			Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "failed to find users",
		})
	}

	userReads := make([]models.UserRead, len(users))
	for i, user := range users {
		userReads[i] = models.UserRead{
			ID:                       user.ID,
			Login:                    user.Login,
			Name:                     user.Name,
			Surname:                  user.Surname,
			Email:                    user.Email,
			SubscriptionType:         user.SubscriptionPlan,
			Status:                   user.Status,
			LastActivity:             user.LastActivity,
			NumberOfWorkouts:         user.NumberOfWorkouts,
			TotalTimeSpentNS:         user.TotalTimeSpent.Nanoseconds(),
			StreakCount:              user.StreakCount,
			AverageWorkoutDurationNS: user.AverageWorkoutDuration.Nanoseconds(),
		}
	}

	return c.Status(fiber.StatusOK).JSON(userReads)
}

// GetUserCount
// @Security BearerAuth
// @Summary Get the total number of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.CountResponse "User count object"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /users/count [get]
func GetUserCount(c *fiber.Ctx) error {
	var count int64

	if err := database.DB.Model(&schemas.User{}).Count(&count).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: "failed to count users",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.CountResponse{
		Count: count,
	})
}
