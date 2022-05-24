package repository

import (
	"fmt"
	"testing"
	"time"
)

func TestVideo(t *testing.T) {
	video := Video{
		UserId:     1,
		PlayUrl:    "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:   "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		Title:      "bear!",
		CreateTime: time.Now(),
	}
	_, err := AddVideo(&video)
	if err != nil {
		fmt.Printf("%v", err.Error())
		t.FailNow()
	}
	videos, err := SelectVideoListByLatest(time.Now())
	if err != nil {
		fmt.Printf("%v", err.Error())
		t.FailNow()
	}
	fmt.Printf("%v", videos)
}
