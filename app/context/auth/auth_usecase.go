/*
 * Copyright (c) 2019
 * Created at 5/20/19 9:36 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/domain/repository"
)

type AuthUsecase interface {
	GoogleLogin(AuthRequest) (AuthResponse, error)
	FacebookLogin(AuthRequest) (AuthResponse, error)
	PhoneLogin(PhoneLoginRequest) (AuthResponse, error)
	EmailLogin(EmailLoginRequest) (AuthResponse, error)
}

type AuthResponse interface {
}

/*
 * request for social login
 */
type AuthRequest interface {
	GetSocialType() string
	GetSocialToken() string
}

/*
 * request for phone login
 */
type PhoneLoginRequest interface {
	GetPhone() string
	GetPin() string
}

/*
 * request for email login
 */
type EmailLoginRequest interface {
	GetEmail() string
	GetPin() string
}

/*
 * repository inject
 */
type authUsecase struct {
	userRepo       repository.UserRepository
	userSocialRepo repository.UserSocialRepository
	db             *gorm.DB
}

func NewLoginUsecase(userRepo repository.UserRepository,
	userSocialRepo repository.UserSocialRepository,
	db *gorm.DB) AuthUsecase {
	return &authUsecase{
		userRepo,
		userSocialRepo,
		db,
	}
}
