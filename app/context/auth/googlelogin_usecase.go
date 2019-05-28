/*
 * Copyright (c) 2019
 * Created at 5/24/19 1:40 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"errors"
	"time"

	"github.com/oktopriima/mark-i/app/config"
	"github.com/oktopriima/mark-i/domain/model"
	"github.com/oktopriima/mark-i/libraries/helper"
	"github.com/oktopriima/mark-i/libraries/middleware"
	"golang.org/x/crypto/bcrypt"
)

func (uc *authUsecase) GoogleLogin(request AuthRequest) (AuthResponse, error) {
	/** get google data */
	gResp, err := helper.GetGoogleData(request.GetSocialToken())
	if err != nil {
		return nil, err
	}

	/** get users */
	user, err := uc.getUserByEmail(gResp.Emails[0].Value)
	if err != nil || user.ID == 0 {
		/** register users */
		if user, err = uc.registerFromGoogle(*gResp); err != nil {
			return nil, err
		}

		if _, err = uc.registerSocialMedia(user, gResp.ID, request); err != nil {
			return nil, err
		}

	}

	/** get user social */
	social, err := uc.getUserSocial(gResp.ID)
	if err != nil {
		return nil, err
	}

	if social.UserID != user.ID {
		return nil, errors.New("user didn't match")
	}

	/** generate parameter for Custom JWT */
	param := middleware.TokenStructure{}
	param.UserID = user.ID
	param.Email = user.Email
	param.Role = "consumer"

	auth := middleware.NewCustomAuth([]byte(config.SIGNATURE))
	token, err := auth.GenerateToken(param)
	if err != nil {
		return nil, err
	}

	if err = uc.updateLastLogin(user); err != nil {
		return nil, err
	}

	return token, nil
}

func (uc *authUsecase) registerFromGoogle(gResp helper.GoogleResponse) (*model.Users, error) {
	pass, _ := bcrypt.GenerateFromPassword([]byte(helper.RandString(10)), bcrypt.DefaultCost)

	m := new(model.Users)

	m.Name = gResp.Name.FamilyName
	m.AutogeneratedName = true
	m.Email = gResp.Emails[0].Value
	m.Password = string(pass)
	m.IsResetPassword = false
	m.Phone = ""
	m.IsVerified = false
	m.Gcm = ""
	m.RememberToken = new(string)
	m.LastLogin = time.Now()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	m.SecondaryEmail = gResp.Emails[0].Value
	m.Avatar = gResp.Image.URL
	m.DeviceType = ""
	m.HasPassword = true

	if m, err := uc.userRepo.Create(m); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (uc *authUsecase) registerSocialMedia(
	users *model.Users,
	socialID string,
	request AuthRequest) (*model.UserSocial, error) {

	m := new(model.UserSocial)
	m.UserID = users.ID
	m.SocialID = socialID
	m.SocialMedia = request.GetSocialType()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	if m, err := uc.userSocialRepo.Create(m); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}

func (uc *authUsecase) getUserSocial(socialID string) (m *model.UserSocial, err error) {
	criteria := make(map[string]interface{})
	criteria["social_id"] = socialID

	m, err = uc.userSocialRepo.FindOneBy(criteria)
	return
}

func (uc *authUsecase) getUserByEmail(email string) (m *model.Users, err error) {
	criteria := make(map[string]interface{})
	criteria["email"] = email

	m, err = uc.userRepo.FindOneBy(criteria)
	return
}

func (uc *authUsecase) updateLastLogin(users *model.Users) (err error) {
	users.LastLogin = time.Now()

	err = uc.userRepo.Update(users)
	return
}
