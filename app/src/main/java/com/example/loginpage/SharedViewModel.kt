// SharedViewModel.kt
package com.example.loginpage

import androidx.compose.runtime.mutableStateOf
import androidx.lifecycle.ViewModel

class SharedViewModel : ViewModel() {
    var profileCardData = mutableStateOf<ProfileCardData?>(null)
}
