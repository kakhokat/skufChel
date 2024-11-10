// CourseCard.kt

package com.example.loginpage

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.colorResource
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController

@Composable
fun CourseCard(
    courseData: CourseCardData,
    navController: NavController,
    onLikeClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    Card(
        modifier = modifier
            .padding(8.dp)
            .fillMaxWidth()
            .clickable {
                navController.navigate("course/${courseData.id}")
            },
        shape = RoundedCornerShape(16.dp),
        elevation = CardDefaults.cardElevation(8.dp)
    ) {
        Column(
            modifier = Modifier
                .fillMaxWidth()
                .height(200.dp)
                .padding(24.dp)
        ) {
            // Course Name
            Text(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(bottom = 16.dp),
                text = courseData.courseName,
                style = MaterialTheme.typography.headlineSmall.copy(
                    fontSize = 20.sp,
                    fontWeight = FontWeight.Bold,
                    color = colorResource(id = R.color.logo_blue)
                ),
                maxLines = 1,
                overflow = TextOverflow.Ellipsis,
            )

            // Main Content: Rows and Columns
            Row(
                modifier = Modifier.fillMaxWidth(),
            ) {
                // Left Column: Profile Image, Likes Count, Like Button
                Column(
                    modifier = Modifier
                        .fillMaxHeight()
                        .fillMaxWidth(0.15f),
                    horizontalAlignment = Alignment.CenterHorizontally
                ) {
                    // Profile Image
                    Image(
                        painter = painterResource(id = courseData.profileImageRes),
                        contentDescription = "Profile Image",
                        modifier = Modifier
                            .size(40.dp)
                            .clip(RoundedCornerShape(50))
                    )

                    // Likes Count
                    Text(
                        text = courseData.likesCount.toString(),
                        style = MaterialTheme.typography.bodySmall
                    )

                    // Like Button
                    IconButton(
                        onClick = onLikeClick
                    ) {
                        Icon(
                            painter = painterResource(
                                id = if (courseData.isLiked)
                                    R.drawable.icom_liked
                                else
                                    R.drawable.icom_not_liked
                            ),
                            contentDescription = "Like Button",
                            tint = if (courseData.isLiked) Color.Red else Color.Gray
                        )
                    }
                }

                // Middle Column: Creator Name and Course Description
                Column(
                    modifier = Modifier
                        .fillMaxHeight()
                        .fillMaxWidth(0.8f)
                ) {
                    // Creator Name
                    Text(
                        text = "Создатель: ${courseData.creatorName}",
                        style = MaterialTheme.typography.bodyMedium,
                        maxLines = 1,
                        overflow = TextOverflow.Ellipsis
                    )

                    Spacer(modifier = Modifier.height(8.dp))

                    // Course Description
                    Text(
                        text = courseData.courseDescription,
                        style = MaterialTheme.typography.bodySmall,
                        maxLines = 2,
                        overflow = TextOverflow.Ellipsis
                    )
                }

                // Right Column: Lessons Icon, Lessons Count, Arrow Icon
                Column(
                    modifier = Modifier
                        .fillMaxHeight()
                        .fillMaxWidth()
                ) {
                    Column(
                        modifier = Modifier
                            .fillMaxHeight()
                            .fillMaxWidth(),
                        horizontalAlignment = Alignment.CenterHorizontally
                    ) {
                        // Clock Icon
                        Icon(
                            painter = painterResource(id = R.drawable.icom_clock),
                            contentDescription = "Lessons Icon",
                            modifier = Modifier.size(24.dp)
                        )

                        // Lessons Count
                        Text(
                            text = "${courseData.lessonsCount}",
                            style = MaterialTheme.typography.bodySmall,
                        )

                        Spacer(modifier = Modifier.height(24.dp))

                        // Right Arrow Icon
                        Icon(
                            painter = painterResource(id = R.drawable.icom_arrow_right),
                            contentDescription = "Arrow Icon",
                            modifier = Modifier.size(24.dp)
                        )
                    }
                }
            }
        }
    }
}