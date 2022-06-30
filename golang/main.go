package main

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/andskur/argon2-hashing"
    "fmt"
    "log"
)

func main() {
    runBcrypt()
    fmt.Println()
    runArgon2()
}

func runBcrypt() {
    fmt.Println("=== BCrypt ===")
    password := []byte("MyP@ssword123")

    // Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(hashedPassword))

    // Comparing the password with the hash
    err = bcrypt.CompareHashAndPassword(hashedPassword, password)

    if err == nil {
        fmt.Println("Password match!")
    } else {
        fmt.Println("Password does not match!")
        fmt.Println(err)
    }
}

func runArgon2() {
    fmt.Println("=== Argon2 ===")
    passwordFromForm := "MyP@ssword123"

    // Hashing the password
    hash, err := argon2.GenerateFromPassword([]byte(passwordFromForm), argon2.DefaultParams)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%s\n", hash)

    // Comparing the password with the hash
    err = argon2.CompareHashAndPassword(hash, []byte(passwordFromForm))
    if err == nil {
        fmt.Println("Password match!")
    } else {
        fmt.Println("Password does not match!")
        fmt.Println(err)
    }
}