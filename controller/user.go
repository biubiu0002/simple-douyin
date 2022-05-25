package controller

import (
	"fmt"
	"simple-douyin/service"
	"strconv"
)

type LoginData struct {
	Response
	service.LoginInfo
}

type UserData struct {
	Response
	service.UserInfo
}

func Register(username string, password string) *LoginData {
	loginInfo, err := service.Register(username, password)

	if err != nil {
		return &LoginData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		}
	} else {
		return &LoginData{
			Response: Response{
				StatusCode: 0,
			},
			LoginInfo: *loginInfo,
		}
	}
}

func Login(username string, password string) *LoginData {
	loginInfo, err := service.Login(username, password)

	if err != nil {
		return &LoginData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		}
	} else {
		return &LoginData{
			Response:  Response{StatusCode: 0},
			LoginInfo: *loginInfo,
		}
	}
}

func UserInfo(userIdStr, token string) *UserData {
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return &UserData{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("不合法的用户id: %v", userIdStr),
			},
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
			UserInfo: *userInfo,
		}
	}
}
