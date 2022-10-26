package domain

import (
	"errors"
	"time"
)

type Question struct {
	ID            int
	ExerciseID    int
	Body          string
	OptionA       string
	OptionB       string
	OptionC       string
	OptionD       string
	CorrectAnswer string
	Score         int
	CreatorID     int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewQuestion(ExerciseID int, Body string,
	OptionA string, OptionB string, OptionC string, OptionD string,
	CorrectAnswer string, Score int, CreatorID int) (Question, error) {
	if Body == "" {
		return Question{}, errors.New("Body cannot be empty")
	}

	return Question{
		ExerciseID:    ExerciseID,
		Body:          Body,
		OptionA:       OptionA,
		OptionB:       OptionB,
		OptionC:       OptionC,
		OptionD:       OptionD,
		CorrectAnswer: CorrectAnswer,
		Score:         Score,
		CreatorID:     CreatorID,
	}, nil
}
