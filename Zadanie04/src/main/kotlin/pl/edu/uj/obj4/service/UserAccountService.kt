package pl.edu.uj.obj4.service

import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import pl.edu.uj.obj4.entity.User
import pl.edu.uj.obj4.repository.UserRepository

@Service
class UserAccountService {
    @Autowired
    lateinit var userRepository: UserRepository

    fun registerUser(email: String, password: String)  {
        if (userRepository.existsByEmail(email)) {
            throw UserAccountExistsException()
        }

        userRepository.save(User(email, password))
    }

    fun login(email: String, password: String):Boolean {
        val user = userRepository.findByEmail(email)
        user?: throw NoSuchUserException()
        return user.password == password
    }

}