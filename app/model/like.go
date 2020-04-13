package model

import (
	"time"
)

type Likes struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
	//Users []User `gorm:"foreignkey:LID;AssociationForeignKey:ID"`
	Users []User `gorm:"foreignkey:LID;AssociationForeignKey:ID" json:"users"`
}

type UserLikesProfile struct {
	Name   string
	UserID uint
}


func GetLikes(PageNum int, PageSize int) (likes Likes) {
	//db.Offset(PageNum).Limit(PageSize).Find(&like)
	Eloquent.First(&likes)
	Eloquent.Model(&likes).Related(&likes.Users, "users")
	//db.Offset(PageNum).Limit(PageSize).Find(&like)
	return
}

