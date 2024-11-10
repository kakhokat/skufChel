// CourseStatisticsData.kt
package com.example.loginpage

data class CourseStatisticsData(
    val enrolledCount: Int,
    val completedCount: Int,
    val lastUpdated: String // You can use a Date type if preferred
)
