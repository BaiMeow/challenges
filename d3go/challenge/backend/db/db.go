package db

import (
	"d3go/config"
	"d3go/model"
	"encoding/hex"
	"errors"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ErrDatabase = errors.New("database error")

var db *gorm.DB

func Init() {
	if err := tryOpen(); err != nil {
		log.Fatalln(err)
	}
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalln(err)
	}

	// create admin
	rand.Seed(time.Now().UnixMicro())
	var rad [32]byte
	rand.Read(rad[:])
	if ok, _ := IsFirstRegistered(); ok {
		db.Save(&model.User{
			Username: "admin",
			Password: hex.EncodeToString(rad[:]),
		})
	}
}

func tryOpen() error {
	var err error
	var database *gorm.DB
	for i := 0; i < 100; i++ {
		database, err = gorm.Open(mysql.Open(config.Conf.DBUser+":"+config.Conf.DBPasswd+"@tcp("+config.Conf.DBHost+":"+config.Conf.DBPort+")/db?parseTime=True"), &gorm.Config{})
		if err != nil {
			time.Sleep(time.Second * 3)
			continue
		}
		db = database
		return nil
	}
	return err
}

func IsAdmin(u *model.User) bool {
	var admin model.User
	if err := db.First(&admin).Error; err != nil {
		log.Error(err)
	}
	return u.Username == admin.Username
}

func AddUser(u *model.User) error {
	if err := db.Save(u).Error; err != nil {
		log.WithField("user", u).Error(err)
		return ErrDatabase
	}
	return nil
}

func CheckAuth(u *model.User) (bool, error) {
	if err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		log.WithField("user", u).Error(err)
		return false, ErrDatabase
	}
	return true, nil
}

func IsFirstRegistered() (bool, error) {
	if err := db.First(&model.User{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		log.WithField("user", model.User{}).Error(err)
		return false, ErrDatabase
	}
	return false, nil
}
