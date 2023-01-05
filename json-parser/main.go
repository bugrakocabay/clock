package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type permissions map[string]bool

type user struct {
	Name       string      `json:"username"`
	Password   string      `json:"-"`
	Permission permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"osman", "qwerty", nil},
		{"aise", "123456", permissions{"admin": false}},
		{"abuzer", "!str0n6p4s5", permissions{"read": true, "write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))

	decode()
}

type userParsed struct {
	Name       string      `json:"username"`
	Permission permissions `json:"perms"`
}

func decode() {
	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []userParsed

	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ ", user.Name)

		switch p := user.Permission; {
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["read"]:
			fmt.Print(" can read.")
		case p["write"]:
			fmt.Print(" can write.")
		}
	}
}
