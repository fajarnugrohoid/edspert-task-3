package domain

import (
	"errors"
	"time"
)

type Answer struct {
	ID         int
	ExerciseID int
	QuestionID int
	UserID     int
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewAnswer(exerciseID int, questionID int, userID int, answer string) (Answer, error) {
	if answer == "" {
		return Answer{}, errors.New("answer cannot be empty")
	}

	return Answer{
		ExerciseID: exerciseID,
		QuestionID: questionID,
		UserID:     userID,
		Answer:     answer,
	}, nil
}
