package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique_index"`
	FullName string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
}

func (user Users) GetAllUsers(where ...interface{}) []Users {
	var users []Users
	DB.Find(&users, where...)
	return users
}

func (user Users) GetUser(username string) Users {
	DB.Where("username =?", username).First(&user)
	return user
}

func (user Users) CreateUser() {
	DB.Create(&user)
}

func (user Users) UpdateUser(data Users) {
	DB.Model(&user).Updates(data)
}

func (user Users) DeleteUser(where ...interface{}) {
	DB.Unscoped().Delete(&user)
}

func (user Users) GetUserEmail(email string) bool {
	DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func (user Users) GetUserUsername(username string) bool {
	DB.Where("username =? ", username).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
