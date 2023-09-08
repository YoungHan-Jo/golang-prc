package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err *PasswordError) Error() string {
	return "it's too short to set password"
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return &PasswordError{len(password), 8}
	}

	return nil
}

func main() {
	err := RegisterAccount("myID", "pass")
	if err != nil {
		if errInfo, ok := err.(*PasswordError); ok {
			fmt.Printf("len : %d, required len: %d", errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("join successful")
	}
}
