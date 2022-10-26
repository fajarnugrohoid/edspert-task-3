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
