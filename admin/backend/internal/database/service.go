package database

import (
	"admin/config"
	"admin/internal/database/schemas"
	"admin/internal/models"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var (
	secret        = config.C.Secret
	jwtSecret     = config.C.JwtSecret
	dbHost        = config.C.DbHost
	dbUser        = config.C.DbUser
	dbPassword    = config.C.DbPassword
	dbName        = config.C.DbName
	dbPort        = config.C.DbPort
	adminPassword = config.C.AdminPassword
)

func Hash(login, password string) string {
	data := login + ":" + password + ":" + secret
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func CreateTokenForAdmin(user schemas.Admin) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"login": user.Login,
		"role":  "admin",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func CreateTokenForUser(user schemas.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"login": user.Login,
		"role":  "user",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func InitDatabase() error {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	switch config.C.AppEnv {
	case "PROD":
		DB, err = gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{
				TranslateError: true,                                  // fix to properly return errors
				Logger:         logger.Default.LogMode(logger.Silent), // silence the gorm logger
			},
		)
		// DB = DB.Debug() // debug postgres queries if needed
	case "DEV":
		DB, err = gorm.Open(
			sqlite.Open("devDb.db"),
			&gorm.Config{
				TranslateError: true, // fix to properly return errors
				// Logger: logger.Default.LogMode(logger.Silent), // silence the gorm logger
			},
		)
		DB = DB.Debug() // outputs generated sql to stdout
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(
		&schemas.User{},
		&schemas.Admin{},
		&schemas.UserActivity{},
		&schemas.Workout{},
		&schemas.Exercise{},
		&schemas.ExerciseSet{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return UpsertStaticData()
}

func UpsertStaticData() error {
	err := UpsertStaticUsers()
	if err != nil {
		return fmt.Errorf("failed to upsert static users: %w", err)
	}

	err = UpsertStaticAdmins()
	if err != nil {
		return fmt.Errorf("failed to upsert static admins: %w", err)
	}

	err = UpsertStaticExercises()
	if err != nil {
		return fmt.Errorf("failed to upsert static exercises: %w", err)
	}

	err = UpsertStaticWorkouts()
	if err != nil {
		return fmt.Errorf("failed to upsert static workouts: %w", err)
	}

	err = UpsertStaticUserActivities()
	if err != nil {
		return fmt.Errorf("failed to upsert static user activities: %w", err)
	}
	
	return nil
}

func UpsertStaticAdmins() error {
	var admin schemas.Admin

	admin.Login = "admin"
	admin.Hash = Hash(admin.Login, adminPassword)

	err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "login"}},
		UpdateAll: true,
	}).Create(&admin).Error
	if err != nil {
		return err
	}

	return nil
}

func UpsertStaticUsers() error {
	data, err := os.ReadFile("assets/users.json")
	if err != nil {
		return err
	}

	var rawUsers []models.UserCreateFull
	err = json.Unmarshal(data, &rawUsers)
	if err != nil {
		return err
	}
	var users []schemas.User
	for _, u := range rawUsers {
		dbUser := schemas.User{
			ID:                     u.ID,
			Name:                   u.Name,
			Surname:                u.Surname,
			Login:                  u.Login,
			Email:                  u.Email,
			Hash:                   Hash(u.Login, u.Password),
			SubscriptionPlan:       u.SubscriptionPlan,
			SubscriptionStatus:     u.SubscriptionStatus,
			Status:                 u.Status,
			StreakCount:            u.StreakCount,
			AverageWorkoutDuration: time.Duration(u.AverageWorkoutDurationNS),
			NumberOfWorkouts:       u.NumberOfWorkouts,
			TotalTimeSpent:         time.Duration(u.TotalTimeSpentNS),
			LastActivity:           u.LastActivity,
		}
		users = append(users, dbUser)
	}
	
	// skip hooks to avoid default user initizalization metadata
	err = DB.Session(&gorm.Session{SkipHooks: true}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		UpdateAll: true,
	}).Create(&users).Error
	if err != nil {
		return err
	}

	return nil
}

// reads assets/exercises.json and upserts them into the db
func UpsertStaticExercises() error {
	data, err := os.ReadFile("assets/exercises.json")
	if err != nil {
		return err
	}

	var exercises []schemas.Exercise
	err = json.Unmarshal(data, &exercises)
	if err != nil {
		return err
	}

	err = DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		UpdateAll: true,
	}).Create(&exercises).Error
	if err != nil {
		return err
	}

	return nil
}

func UpsertStaticWorkouts() error {
	data, err := os.ReadFile("assets/workouts.json")
	if err != nil {
		return err
	}

	var workoutsCreate []models.WorkoutCreate

	err = json.Unmarshal(data, &workoutsCreate)
	if err != nil {
		return err
	}

	for _, workoutCreate := range workoutsCreate {
		workoutDuration := time.Duration(workoutCreate.DurationNS)

		var exerciseSets []schemas.ExerciseSet
		for _, exerciseSet := range workoutCreate.ExerciseSets {
			exerciseSets = append(exerciseSets, schemas.ExerciseSet{
				ExerciseID: exerciseSet.ExerciseID,
				Weight:     exerciseSet.Weight,
				Reps:       exerciseSet.Reps,
			})
		}

		workout := &schemas.Workout{
			UserID:       workoutCreate.UserID,
			Duration:     workoutDuration,
			StartTime:    workoutCreate.StartTime,
			ExerciseSets: exerciseSets,
		}

		if err := DB.Create(&workout).Error; err != nil {
			return err
		}
	}

	return nil
}

func UpsertStaticUserActivities() error {
	data, err := os.ReadFile("assets/user_activities.json")
	if err != nil {
		return err
	}

	var userActivityCreates []models.UserActivityCreateFull

	err = json.Unmarshal(data, &userActivityCreates)
	if err != nil {
		return err
	}

	for _, userActivityCreate := range userActivityCreates {
		userActivity := &schemas.UserActivity{
			UserID:       userActivityCreate.UserID,
			ID: userActivityCreate.ID,
			Date: userActivityCreate.Date,
		}

		if err := DB.Create(userActivity).Error; err != nil {
			return err
		}
	}

	return nil
}