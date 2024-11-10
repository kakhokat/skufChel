// ProfileCard.kt
package com.example.loginpage

import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.foundation.border

@Composable
fun ProfileCard(
    profileCardData: ProfileCardData,
    onExitClick: () -> Unit,
    modifier: Modifier = Modifier
) {
    Card(
        modifier = modifier
            .padding(8.dp)
            .fillMaxWidth()
            .border(
                width = 5.dp, // Adjust the width as needed
                color = Color.Blue,
                shape = RoundedCornerShape(16.dp)
            ),
        shape = RoundedCornerShape(16.dp),
    ) {
        Column(
            modifier = Modifier
                .padding(20.dp)
                .fillMaxWidth()
                .height(170.dp)
        ) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                verticalAlignment = Alignment.CenterVertically,
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                Image(
                    painter = painterResource(id = profileCardData.profileImageRes),
                    contentDescription = "Profile Image",
                    modifier = Modifier
                        .size(50.dp)
                        .clip(RoundedCornerShape(50))
                )

                Text(
                    text = profileCardData.nickName,
                    style = MaterialTheme.typography.headlineSmall.copy(
                        fontSize = 20.sp,
                        color = Color.Black
                    ),
                    modifier = Modifier
                        .weight(1f)
                        .padding(start = 16.dp)
                )

                IconButton(
                    onClick = onExitClick
                ) {
                    Icon(
                        painter = painterResource(id = R.drawable.exit),
                        contentDescription = "Exit Button",
                        tint = Color.Blue,
                        modifier = Modifier.size(26.dp)
                    )
                }
            }

            Spacer(modifier = Modifier.height(16.dp))

            Row(
                modifier = Modifier.fillMaxWidth()
            ) {
                Column(
                    modifier = Modifier
                        .fillMaxHeight()
                        .weight(2f),
                    verticalArrangement = Arrangement.spacedBy(15.dp),
                    horizontalAlignment = Alignment.Start
                ) {
                    Text(
                        text = "Пройдено курсов",
                        style = MaterialTheme.typography.bodyMedium
                    )
                    Text(
                        text = "Дней без перерыва",
                        style = MaterialTheme.typography.bodyMedium
                    )
                    Text(
                        text = "Дней без перерыва (макс.)",
                        style = MaterialTheme.typography.bodyMedium
                    )
                }

                Column(
                    modifier = Modifier
                        .fillMaxHeight()
                        .weight(1f),
                    verticalArrangement = Arrangement.spacedBy(15.dp),
                    horizontalAlignment = Alignment.End
                ) {
                    Text(
                        text = profileCardData.coursesCompleted.toString(),
                        style = MaterialTheme.typography.bodyMedium
                    )
                    Text(
                        text = profileCardData.daysStreak.toString(),
                        style = MaterialTheme.typography.bodyMedium
                    )
                    Text(
                        text = profileCardData.maxDaysStreak.toString(),
                        style = MaterialTheme.typography.bodyMedium
                    )
                }
            }
        }
    }
}

@Preview(showBackground = true)
@Composable
fun ProfileCardPreview() {
    ProfileCard(
        profileCardData = ProfileCardData(
            nickName = "Иван Иванов",
            profileImageRes = R.drawable.logo,
            coursesCompleted = 5,
            daysStreak = 7,
            maxDaysStreak = 10
        ),
        onExitClick = { /* Handle exit */ }
    )
}
