// LessonCard.kt
package com.example.loginpage

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.CheckCircle
import androidx.compose.material.icons.filled.Favorite
import androidx.compose.material.icons.filled.FavoriteBorder
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

@Composable
fun LessonCard(
    lessonData: LessonData,
    onLikeClick: (LessonData) -> Unit,
    modifier: Modifier = Modifier
) {
    // Local state to update the UI when properties change
    var isLiked by remember { mutableStateOf(lessonData.isLiked) }

    Card(
        modifier = modifier
            .fillMaxWidth()
            .padding(8.dp),
        shape = MaterialTheme.shapes.medium,
        elevation = CardDefaults.cardElevation(4.dp)
    ) {
        Row(
            modifier = Modifier
                .padding(16.dp)
                .fillMaxWidth(),
            verticalAlignment = Alignment.CenterVertically
        ) {
            // Left side: Lesson title and description
            Column(
                modifier = Modifier.weight(1f)
            ) {
                Text(
                    text = lessonData.title,
                    style = MaterialTheme.typography.titleMedium
                )
                Spacer(modifier = Modifier.height(4.dp))
                Text(
                    text = lessonData.description,
                    style = MaterialTheme.typography.bodyMedium
                )
            }

            // Right side: Like button and Checkmark icon (if completed)
            Row(
                verticalAlignment = Alignment.CenterVertically
            ) {
                // Like Button


                // Checkmark Icon (displayed only if isCompleted == true)
                if (lessonData.isCompleted) {
                    Icon(
                        imageVector = Icons.Default.CheckCircle,
                        contentDescription = "Completed",
                        tint = MaterialTheme.colorScheme.primary,
                        modifier = Modifier.size(24.dp)
                    )
                    Spacer(modifier = Modifier.width(8.dp)) // Optional: Add space between icons
                }
                IconButton(
                    onClick = {
                        isLiked = !isLiked
                        onLikeClick(lessonData.copy(isLiked = isLiked))
                    }
                ) {
                    Icon(
                        imageVector = if (isLiked) Icons.Default.Favorite else Icons.Default.FavoriteBorder,
                        contentDescription = if (isLiked) "Liked" else "Not Liked",
                        tint = if (isLiked) MaterialTheme.colorScheme.error else MaterialTheme.colorScheme.onSurfaceVariant
                    )
                }
            }
        }
    }
}
