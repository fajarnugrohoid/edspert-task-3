package exercise

import (
	"course/internal/domain"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseUsecase struct {
	db *gorm.DB
}

func NewExerciseUsecase(db *gorm.DB) *ExerciseUsecase {
	return &ExerciseUsecase{db: db}
}

func (eu ExerciseUsecase) GetExerciseByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid id",
		})
		return
	}
	var exercise domain.Exercise
	err = eu.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "not found",
		})
		return
	}
	c.JSON(200, exercise)
}

func (eu ExerciseUsecase) GetScore(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(400, map[string]string{
			"message": "invalid id",
		})
		return
	}
	var exercise domain.Exercise
	err = eu.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, map[string]string{
			"message": "not found",
		})
		return
	}
	userID := c.Request.Context().Value("user_id").(int)
	fmt.Println("user_id", userID)
	var answers []domain.Answer
	err = eu.db.Where("exercise_id = ? AND user_id = ?", id, userID).Find(&answers).Error
	if err != nil {
		c.JSON(404, map[string]interface{}{
			"message": "not answered yet",
		})
		return
	}

	mapQA := make(map[int]domain.Answer)
	for _, answer := range answers {
		mapQA[answer.QuestionID] = answer
	}

	var score Score
	wg := new(sync.WaitGroup)
	for _, question := range exercise.Questions {
		wg.Add(1)
		go func(question domain.Question) {
			defer wg.Done()
			if strings.EqualFold(question.CorrectAnswer, mapQA[question.ID].Answer) {
				score.Inc(question.Score)
			}
		}(question)
	}
	wg.Wait()
	c.JSON(200, map[string]int{
		"score": score.totalScore,
	})
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
