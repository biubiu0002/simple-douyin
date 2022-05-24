package controller

import (
	"fmt"
	"simple-douyin/repository"
	"testing"
)

func TestVideo(t *testing.T) {
	repository.Init()
	data := Feed("1653226692123", "")
	fmt.Printf("%v", data)
}
