package repository

import (
	"time"
)

type Video struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	PlayUrl    string    `gorm:"column:play_url"`
	CoverUrl   string    `gorm:"column:cover_url"`
	Title      string    `gorm:"column:title"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func AddVideo(video *Video) (*Video, error) {
	err := db.Create(video).Error
	return video, err
}

func SelectVideoListByLatest(latestTime time.Time) ([]Video, error) {
	var videos []Video
	err := db.Where("create_time < ?", latestTime).Order("create_time desc").Find(&videos).Error
	return videos, err
}
