// LoginScreen.kt
package com.example.loginpage

import android.widget.Toast
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.FocusDirection
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalFocusManager
import androidx.compose.ui.res.colorResource
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController

@Composable
fun LoginScreen(
    navController: NavController? = null,
    sharedViewModel: SharedViewModel
) {
    var username by remember { mutableStateOf("") }
    var password by remember { mutableStateOf("") }
    val focusManager = LocalFocusManager.current
    val context = LocalContext.current


    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(5.dp),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        // Текст логотипа
        Text(
            text = stringResource(id = R.string.logo),
            color = colorResource(id = R.color.logo_blue),
            fontSize = 68.sp,
            fontWeight = FontWeight.Bold
        )

        Spacer(modifier = Modifier.height(8.dp))

        // Изображение логотипа
        Image(
            painter = painterResource(id = R.drawable.logo),
            contentDescription = "App Logo",
            modifier = Modifier.size(400.dp)
        )

        Spacer(modifier = Modifier.height(8.dp))

        // Вложенная колонка для кнопок и полей ввода
        Column(
            modifier = Modifier
                .padding(start = 25.dp, end = 25.dp),
            verticalArrangement = Arrangement.Center,
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            // Кнопки "Регистрация" и "Забыли пароль?"
            Row(
                modifier = Modifier
                    .height(40.dp)
                    .fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                Button(
                    modifier = Modifier
                        .weight(1f)
                        .padding(end = 8.dp),
                    onClick = { /* Действие для регистрации */ },
                    enabled = false
                ) {
                    Text("Регистрация")
                }

                Button(
                    modifier = Modifier
                        .weight(1f)
                        .padding(start = 8.dp),
                    onClick = { /* Действие для забытого пароля */ },
                    enabled = false
                ) {
                    Text("Забыли пароль?")
                }
            }

            Spacer(modifier = Modifier.height(8.dp))

            // Поле ввода логина
            TextField(
                value = username,
                onValueChange = { username = it },
                label = { Text("Username") },
                singleLine = true,
                keyboardOptions = KeyboardOptions(
                    imeAction = ImeAction.Next  // Настройка действия "Next" для клавиши Enter
                ),
                keyboardActions = KeyboardActions(
                    onNext = { focusManager.moveFocus(FocusDirection.Down) }  // Перемещаем фокус на следующее поле
                ),
                modifier = Modifier
                    .fillMaxWidth()
                    .background(color = colorResource(id = R.color.background))
            )

            Spacer(modifier = Modifier.height(8.dp))

            // Поле ввода пароля
            TextField(
                value = password,
                onValueChange = { password = it },
                label = { Text("Password") },
                visualTransformation = PasswordVisualTransformation(),
                singleLine = true,
                keyboardOptions = KeyboardOptions(
                    imeAction = ImeAction.Done  // Настройка действия "Done" для клавиши Enter
                ),
                keyboardActions = KeyboardActions(
                    onDone = {
                        focusManager.clearFocus()  // Скрываем клавиатуру после нажатия "Done"
                    }
                ),
                modifier = Modifier.fillMaxWidth()
            )

            Spacer(modifier = Modifier.height(16.dp))

            // Кнопка входа
            Button(
                onClick = {
                    focusManager.clearFocus()
                    if (username.isNotEmpty() && password.isNotEmpty()) {
                        val profileCardData = ProfileCardData(
                            nickName = username,
                            profileImageRes = R.drawable.logo,
                            coursesCompleted = 0,
                            daysStreak = 0,
                            maxDaysStreak = 0
                        )
                        sharedViewModel.profileCardData.value = profileCardData
                        navController?.navigate("profile")
                    } else {
                        Toast.makeText(
                            context,
                            "Заполните поля логин и пароль",
                            Toast.LENGTH_SHORT
                        ).show()
                    }
                },
                modifier = Modifier.fillMaxWidth()
            ) {
                Text("Вход")
            }
        }
    }
}