// AppNavigation.kt
package com.example.loginpage

import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.navigation.NavType
import androidx.navigation.compose.*
import androidx.navigation.navArgument

@Composable
fun AppNavigation(sharedViewModel: SharedViewModel) {
    val navController = rememberNavController()

    val courseCards = com.example.loginpage.courseCards

    NavHost(navController = navController, startDestination = "login") {
        composable("login") {
            LoginScreen(navController = navController, sharedViewModel = sharedViewModel)
        }
        composable("profile") {
            ProfileScreen(
                profileCardData = sharedViewModel.profileCardData.value,
                courseCards = courseCards,
                sharedViewModel = sharedViewModel,
                navController = navController
            )
        }
        composable(
            route = "course/{courseId}",
            arguments = listOf(navArgument("courseId") { type = NavType.StringType })
        ) { backStackEntry ->
            val courseId = backStackEntry.arguments?.getString("courseId")
            val courseCardData = courseCards.find { it.id == courseId }
            if (courseCardData != null) {
                CourseMainPage(courseCardData = courseCardData, navController = navController)
            } else {
                // Handle course not found
                Text("Course not found")
            }
        }
        // Новый маршрут для всех курсов
        composable("all_courses") {
            AllCoursesScreen(
                navController = navController,
                courseCards = courseCards
            )
        }
    }
}
