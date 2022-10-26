package domain

type Exercise struct {
	ID          int
	Title       string
	Description string
	Questions   []Question
}
