// LessonData.kt
package com.example.loginpage

data class LessonData(
    val courseId: String,
    val lessonId: String,
    val title: String,
    val description: String,
    val isCompleted: Boolean = false,
    val isLiked: Boolean = false
)