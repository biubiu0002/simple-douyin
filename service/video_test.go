package service

import (
	"fmt"
	"testing"
	"time"
)

func TestFeed(t *testing.T) {
	fmt.Printf("%v", time.Unix(1653223495, 781))
	fmt.Printf("%v", time.Now().Unix())
	videoInfoFlow, err := Feed(time.Now(), "")
	if err != nil {
		fmt.Printf("get Feed Failed: %v", err.Error())
		t.FailNow()
	}
	fmt.Printf("%v", videoInfoFlow)
}
