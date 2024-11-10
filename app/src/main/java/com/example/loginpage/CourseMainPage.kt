// CourseMainPage.kt
package com.example.loginpage

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.CheckCircle
import androidx.compose.material.icons.filled.Favorite
import androidx.compose.material.icons.filled.FavoriteBorder
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import androidx.navigation.compose.rememberNavController
import androidx.compose.ui.tooling.preview.Preview

@Composable
fun CourseMainPage(
    courseCardData: CourseCardData,
    navController: NavController
) {
    // Remember the list of lessons to manage state
    val lessons = remember { mutableStateListOf<LessonData>().apply { addAll(courseCardData.lessons) } }

    Column(
        modifier = Modifier.fillMaxSize(),
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        // Scrollable content
        Column(
            modifier = Modifier
                .weight(1f)
                .padding(16.dp)
                .verticalScroll(rememberScrollState()),
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            // Course Title
            Text(
                text = courseCardData.courseName,
                style = MaterialTheme.typography.headlineMedium,
                textAlign = TextAlign.Center
            )

            Spacer(modifier = Modifier.height(16.dp))

            // Author's Photo and Name
            Row(
                verticalAlignment = Alignment.CenterVertically
            ) {
                Image(
                    painter = painterResource(id = courseCardData.profileImageRes),
                    contentDescription = "Author's Photo",
                    modifier = Modifier
                        .size(50.dp)
                        .clip(RoundedCornerShape(25.dp))
                )
                Spacer(modifier = Modifier.width(8.dp))
                Text(text = courseCardData.creatorName)
            }

            Spacer(modifier = Modifier.height(16.dp))

            // Course Description
            Text(
                text = courseCardData.courseDescription,
                textAlign = TextAlign.Center
            )

            Spacer(modifier = Modifier.height(16.dp))

            // Course Statistics Card
            CourseStatisticsCard(
                courseStatisticsData = CourseStatisticsData(
                    enrolledCount = courseCardData.enrolledCount,
                    completedCount = courseCardData.completedCount,
                    lastUpdated = courseCardData.lastUpdated
                )
            )

            Spacer(modifier = Modifier.height(16.dp))

            // "Содержание курса" Text
            Text(
                text = "Содержание курса",
                style = MaterialTheme.typography.headlineSmall,
                textAlign = TextAlign.Center
            )

            Spacer(modifier = Modifier.height(16.dp))

            // Scrollable Area with Lesson Cards
            Column(
                verticalArrangement = Arrangement.spacedBy(8.dp),
                modifier = Modifier.fillMaxWidth()
            ) {
                lessons.forEachIndexed { index, lesson ->
                    LessonCard(
                        lessonData = lesson,
                        onLikeClick = { updatedLesson ->
                            lessons[index] = updatedLesson
                        }
                    )
                }
            }

            Spacer(modifier = Modifier.height(16.dp))
        }

        // Fixed Navigation Hints at the Bottom
        Divider()
        Row(
            modifier = Modifier
                .fillMaxWidth()
                .padding(8.dp)
                .background(MaterialTheme.colorScheme.surface),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            // Column 1: Go Back
            Column(
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                IconButton(onClick = { navController.popBackStack() }) {
                    Icon(
                        painter = painterResource(id = R.drawable.nav_arrow_left),
                        contentDescription = "Go Back"
                    )
                }
                Text(text = "Вернуться")
            }

            // Column 2: Likes
            Column(
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                IconButton(onClick = { /* Handle like action */ }) {
                    Icon(
                        painter = painterResource(
                            id = if (courseCardData.isLiked) R.drawable.icom_liked else R.drawable.icom_not_liked
                        ),
                        contentDescription = "Like"
                    )
                }
                Text(text = courseCardData.likesCount.toString())
            }

            // Column 3: Start
            Column(
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                IconButton(onClick = { /* Handle start action */ }) {
                    Icon(
                        painter = painterResource(id = R.drawable.nav_arrow_right),
                        contentDescription = "Start Course"
                    )
                }
                Text(text = "Приступить")
            }
        }
    }
}