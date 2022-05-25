package service

import (
	"errors"
	"fmt"
	"simple-douyin/repository"
	"simple-douyin/util"
	"time"
)

type LoginInfo struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfo struct {
	UserId        int64  `json:"user_id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

const (
	expiredSecond = 3600
)

func Register(username string, password string) (*LoginInfo, error) {
	//check param
	if len(username) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(username) > 32 {
		return nil, errors.New("用户名支持最长32个字符")
	}
	if len(password) == 0 {
		return nil, errors.New("密码不能为空")
	}
	if len(password) > 32 {
		return nil, errors.New("密码支持最长32个字符")
	}

	user, err := repository.SelectByName(username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("用户名已存在")
	}

	//prepare
	user, err = repository.AddUser(&repository.User{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	token, err := util.GenToken(user.Id, time.Now().Unix()+expiredSecond)
	if err != nil {
		return nil, err
	}

	return &LoginInfo{
		UserId: user.Id,
		Token:  token,
	}, nil
}

func Login(username string, password string) (*LoginInfo, error) {
	//check param
	if len(username) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(username) > 32 {
		return nil, errors.New("用户名支持最长32个字符")
	}
	if len(password) == 0 {
		return nil, errors.New("密码不能为空")
	}
	if len(password) > 32 {
		return nil, errors.New("密码支持最长32个字符")
	}
	user, err := repository.SelectByName(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户名不存在")
	}
	if user.Password != password {
		return nil, errors.New("密码不正确")
	}
	token, err := util.GenToken(user.Id, time.Now().Unix()+expiredSecond)
	if err != nil {
		return nil, err
	}

	return &LoginInfo{
		UserId: user.Id,
		Token:  token,
	}, nil
}

func GetUserInfo(userId int64, token string) (*UserInfo, error) {
	if userId < 0 {
		return nil, fmt.Errorf("不合法的用户id: %v", userId)
	}

	//varify token
	if token != "" {
		_, err := util.ParseToken(token)
		if err != nil {
			return nil, err
		}
	}

	user, err := repository.SelectById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("用户id不存在: %v", userId)
	}

	return &UserInfo{
		UserId:        user.Id,
		Name:          user.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}, nil
}
