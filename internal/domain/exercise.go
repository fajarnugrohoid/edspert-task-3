package domain

import "time"

type Exercise struct {
	ID          int
	Title       string
	Description string
	Questions   []Question
}

type Answer struct {
	ID         int
	ExerciseID int
	QuestionID int
	UserID     int
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
