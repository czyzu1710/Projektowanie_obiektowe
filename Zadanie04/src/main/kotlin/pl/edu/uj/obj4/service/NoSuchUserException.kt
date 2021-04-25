package pl.edu.uj.obj4.service

import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.ResponseStatus
import java.lang.RuntimeException
@ResponseStatus(code = HttpStatus.UNAUTHORIZED, reason = "No such user account")
class NoSuchUserException : RuntimeException()