package answer

import (
	"course/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnswerUsecase struct {
	db *gorm.DB
}

func NewAnswerUsecase(db *gorm.DB) *AnswerUsecase {
	return &AnswerUsecase{db: db}
}

type AnswerRequest struct {
	ExerciseID int    `json:"exercise_id"`
	QuestionID int    `json:"question_id"`
	UserID     int    `json:"user_id"`
	Answer     string `json:"answer"`
}

func (eu AnswerUsecase) GetAnswers(c *gin.Context) {
	var answers []*domain.Answer
	err := eu.db.Find(&answers).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "not found",
		})
		return
	}
	c.JSON(200, answers)
}

func (qu AnswerUsecase) Insert(c *gin.Context) {
	var answerReq AnswerRequest
	if err := c.ShouldBind(&answerReq); err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	answer, err := domain.NewAnswer(
		answerReq.ExerciseID,
		answerReq.QuestionID,
		answerReq.UserID,
		answerReq.Answer,
	)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	if err := qu.db.Create(&answer).Error; err != nil {
		c.JSON(500, map[string]string{
			"message": "error when create question",
		})
		return
	}

	c.JSON(200, map[string]string{
		"insert": "success",
	})
}
