package com.example.loginpage

import android.net.Uri
import android.widget.Toast
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.Image
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
import androidx.compose.ui.unit.dp
import androidx.navigation.NavController
import coil.compose.rememberImagePainter
import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.request.forms.*
import androidx.compose.ui.text.input.ImeAction
import io.ktor.client.statement.*
import io.ktor.http.*
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.sp
import java.io.InputStream

@Preview
@Composable
fun RegistrationScreenPreview() {
    RegistrationScreen(
        navController = null,
        sharedViewModel = SharedViewModel()
        // Provide a mocked or default instance
    )
}

@Composable
fun RegistrationScreen(navController: NavController? = null, sharedViewModel: SharedViewModel) {
    var email by remember { mutableStateOf("") }
    var username by remember { mutableStateOf("") }
    var password by remember { mutableStateOf("") }
    var selectedImageUri by remember { mutableStateOf<Uri?>(null) }
    val context = LocalContext.current
    val focusManager = LocalFocusManager.current

    val imageLauncher = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.GetContent(),
        onResult = { uri ->
            selectedImageUri = uri
        }
    )

    val httpClient = HttpClient()

    Column(
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp),
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
        Image(
            painter = painterResource(id = R.drawable.logo),
            contentDescription = "App Logo",
            modifier = Modifier.size(400.dp)
        )
        Spacer(modifier = Modifier.height(16.dp))
        // Поля ввода
        TextField(
            value = email,
            onValueChange = { email = it },
            label = { Text("Email") },
            keyboardOptions = KeyboardOptions(imeAction = ImeAction.Next),
            keyboardActions = KeyboardActions(
                onNext = { focusManager.moveFocus(FocusDirection.Down) }
            ),
            modifier = Modifier.fillMaxWidth()
        )

        Spacer(modifier = Modifier.height(8.dp))

        TextField(
            value = username,
            onValueChange = { username = it },
            label = { Text("Username") },
            keyboardOptions = KeyboardOptions(imeAction = ImeAction.Next),
            keyboardActions = KeyboardActions(
                onNext = { focusManager.moveFocus(FocusDirection.Down) }
            ),
            modifier = Modifier.fillMaxWidth()
        )

        Spacer(modifier = Modifier.height(8.dp))

        TextField(
            value = password,
            onValueChange = { password = it },
            label = { Text("Password") },
            keyboardOptions = KeyboardOptions(imeAction = ImeAction.Done),
            keyboardActions = KeyboardActions(onDone = { focusManager.clearFocus() }),
            modifier = Modifier.fillMaxWidth()
        )

        Spacer(modifier = Modifier.height(16.dp))

        // Кнопка выбора изображения
        Button(
            onClick = { imageLauncher.launch("image/*") },
            modifier = Modifier.fillMaxWidth()
        ) {
            Text("Выбрать фото")
        }

        Spacer(modifier = Modifier.height(8.dp))

        // Отображение выбранного изображения
        selectedImageUri?.let { uri ->
            Image(
                painter = rememberImagePainter(data = uri),
                contentDescription = "Selected Image",
                modifier = Modifier.size(200.dp)
            )
        }

        Spacer(modifier = Modifier.height(8.dp))

        // Кнопка регистрации
        Button(
            onClick = {
                CoroutineScope(Dispatchers.IO).launch {
                    try {
                        val imageData = selectedImageUri?.let { uri ->
                            context.contentResolver.openInputStream(uri)?.readBytes()
                        }

                        val response: HttpResponse = httpClient.submitFormWithBinaryData(
                            url = "http://10.0.2.2:8080/courses/auth/signup",
                            formData = formData {
                                append("email", email)
                                append("username", username)
                                append("password", password)
                                imageData?.let { data ->
                                    append(
                                        "profile_image",
                                        data,
                                        Headers.build {
                                            append(HttpHeaders.ContentType, "image/jpeg")
                                            append(HttpHeaders.ContentDisposition, "filename=\"profile.jpg\"")
                                        }
                                    )
                                }
                            }
                        )

                        withContext(Dispatchers.Main) {
                            if (response.status == HttpStatusCode.OK) {
                                val EmailData = EmailData(
                                    email = email
                                )
                                sharedViewModel.emailData.value = EmailData
                                Toast.makeText(context, "Регистрация успешна!", Toast.LENGTH_SHORT).show()
                                navController?.navigate("checkcode")
                            } else {
                                val errorText = response.bodyAsText()
                                Toast.makeText(context, "Ошибка: $errorText", Toast.LENGTH_SHORT).show()
                            }
                        }
                    } catch (e: Exception) {
                        withContext(Dispatchers.Main) {
                            Toast.makeText(context, "Ошибка: ${e.localizedMessage}", Toast.LENGTH_SHORT).show()
                        }
                    }
                }
            },
            modifier = Modifier.fillMaxWidth()
        ) {
            Text("Зарегистрироваться")
        }
    }
}
