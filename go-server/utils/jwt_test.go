package utils

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	Generate(123)
}

func TestParse(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjI4NjUzNjAsInVzZXJJZCI6NH0.xOogbymyjYmveAsNqU8j63-fcZnqRhD2Z35_0uf1rz4"
	user, err := ParseToken(token)

	if err != nil {
		println("==============")
		println(err.Error())

	} else {
		fmt.Printf("%v", user.UserId)

	}

}
