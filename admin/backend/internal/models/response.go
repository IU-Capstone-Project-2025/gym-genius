package models

import (
	"time"
)

type UserRead struct {
	ID                       uint      `json:"id" example:"12345"`
	Login                    string    `json:"username"`
	Name                     string    `json:"name" example:"John"`
	Surname                  string    `json:"surname" example:"Doe"`
	Email                    string    `json:"email" example:"john_doe@gmail.com"`
	SubscriptionType         string    `json:"subscription_type" example:"free"` // e.g., "free", "basic", "pro"
	SubscriptionStatus       string    `json:"subscription_status" example:"active"` // e.g., "active", "expired", "cancelled"
	Status                   string    `json:"status" example:"active"`          // e.g., "active", "inactive", "banned"
	LastActivity             time.Time `json:"last_activity" example:"2023-10-01T12:00:00Z"`
	NumberOfWorkouts         uint      `json:"number_of_workouts" example:"0"`
	TotalTimeSpentNS         int64     `json:"total_time_spent_ns" example:"3600"` // in nanoseconds
	StreakCount              uint      `json:"streak_count" example:"0"`
	AverageWorkoutDurationNS int64     `json:"average_workout_duration_ns" example:"3600"` // in nanoseconds
}

type ExerciseSetRead struct {
	Weight     float64 `json:"weight"      example:"10"`
	Reps       uint    `json:"reps"        example:"10"`
	ExerciseID uint    `json:"exercise_id" example:"10"`
	WorkoutID  uint    `json:"-"           example:"10"`
}

type WorkoutRead struct {
	ID           uint              `json:"id" example:"1"`
	DurationNS   int64             `json:"duration_ns" example:"60"` // in nanoseconds
	Timestamp    time.Time         `json:"timestamp" example:"2023-10-01T12:00:00Z"`
	UserID       uint              `json:"user_id" example:"12345"`
	ExerciseSets []ExerciseSetRead `json:"exercise_sets"`
}

type ExerciseRead struct {
	ID           uint     `json:"id" example:"1"`
	Name         string   `json:"name"`
	Description  string   `json:"description" example:"Push-ups are a basic exercise that works the chest, shoulders, and triceps."`
	MuscleGroups []string `json:"muscle_groups" example:"chest,back,triceps"`
	URL          string   `json:"url" example:"https://example.com/image.jpg"` // URL to the exercise image
}

type MessageResponse struct {
	Message string `json:"message" example:"Descriptive message"`
}

type PhotoResponse struct {
	Message string `json:"message" example:"Descriptive message"`
	Data	string `json:"data" example:"https://example.com/image.jpg"` // URL to the exercise photo
}

type CreatedResponse struct {
	Message string `json:"message" example:"User created successfully"`
	ID      uint   `json:"id" example:"12345"`
}

type TokenResponse struct {
	Message string `json:"message" example:"Token created successfully"`
	Token string `json:"token" example:"token_value_here"` // JWT token
}

type ErrorResponse struct {
	Error string `json:"error" example:"A descriptive error message"`
}

type CountResponse struct {
	Count int64 `json:"count" example:"10"`
}