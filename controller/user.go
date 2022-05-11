package controller

import (
	"simple-douyin/service"
	"strconv"
)

type UserData struct {
	Response
	Data interface{}
}

func Register(username string, password string) *UserData {
	loginInfo, err := service.Register(username, password)

	if err != nil {
		return &UserData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		}
	} else {
		return &UserData{
			Response: Response{
				StatusCode: 0,
			},
			Data: loginInfo,
		}
	}
}

func Login(username string, password string) *UserData {
	loginInfo, err := service.Login(username, password)

	if err != nil {
		return &UserData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		}
	} else {
		return &UserData{
			Response: Response{StatusCode: 0},
			Data:     loginInfo,
		}
	}
}

func UserInfo(userIdStr, token string) *UserData {
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return &UserData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "不合法的用户id"},
		}
	}
	userInfo, err := service.GetUserInfo(userId, token)
	if err != nil {
		return &UserData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		}
	} else {
		return &UserData{
			Response: Response{StatusCode: 0},
			Data:     userInfo,
		}
	}
}
