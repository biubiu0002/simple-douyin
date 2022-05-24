package repository

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := Init()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
}

func TestAddUser(t *testing.T) {
	user := User{
		Username: "test",
		Password: "1234",
	}
	id, err := AddUser(&user)
	if err != nil {
		fmt.Printf("%v\n", err)
		t.FailNow()
	}
	fmt.Printf("user id: %v\n", id)
}

func TestSelectById(t *testing.T) {
	user, err := SelectById(1)
	if err != nil {
		fmt.Printf("%v\n", err)
		t.FailNow()
	}
	fmt.Printf("%v\n", user)
}
