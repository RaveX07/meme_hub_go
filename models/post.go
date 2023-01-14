package models

import "time"

type Post struct{
	ID 			uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	UserID 		uint32   	`json:"userID"`
	Content 	string		`gorm:"size:240;not null" json:"content"`
	Image 		[]byte    	`gorm:"size:254;" json:"image"`
	Likes 		int32		`gorm:"default:0" json:"likes"`
	CreatedAt 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}