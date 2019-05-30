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

	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/app/config"
	"github.com/oktopriima/mark-i/domain/model"
	"github.com/oktopriima/mark-i/libraries/helper"
	"github.com/oktopriima/mark-i/libraries/middleware"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) GoogleLogin(request AuthRequest) (AuthResponse, error) {
	/** get google data */
	gResp, err := helper.GetGoogleData(request.GetSocialToken())
	if err != nil {
		return nil, err
	}

	/** get users */
	user, err := uc.getUserByEmail(gResp.Emails[0].Value)
	if err != nil || user.ID == 0 {
		/** register users */
		tx := uc.db.Begin()
		user, err = uc.registerFromGoogle(*gResp, tx);
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		_, err = uc.createSocialMedia(user, gResp.ID, request, tx);
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
	}

	/** get user social */
	social, err := uc.getUserSocial(gResp.ID)
	if err != nil {
		return nil, err
	}

	if social.UsersID != user.ID {
		return nil, errors.New("user didn't match")
	}

	var role string
	roles, err := uc.findRoleUser(user)
	if err != nil {
		role = config.CONSUMER
	} else {
		role = roles.Value
	}

	/** generate parameter for Custom JWT */
	param := middleware.TokenStructure{}
	param.UserID = user.ID
	param.Email = user.Email
	param.Role = role

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

func (uc *usecase) registerFromGoogle(gResp helper.GoogleResponse, tx *gorm.DB) (*model.Users, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(helper.RandString(10)), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	m := new(model.Users)

	m.Name = gResp.Name.FamilyName
	m.Username = gResp.Emails[0].Value
	m.Email = gResp.Emails[0].Value
	m.Phone = ""
	m.Password = string(pass)
	m.LastLogin = time.Now()
	m.SecondaryEmail = gResp.Emails[0].Value
	m.IsVerified = false
	m.Avatar = gResp.Image.URL

	if m, err := uc.userRepo.Create(m, tx); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}
