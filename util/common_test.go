package util

import (
	"fmt"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	tokenString, err := GenToken(1, time.Now().Unix())
	if err != nil {
		fmt.Printf("generate token faild: %v\n", err.Error())
		t.FailNow()
	}
	fmt.Printf("%v\n", tokenString)
	userId, err := ParseToken(tokenString)
	if err != nil {
		fmt.Printf("parse token faild: %v\n", err.Error())
		t.FailNow()
	}
	fmt.Printf("%v\n", userId)
}
