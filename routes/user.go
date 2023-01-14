package routes

import "github.com/meme_hub/models"

type User struct{
	ID 			uint32    	`gorm:"primary_key;auto_increment" json:"id"`
	Nickname  	string    	`gorm:"size:255;not null;unique" json:"nickname"`
    Email     	string    	`gorm:"size:100;not null;unique" json:"email"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Nickname: userModel.Nickname, Email: userModel.Email}
}
