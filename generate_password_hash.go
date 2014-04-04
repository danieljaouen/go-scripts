package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"crypto/rand"
	"crypto/sha512"
	"os"
	"strings"

	"code.google.com/p/gopass"
)

const (
	MagicPrefix   = "$6$"
	RandomSalt    = ""
	RoundsDefault = 5000
	RoundsMax     = 999999999
	RoundsMin     = 1000
	SaltLenMax    = 16
	SaltLenMin    = 1
)


func GenerateSalt(length, rounds int) string {
	if length > SaltLenMax {
		length = SaltLenMax
	} else if length < SaltLenMin {
		length = SaltLenMin
	}
	rlen := (length * 6 / 8)
	if (length * 6) % 8 != 0 {
		rlen += 1
	}
	if rounds < RoundsMin {
		rounds = RoundsMin
	} else if rounds > RoundsMax {
		rounds = RoundsMax
	}

	buf := make([]byte, rlen)
	rand.Read(buf)
	salt := Hash64(buf)
	if rounds == RoundsDefault {
		return fmt.Sprintf("%s%s", MagicPrefix, salt)
	}
	return fmt.Sprintf("%srounds=%d$%s", MagicPrefix, rounds, salt)
}


func Prompt() string {
	var password1 string
	var password2 string

	prompt1 := "Enter your desired password: "
	prompt2 := "Once more: "
	invalid := "Error, your passwords didn't match.  Please try again.\n"

	password1, _ = gopass.GetPass(prompt1)
	password2, _ = gopass.GetPass(prompt2)

	for password1 != password2 {
		fmt.Println(invalid)

		password1, _ = gopass.GetPass(prompt1)
		password2, _ = gopass.GetPass(prompt2)
	}

	return password1
}


func HashPassword(password string, rounds int) string {
	hash := sha512.New()
}


func main() {
	password := Prompt()
	hashedPassword := HashPassword(password, 5000)
	fmt.Println()
	fmt.Println("-------------------------------------")
	fmt.Println("-- Your password hash is:")
	fmt.Println(hashedPassword)
	fmt.Println("-------------------------------------")

	who := "World!"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)
}
