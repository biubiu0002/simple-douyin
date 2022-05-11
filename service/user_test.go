package service

import (
	"fmt"
	"simple-douyin/repository"
	"testing"
)

func TestMain(m *testing.M) {
	repository.Init()
	m.Run()
}

func TestRegister(t *testing.T) {
	user, err := Register("test", "1234")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		t.FailNow()
	}
	fmt.Printf("%v\n", user)
}

func TestLogin(t *testing.T) {
	user, err := Login("test", "1234")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		t.FailNow()
	}
	fmt.Printf("%v\n", user)
}

func TestUserInfo(t *testing.T) {
	user, err := GetUserInfo(5, "test1234 test1234")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		t.FailNow()
	}
	fmt.Printf("%v\n", user)
}
