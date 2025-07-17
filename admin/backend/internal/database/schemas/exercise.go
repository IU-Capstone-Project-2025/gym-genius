package schemas

type Exercise struct {
	ID           uint     `gorm:"primaryKey;autoincrement"           json:"id"`
	Name         string   `gorm:"not null;uniqueIndex"               json:"name"`
	Description  string   `gorm:"not null"                           json:"description"`
	MuscleGroups []string `gorm:"type:text;serializer:json;not null" json:"muscle_groups"`
	ImagePath          string   `gorm:"not null"                           json:"image_path"`
}

type ExerciseSet struct {
	Weight     float64 `gorm:"not null"`
	Reps       uint    `gorm:"not null"`
	ExerciseID uint    `gorm:"primaryKey;autoIncrement:false"`
	WorkoutID  uint    `gorm:"primaryKey;autoIncrement:false"`
	Exercise   Exercise
	Workout    Workout
}
