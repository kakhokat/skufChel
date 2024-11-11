// ProfileScreen.kt
package com.example.loginpage

import androidx.compose.foundation.gestures.detectHorizontalDragGestures
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.input.pointer.pointerInput
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController

@Composable
fun ProfileScreen(
    profileCardData: ProfileCardData?,
    courseCards: List<CourseCardData>,
    sharedViewModel: SharedViewModel,
    navController: NavController
) {
    if (profileCardData != null) {
        // Состояние выбранной вкладки
        var selectedTab by remember { mutableStateOf(CourseStatus.IN_PROGRESS) }
        val filteredCourses = courseCards.filter { it.status == selectedTab }

        Box(
            modifier = Modifier
                .fillMaxSize()
                .pointerInput(Unit) {
                    detectHorizontalDragGestures { change, dragAmount ->
                        change.consume()
                        if (dragAmount < -50) {
                            // Обнаружен свайп влево
                            navController.navigate("all_courses")
                        }
                    }
                }
        ) {
            // Основной контент
            Column(
                modifier = Modifier
                    .fillMaxSize()
                    .padding(16.dp)
            ) {
                // Профиль пользователя
                ProfileCard(
                    profileCardData = profileCardData,
                    onExitClick = {
                        // Переход на экран входа
                        navController.navigate("login") {
                            popUpTo("login") { inclusive = true }
                        }
                    }
                )

                Spacer(modifier = Modifier.height(16.dp))

                // Кнопки переключения статуса курсов
                Row(
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(bottom = 16.dp),
                    horizontalArrangement = Arrangement.SpaceEvenly
                ) {
                    Button(
                        onClick = { selectedTab = CourseStatus.IN_PROGRESS },
                        colors = ButtonDefaults.buttonColors(
                            containerColor = if (selectedTab == CourseStatus.IN_PROGRESS) Color.Blue else Color.Gray
                        )
                    ) {
                        Text(text = "Прохожу")
                    }
                    Button(
                        onClick = { selectedTab = CourseStatus.COMPLETED },
                        colors = ButtonDefaults.buttonColors(
                            containerColor = if (selectedTab == CourseStatus.COMPLETED) Color.Blue else Color.Gray
                        )
                    ) {
                        Text(text = "Пройденные")
                    }
                }

                // Список курсов
                Column(
                    modifier = Modifier
                        .fillMaxWidth()
                        .weight(1f)
                        .verticalScroll(rememberScrollState())
                ) {
                    if (filteredCourses.isEmpty()) {
                        Text(
                            text = "Нет курсов для отображения",
                            modifier = Modifier.align(Alignment.CenterHorizontally)
                        )
                    } else {
                        filteredCourses.forEach { courseCard ->
                            CourseCard(
                                courseData = courseCard,
                                navController = navController,
                                onLikeClick = { /* Обработка лайка */ }
                            )
                        }
                    }
                }
            }

            // Навигационная кнопка, наложенная поверх контента
            IconButton(
                onClick = {
                    navController.navigate("all_courses")
                },
                modifier = Modifier
                    .align(Alignment.BottomEnd)
                    .padding(16.dp)
            ) {
                Column(
                    horizontalAlignment = Alignment.CenterHorizontally
                ) {
                    Icon(
                        painter = painterResource(id = R.drawable.nav_arrow_right),
                        contentDescription = "Все курсы",
                        modifier = Modifier.size(48.dp)
                    )
                    Text(
                        text = "Все курсы",
                        fontSize = 20.sp,
                        color = MaterialTheme.colorScheme.onSurface
                    )
                }
            }
        }
    } else {
        // Обработка случая, когда данные профиля недоступны
        Text("Нет данных профиля")
    }
}
