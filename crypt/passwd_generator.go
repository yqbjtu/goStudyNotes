package main

import (
    "fmt"
    crypto_rand "crypto/rand"
	math_rand "math/rand"
	"math/big"
    "strings"
    "time"
)

var (
    lowerCharSet   = "abcdedfghijklmnopqrst"
    upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialCharSet = "!@#$%&*"
    numberSet      = "0123456789"
    allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
    math_rand.Seed(time.Now().Unix())
    minSpecialChar := 1
    minNum := 1
    minUpperCase := 1
    passwordLength := 8
    password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
    fmt.Println(password)

    minSpecialChar = 2
    minNum = 2
    minUpperCase = 2
    passwordLength = 20
    password = generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
    fmt.Println(password)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
    var password strings.Builder

    //Set special character
    for i := 0; i < minSpecialChar; i++ {
        //random := rand.Intn(len(specialCharSet))
		nBig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(len(specialCharSet))))
		if err != nil {
			panic(err)
		}
		random := nBig.Int64()

        password.WriteString(string(specialCharSet[random]))
    }

    //Set numeric
    for i := 0; i < minNum; i++ {
        //random := rand.Intn(len(numberSet))
		nBig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(len(numberSet))))
		if err != nil {
			panic(err)
		}
		random := nBig.Int64()
        password.WriteString(string(numberSet[random]))
    }

    //Set uppercase
    for i := 0; i < minUpperCase; i++ {
        //random := rand.Intn(len(upperCharSet))
		nBig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(len(upperCharSet))))
		if err != nil {
			panic(err)
		}
		random := nBig.Int64()
        password.WriteString(string(upperCharSet[random]))
    }

    remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
    for i := 0; i < remainingLength; i++ {
        //random := rand.Intn(len(allCharSet))
		nBig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(len(allCharSet))))
		if err != nil {
			panic(err)
		}
		random := nBig.Int64()
        password.WriteString(string(allCharSet[random]))
    }

    inRune := []rune(password.String())
	math_rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
