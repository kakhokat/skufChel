// MainActivity.kt
package com.example.loginpage

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.lifecycle.viewmodel.compose.viewModel
import com.example.loginpage.ui.theme.LoginPageTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            LoginPageTheme {
                val sharedViewModel: SharedViewModel = viewModel()
                Surface(color = MaterialTheme.colorScheme.background) {
                    AppNavigation(sharedViewModel)
                }
            }
        }
    }
}
