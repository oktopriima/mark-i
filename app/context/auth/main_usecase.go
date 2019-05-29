/*
 * Copyright (c) 2019
 * Created at 5/20/19 9:36 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/domain/model"
	"github.com/oktopriima/mark-i/domain/repository"
)

type Usecase interface {
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
type usecase struct {
	userRepo       repository.UserRepository
	userSocialRepo repository.UserSocialRepository
	db             *gorm.DB
}

func NewUsecase(userRepo repository.UserRepository,
	userSocialRepo repository.UserSocialRepository,
	db *gorm.DB) Usecase {
	return &usecase{
		userRepo,
		userSocialRepo,
		db,
	}
}

/** insert registered user to social media */
func (uc *usecase) createSocialMedia(
	users *model.Users,
	socialID string,
	request AuthRequest,
	tx *gorm.DB) (*model.UserSocial, error) {

	m := new(model.UserSocial)
	m.UsersID = users.ID
	m.SocialID = socialID
	m.SocialName = request.GetSocialType()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	if m, err := uc.userSocialRepo.Create(m, tx); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (uc *usecase) getUserSocial(socialID string) (m *model.UserSocial, err error) {
	criteria := make(map[string]interface{})
	criteria["social_id"] = socialID

	m, err = uc.userSocialRepo.FindOneBy(criteria)
	return
}

func (uc *usecase) getUserByEmail(email string) (m *model.Users, err error) {
	criteria := make(map[string]interface{})
	criteria["email"] = email

	m, err = uc.userRepo.FindOneBy(criteria)
	return
}

func (uc *usecase) updateLastLogin(users *model.Users) (err error) {
	users.LastLogin = time.Now()

	err = uc.userRepo.Update(users)
	return
}
