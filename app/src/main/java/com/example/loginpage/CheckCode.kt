package com.example.loginpage

import android.widget.Toast
import androidx.compose.foundation.Image
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.Button
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
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
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.navigation.NavController
import io.ktor.client.HttpClient
import io.ktor.client.request.forms.formData
import io.ktor.client.request.forms.submitFormWithBinaryData
import io.ktor.client.plugins.contentnegotiation.*
import io.ktor.client.plugins.logging.LogLevel
import io.ktor.client.plugins.logging.Logging
import io.ktor.serialization.kotlinx.json.*
import io.ktor.client.statement.HttpResponse
import io.ktor.client.statement.bodyAsText
import io.ktor.http.Headers
import io.ktor.http.HttpHeaders
import io.ktor.client.request.post
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.setBody
import io.ktor.http.ContentType
import io.ktor.serialization.kotlinx.json.json
import io.ktor.http.HttpStatusCode
import io.ktor.http.contentType
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

@Preview
@Composable
fun CheckCodePreview() {
    CheckCode(
        navController = null,
        sharedViewModel = SharedViewModel()
        // Provide a mocked or default instance
    )
}

@Composable
fun CheckCode(navController: NavController? = null, sharedViewModel: SharedViewModel) {
    var code by remember { mutableStateOf("") }
    var email by remember { mutableStateOf("") }
    val context = LocalContext.current
    val focusManager = LocalFocusManager.current
    val httpClient = HttpClient {
        install(ContentNegotiation) {
            json() // Подключаем сериализацию JSON
        }
        install(Logging) {
            level = LogLevel.ALL
        }
    }

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

        Spacer(modifier = Modifier.height(8.dp))

        // Поле ввода кода
        TextField(
            value = code,
            onValueChange = { code = it },
            label = { Text("Code") },
            keyboardOptions = KeyboardOptions(imeAction = ImeAction.Done),
            keyboardActions = KeyboardActions(
                onDone = { focusManager.clearFocus() }
            ),
            modifier = Modifier.fillMaxWidth()
        )

        Spacer(modifier = Modifier.height(16.dp))

        // Кнопка подтверждения
        Button(
            onClick = {
                CoroutineScope(Dispatchers.IO).launch {
                    try {
                        val response: HttpResponse = httpClient.post("http://10.0.2.2:8080/courses/auth/checkkey") {
                            contentType(io.ktor.http.ContentType.Application.Json)
                            setBody(mapOf("checkKey" to code, "mail" to sharedViewModel.emailData.value?.email))
                        }

                        withContext(Dispatchers.Main) {
                            if (response.status == HttpStatusCode.OK) {
                                Toast.makeText(context, "Код успешно проверен!", Toast.LENGTH_SHORT).show()
                                navController?.navigate("login")
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
            Text("Проверить код")
        }
    }
}
