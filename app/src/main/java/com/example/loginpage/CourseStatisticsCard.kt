// CourseStatisticsCard.kt
package com.example.loginpage

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp

@Composable
fun CourseStatisticsCard(
    courseStatisticsData: CourseStatisticsData,
    modifier: Modifier = Modifier
) {
    Card(
        modifier = modifier
            .padding(8.dp)
            .fillMaxWidth(),
        shape = RoundedCornerShape(16.dp),
        elevation = CardDefaults.cardElevation(4.dp)
    ) {
        Row(
            modifier = Modifier
                .padding(16.dp)
        ) {
            // Left Column
            Column(
                modifier = Modifier.weight(1f),
                verticalArrangement = Arrangement.spacedBy(8.dp)
            ) {
                Text(text = "Записалось на курс:")
                Text(text = "Завершило курс:")
                Text(text = "Последнее обновление:")
            }
            // Right Column
            Column(
                modifier = Modifier.weight(1f),
                horizontalAlignment = Alignment.End,
                verticalArrangement = Arrangement.spacedBy(8.dp)
            ) {
                Text(text = courseStatisticsData.enrolledCount.toString())
                Text(text = courseStatisticsData.completedCount.toString())
                Text(text = courseStatisticsData.lastUpdated)
            }
        }
    }
}
