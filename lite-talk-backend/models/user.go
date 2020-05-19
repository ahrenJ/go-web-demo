package models

import "time"

type User struct {
	Id         int
	Nickname   string
	EncryptPwd string
	Avatar     string
	CreateTime time.Time
	UpdateTime time.Time
}

func AddUser(nickname string, password string) {
	var user = User{Nickname: nickname, EncryptPwd: password, CreateTime: time.Now(), UpdateTime: time.Now()}
	db.Create(&user)
}

func QueryUserById(id int) (user User, existed bool) {
	existed = !db.Where("id = ?", id).First(&user).RecordNotFound()
	return user, existed
}

func QueryUserExisted(nickname string) bool {
	return !db.Where("nickname = ?", nickname).First(new(User)).RecordNotFound()
}

func QueryUserByNickname(nickname string) (user User, existed bool) {
	existed = !db.Where("nickname = ?", nickname).First(&user).RecordNotFound()
	return user, existed
}
