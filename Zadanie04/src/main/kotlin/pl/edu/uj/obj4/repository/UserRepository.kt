package pl.edu.uj.obj4.repository

import org.springframework.data.repository.CrudRepository
import org.springframework.stereotype.Repository
import pl.edu.uj.obj4.entity.User

@Repository
interface UserRepository : CrudRepository<User, Long> {
    fun existsByEmail(email:String): Boolean
    fun findByEmail(email:String): User?
}