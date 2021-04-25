package pl.edu.uj.obj4.service

import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.ResponseStatus

@ResponseStatus(code = HttpStatus.UNAUTHORIZED, reason = "Credentials do not much for given user")
class CredentialsDoNotMachException : RuntimeException()