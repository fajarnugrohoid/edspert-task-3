package question

import (
	"course/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionUsecase struct {
	db *gorm.DB
}

type QuestionRequest struct {
	ExerciseID    int    `json:"exercise_id"`
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
	Score         int    `json:"score"`
	CreatorID     int    `json:"creator_id"`
}

func NewQuestionUsecase(db *gorm.DB) *QuestionUsecase {
	return &QuestionUsecase{db: db}
}

func (eu QuestionUsecase) GetQuestions(c *gin.Context) {

	var questions []*domain.Question
	err := eu.db.Find(&questions).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "not found",
		})
		return
	}
	c.JSON(200, questions)
}

func (qu QuestionUsecase) Insert(c *gin.Context) {
	var questionReq QuestionRequest
	if err := c.ShouldBind(&questionReq); err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	question, err := domain.NewQuestion(
		questionReq.ExerciseID,
		questionReq.Body,
		questionReq.OptionA,
		questionReq.OptionB,
		questionReq.OptionC,
		questionReq.OptionD,
		questionReq.CorrectAnswer,
		questionReq.Score,
		questionReq.CreatorID,
	)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": err.Error(),
		})
		return
	}

	if err := qu.db.Create(&question).Error; err != nil {
		c.JSON(500, map[string]string{
			"message": "error when create question",
		})
		return
	}

	c.JSON(200, map[string]string{
		"insert": "success",
	})
}
