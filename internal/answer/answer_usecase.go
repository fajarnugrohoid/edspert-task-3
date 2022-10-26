package answer

import (
	"course/internal/domain"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AnswerUsecase struct {
	db *gorm.DB
}

func NewAnswerUsecase(db *gorm.DB) *AnswerUsecase {
	return &AnswerUsecase{db: db}
}

func (eu AnswerUsecase) GetAnswerByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid id",
		})
		return
	}
	var answer domain.Answer
	err = eu.db.Where("id = ?", id).Preload("Answers").Take(&answer).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "not found",
		})
		return
	}
	c.JSON(200, answer)
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
