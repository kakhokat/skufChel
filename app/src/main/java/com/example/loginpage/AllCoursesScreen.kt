// AllCoursesScreen.kt
package com.example.loginpage

// Import for property delegation with 'by'
import android.widget.Toast
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
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Search
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults.buttonColors
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.input.pointer.pointerInput
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.input.TextFieldValue
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController

@Composable
fun AllCoursesScreen(
    navController: NavController,
    courseCards: List<CourseCardData>
) {
    val context = LocalContext.current

    // State variables
    var searchText by remember { mutableStateOf(TextFieldValue("")) }
    var activeSort by remember { mutableStateOf(SortOption.NEW) }

    Box(
        modifier = Modifier
            .fillMaxSize()
            .pointerInput(Unit) {
                detectHorizontalDragGestures { change, dragAmount ->
                    change.consume()
                    if (dragAmount > 50) {
                        // Swipe right detected
                        navController.navigate("profile") {
                            popUpTo("profile") { inclusive = true }
                        }
                    }
                }
            }
    ) {
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(horizontal = (44 * 0.5f).dp)
        ) {
            Spacer(modifier = Modifier.height(16.dp))

            // Search bar and icon
            Row(
                verticalAlignment = Alignment.CenterVertically,
                modifier = Modifier.fillMaxWidth()
            ) {
                // Search field
                TextField(
                    value = searchText,
                    onValueChange = { searchText = it },
                    placeholder = { Text(text = "Поиск...") },
                    modifier = Modifier
                        .weight(1f)
                )

                Spacer(modifier = Modifier.width(20.dp))

                // Search icon
                IconButton(
                    onClick = {
                        Toast.makeText(context, "Поиск: ${searchText.text}", Toast.LENGTH_SHORT).show()
                    }
                ) {
                    Icon(
                        imageVector = Icons.Default.Search,
                        contentDescription = "Search Icon"
                    )
                }
            }

            Spacer(modifier = Modifier.height(16.dp))

            // Sort buttons
            Row(
                modifier = Modifier
                    .fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                SortOption.values().forEach { sortOption ->
                    Button(
                        onClick = {
                            activeSort = sortOption
                            Toast.makeText(context, "Сортировка: ${sortOption.description}", Toast.LENGTH_SHORT).show()
                        },
                        colors = buttonColors(
                            containerColor = if (activeSort == sortOption) Color.Blue else Color.Gray
                        ),
                        modifier = Modifier.weight(1f).padding(horizontal = 4.dp)
                    ) {
                        Text(text = sortOption.title)
                    }
                }
            }

            Spacer(modifier = Modifier.height(16.dp))

            // Courses list
            Column(
                modifier = Modifier
                    .fillMaxWidth()
                    .weight(1f)
                    .verticalScroll(rememberScrollState())
            ) {
                courseCards.forEach { courseCard ->
                    CourseCard(
                        courseData = courseCard,
                        navController = navController,
                        onLikeClick = { /* Handle like action */ }
                    )
                }
            }
        }

        // Navigation button overlaid on top of the content
        IconButton(
            onClick = {
                navController.navigate("profile") {
                    popUpTo("profile") { inclusive = true }
                }
            },
            modifier = Modifier
                .align(Alignment.BottomStart)
                .padding(16.dp)
        ) {
            Column(
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                Icon(
                    painter = painterResource(id = R.drawable.nav_arrow_left),
                    contentDescription = "Back to Profile",
                    modifier = Modifier.size(48.dp)
                )
                Text(
                    text = "Профиль",
                    fontSize = 20.sp,
                    color = MaterialTheme.colorScheme.onSurface
                )
            }
        }
    }
}

// Definition of SortOption enum
enum class SortOption(val title: String, val description: String) {
    NEW("Новые", "Сортировка по дате создания"),
    POPULAR("Популярные", "Сортировка по количеству записавшихся"),
    BEST("Лучшие", "Сортировка по соотношению лайков и прошедших курс")
}
