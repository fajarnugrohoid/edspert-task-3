package question

import (
	"course/internal/domain"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionUsecase struct {
	db *gorm.DB
}

func NewQuestionUsecase(db *gorm.DB) *QuestionUsecase {
	return &QuestionUsecase{db: db}
}

func (eu QuestionUsecase) GetQuestionByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid id",
		})
		return
	}
	var question domain.Question
	err = eu.db.Where("id = ?", id).Preload("Questions").Take(&question).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "not found",
		})
		return
	}
	c.JSON(200, question)
}

type Score struct {
	totalScore int
	mu         sync.Mutex
}

func (s *Score) Inc(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalScore += value
}
