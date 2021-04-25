package pl.edu.uj.obj4.entity;

import javax.persistence.*

@Entity
class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long = 0
    @Column(unique = true)
    var email: String = ""
    var password: String = ""

    constructor(email: String, password: String) {
        this.email = email
        this.password = password
    }
}
