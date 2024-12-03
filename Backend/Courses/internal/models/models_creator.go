package models

type Course struct {
	CourseId      int      `json:"courseId"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	CourseUsers   int      `json:"countUsers"`
	CourseLessons int      `json:"countLessons"`
	Lessons       []Lesson `json:"lessons"`
	Likes         int      `json:"likes"`
	CreatorId     int      `json:"creatorId"`
}

type Lesson struct {
	Id          int    `json:"lessonId"`
	Name        string `json:"name"`
	Likes       int    `json:"likes"`
	Description string `json:"description"`
	Test        Test   `json:"test"`
	Video       Video  `json:"video"`
}

type Test struct {
	Id           int    `json:"testId"`
	Questiong    string `json:"question"`
	Answers      string `json:"answers"`
	CurrentAnser string `json:"current_answer"`
}

type Video struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}
