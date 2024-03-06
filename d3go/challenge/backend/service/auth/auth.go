package auth

import (
	"d3go/config"
	"d3go/db"
	"d3go/model"
)

const (
	UnAuthed int = iota - 1
	User
	Admin
)

func Login(u *model.User) (int, error) {
	ok, err := db.CheckAuth(u)
	if !ok || err != nil {
		return UnAuthed, err
	}
	if config.Conf.NoAdminLogin && u.ID == 1 {
		return UnAuthed, nil
	}
	if db.IsAdmin(u) {
		return Admin, nil
	}
	return User, nil
}

func Register(u *model.User) error {
	return db.AddUser(u)
}
