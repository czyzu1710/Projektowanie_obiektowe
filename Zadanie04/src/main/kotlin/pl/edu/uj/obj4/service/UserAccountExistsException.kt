package pl.edu.uj.obj4.service

import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.ResponseStatus

@ResponseStatus(code = HttpStatus.CONFLICT, reason = "This email is already registered")
class UserAccountExistsException : RuntimeException()