package mysql

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID       uint64 `gorm:"primaryKey"`
	Username     string
	Password     string
	NickName     string
	Follower_cnt int64
	Follow_cnt   int64
	RegisterTime time.Time
}

type Video struct {
	gorm.Model
	VideoID      uint64 `gorm:"primaryKey"`
	AutherID     uint64
	PlayUrl      string
	CoverUrl     string
	LikeCount    int64
	CommentCount int64
	Title        string
	Abstract     string
}

type Comment struct {
	gorm.Model
	CommentID   uint64 `gorm:"primaryKey"`
	VideoID     uint64
	UserID      uint64
	Context     string
	LikeCount   int64
	IsLike      bool
	CommentTime time.Time
}

type Like struct {
	gorm.Model
	KeyID    uint64 `gorm:"primaryKey"`
	OwnerID  uint64
	VideoID  uint64
	LikeTime time.Time
}

type Follow struct {
	gorm.Model
	KeyID     uint64 `gorm:"primaryKey"`
	WatcherID uint64
	WatchedID uint64
	WatchTime time.Time
}
