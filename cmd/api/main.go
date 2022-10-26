package main

import (
	"course/internal/answer"
	"course/internal/database"
	"course/internal/exercise"
	"course/internal/middleware"
	"course/internal/question"
	"course/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/hello", func(c *gin.Context) {
		c.JSON(200, map[string]string{
			"message": "hello world",
		})
	})

	dbConn := database.NewDatabaseConn()
	eu := exercise.NewExerciseUsecase(dbConn)
	qu := question.NewQuestionUsecase(dbConn)
	an := answer.NewAnswerUsecase(dbConn)
	uu := user.NewUserUsecase(dbConn)

	// usecase endpoint
	route.GET("/exercises/:id", middleware.WithAuth(), eu.GetExerciseByID)
	route.GET("/exercises/:id/scores", middleware.WithAuth(), eu.GetScore)

	// usecase endpoint
	route.GET("/questions", middleware.WithAuth(), qu.GetQuestions)
	route.GET("/answer", middleware.WithAuth(), an.GetAnswers)

	// user endpoint
	route.POST("/register", uu.Register)
	route.POST("/login", uu.Login)

	route.Run(":1234")
}
