package controller

import (
	"fmt"
	"simple-douyin/service"
	"strconv"
	"time"
)

type VideoData struct {
	Response
	service.VideoInfoFlow
}

func Feed(latestTime, token string) *VideoData {
	tmp, err := strconv.ParseInt(latestTime, 10, 64)
	if err != nil {
		return &VideoData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("错误的时间戳: %v", tmp),
			},
		}
	}
	videoInfoFlow, err := service.Feed(time.Unix(tmp/1000, 0), token)
	if err != nil {
		return &VideoData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
	} else {
		return &VideoData{
			Response: Response{
				StatusCode: 0,
			},
			VideoInfoFlow: *videoInfoFlow,
		}
	}
}
