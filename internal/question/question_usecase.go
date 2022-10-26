package question

import (
	"course/internal/domain"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionUsecase struct {
	db *gorm.DB
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
