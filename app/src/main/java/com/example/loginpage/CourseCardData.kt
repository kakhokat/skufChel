// CourseCardData.kt
package com.example.loginpage

enum class CourseStatus {
    IN_PROGRESS,
    COMPLETED
}

data class CourseCardData(
    val id: String,
    val courseName: String,
    val creatorName: String,
    val courseDescription: String,
    val lessonsCount: Int,
    val profileImageRes: Int,
    val likesCount: Int,
    val isLiked: Boolean,
    val enrolledCount: Int,
    val completedCount: Int,
    val lastUpdated: String,
    val status: CourseStatus,
    val lessons: List<LessonData>
)

// Sample data with adjusted lessonsCount and lessons
val courseCards = listOf(
    CourseCardData(
        id = "1",
        courseName = "Курс по Jetpack Compose",
        creatorName = "Иван Иванов",
        courseDescription = "Изучите основы Jetpack Compose.",
        lessonsCount = 3,
        profileImageRes = R.drawable.logo,
        likesCount = 123,
        isLiked = false,
        enrolledCount = 200,
        completedCount = 150,
        lastUpdated = "2023-10-01",
        status = CourseStatus.IN_PROGRESS,
        lessons = listOf(
            LessonData(
                courseId = "1",
                lessonId = "1-1",
                title = "Введение в Jetpack Compose",
                description = "Основы использования Jetpack Compose.",
                isCompleted = true,
                isLiked = false
            ),
            LessonData(
                courseId = "1",
                lessonId = "1-2",
                title = "Базовые компоненты",
                description = "Изучение базовых компонентов Compose.",
                isCompleted = false,
                isLiked = false
            ),
            LessonData(
                courseId = "1",
                lessonId = "1-3",
                title = "Layout и компоновка",
                description = "Узнайте о принципах компоновки в Compose.",
                isCompleted = false,
                isLiked = false
            )
        )
    ),
    // Existing courses 2 through 8 (not shown here for brevity)
    // Course ID 9
    CourseCardData(
        id = "9",
        courseName = "Разработка игр на Unity",
        creatorName = "Виктор Викторов",
        courseDescription = "Создавайте игры с помощью Unity Engine.",
        lessonsCount = 2,
        profileImageRes = R.drawable.logo,
        likesCount = 95,
        isLiked = true,
        enrolledCount = 220,
        completedCount = 180,
        lastUpdated = "2023-11-01",
        status = CourseStatus.IN_PROGRESS,
        lessons = listOf(
            LessonData(
                courseId = "9",
                lessonId = "9-1",
                title = "Введение в Unity",
                description = "Познакомьтесь с основами Unity.",
                isCompleted = false,
                isLiked = false
            ),
            LessonData(
                courseId = "9",
                lessonId = "9-2",
                title = "Создание первого проекта",
                description = "Создайте свой первый проект в Unity.",
                isCompleted = false,
                isLiked = false
            )
        )
    ),
    // Course ID 10
    CourseCardData(
        id = "10",
        courseName = "Веб-разработка с Django",
        creatorName = "Ольга Ольгова",
        courseDescription = "Создание веб-приложений на Python с использованием Django.",
        lessonsCount = 3,
        profileImageRes = R.drawable.logo,
        likesCount = 130,
        isLiked = true,
        enrolledCount = 400,
        completedCount = 350,
        lastUpdated = "2023-10-20",
        status = CourseStatus.COMPLETED,
        lessons = listOf(
            LessonData(
                courseId = "10",
                lessonId = "10-1",
                title = "Введение в Django",
                description = "Основы веб-разработки с Django.",
                isCompleted = true,
                isLiked = false
            ),
            LessonData(
                courseId = "10",
                lessonId = "10-2",
                title = "Модели и базы данных",
                description = "Работа с моделями и базами данных в Django.",
                isCompleted = true,
                isLiked = false
            ),
            LessonData(
                courseId = "10",
                lessonId = "10-3",
                title = "Создание шаблонов",
                description = "Использование шаблонов для отображения данных.",
                isCompleted = true,
                isLiked = false
            )
        )
    )
)
