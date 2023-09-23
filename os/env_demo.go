package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Shell:", os.Getenv("SHELL"))

	getEnv := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	getEnv("EDITOR")
	getEnv("SHELL")

	fmt.Println("before editor:", os.Getenv("EDITOR"))
	os.Setenv("EDITOR", "emacs")
	fmt.Println("after setenv, editor:", os.Getenv("EDITOR"))

	// os.Environ returns a copy of strings representing the environment, in the form "key=value".
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s: %s\n", pair[0], pair[1])
	}

	// os.ExpandEnv is a helper function which replaces the $var inside a string into the value of the given variable.
	//	References to undefined variables are replaced by the empty string
	os.Setenv("EDITOR", "emacs")

	fmt.Println(os.ExpandEnv("My editor is $EDITOR."))
	fmt.Println(os.ExpandEnv("My shell is $SHELL."))

	// github.com/joho/godotenv  loads environment variables from the .env file
}
