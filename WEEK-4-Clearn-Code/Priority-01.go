package main

import "fmt"

// user represents a user in the system.
type user struct {
    Email    string // Email of the user
    Password string // Password of the user
}


type userRepo struct {
    DB []user 
}


func (r *userRepo) Register(u user) {
    
    if u.Email == "" || u.Password == "" {
        fmt.Println("Registration failed: Email and Password cannot be empty")
        return 
    }

   
    r.DB = append(r.DB, u)
    fmt.Println("User registered successfully")
}


func (r *userRepo) Login(u user) string {
   
    if u.Email == "" || u.Password == "" {
        fmt.Println("Login failed: Email and Password cannot be empty")
        return "" 
    }

    for _, us := range r.DB {
        if us.Email == u.Email && us.Password == u.Password {
            return "auth token" 
        }
    }

    fmt.Println("Login failed: User not found or incorrect password")
    return "" 
}

func main() {
   
    repo := userRepo{}
    repo.Register(user{Email: "user@example.com", Password: "password123"})
    token := repo.Login(user{Email: "user@example.com", Password: "password123"})
    fmt.Println("Login token:", token)
}
