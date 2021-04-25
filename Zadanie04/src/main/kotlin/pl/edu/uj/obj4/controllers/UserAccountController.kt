package pl.edu.uj.obj4.controllersfont

import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.stereotype.Controller
import org.springframework.ui.Model
import org.springframework.web.bind.annotation.*
import pl.edu.uj.obj4.service.UserAccountService

@Controller
@RequestMapping("/")
class UserAccountController {

    @Autowired
    lateinit var userAccountService: UserAccountService

    @PostMapping("register/")
    fun registerUser(@RequestBody email: String, @RequestBody password: String) {
        userAccountService.registerUser(email, password)
    }

    @GetMapping("register/")
    fun displayRegister(model: Model): String {
        return "register"
    }

    @PostMapping("login/")
    fun login(@RequestBody email: String, @RequestBody password: String): HttpStatus {
        userAccountService.login(email, password)
    }

}