package mysql

import (
	"time"
)

type User struct {
	UserID       int64 `gorm:"primaryKey"`
	Username     string
	Password     string
	NickName     string
	Follower_cnt int64
	Follow_cnt   int64
	RegisterTime time.Time
}

type Video struct {
	VideoID      int64 `gorm:"primaryKey"`
	AutherID     int64
	PlayUrl      string
	CoverUrl     string
	LikeCount    int64
	CommentCount int64
	Title        string
	Abstract     string
}

type Comment struct {
	CommentID   int64 `gorm:"primaryKey"`
	VideoID     int64
	UserID      int64
	Context     string
	LikeCount   int64
	IsLike      bool
	CommentTime time.Time
}

type Like struct {
	KeyID    int64 `gorm:"primaryKey"`
	OwnerID  int64
	VideoID  int64
	LikeTime time.Time
}

type Follow struct {
	KeyID     int64 `gorm:"primaryKey"`
	WatcherID int64
	WatchedID int64
	WatchTime time.Time
}
