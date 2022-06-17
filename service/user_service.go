package service

import (
	"errors"
	"final-project/entity"
	"strings"
)

type UserIface interface {
	Register(user *entity.User) (*entity.User, error)
}

type UserSvc struct{}

func NewUserSvc() UserIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity.User) (*entity.User, error) {
	// validasi format email
	if !strings.Contains(user.Email, "gmail.com") {
		return nil, errors.New("masukkan format email yang benar")
	}
	// email tidak boleh kosong
	if user.Email == "" {
		return nil, errors.New("email tidak boleh kosong")
	}
	// validasi username
	if user.Username == "" {
		return nil, errors.New("username tidak boleh kosong")
	}
	// validasi password
	if user.Password == "" {
		return nil, errors.New("password tidak boleh kosong")
	}
	// validasi age
	if user.Age == 0 || user.Age <= 8 {
		return nil, errors.New("age tidak boleh kosong atau di bawah 8 tahun")
	}
	// jika validasi berhasil
	return user, nil
}
