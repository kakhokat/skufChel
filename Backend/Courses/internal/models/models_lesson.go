package models

type LessonFull struct { //todo нейминг поменять
	Lesson
	IsLiked  bool `json:"isLiked"`
	IsPassed bool `json:"isPassed"`
}

type Answer struct {
	Answer string `json:"answer"`
}

type ResultOfAnswering struct {
	IsCorrect bool `json:"isCorrect"`
}
