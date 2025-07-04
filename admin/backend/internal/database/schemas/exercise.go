package schemas

// TODO verify schemas

type Exercise struct {
	ID           uint   `gorm:"primaryKey;autoincrement"`
	Name         string `gorm:"not null"`
	URL          string
}

// type MuscleGroup struct {
// 	ID            uint            `gorm:"primaryKey"`
// 	Name          string          `gorm:"unique;not null"`
// 	ExerciseInfos []*ExerciseInfo `gorm:"many2many:exercise_info_muscle_groups;constraint:OnDelete:CASCADE;"`
// }

// type Exercise struct {
// 	ID             uint `gorm:"primaryKey"`
// 	WorkoutID      uint
// 	Workout        Workout `gorm:"constraint:OnDelete:CASCADE;"`
// 	ExerciseInfoID uint
// 	ExerciseInfo   ExerciseInfo
// }

type ExerciseSet struct {
	ID         uint    `gorm:"primaryKey"`
	Weight     float64 `gorm:"not null"`
	Reps       uint    `gorm:"not null"`
	ExerciseID uint
	Exercise   Exercise `gorm:"constraint:OnDelete:CASCADE;"`
}
