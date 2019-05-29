/*
 * Copyright (c) 2019
 * Created at 5/28/19 12:09 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"errors"

	"github.com/oktopriima/mark-i/app/config"
	"github.com/oktopriima/mark-i/domain/model"
	"github.com/oktopriima/mark-i/libraries/helper"
	"github.com/oktopriima/mark-i/libraries/middleware"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) PhoneLogin(request PhoneLoginRequest) (AuthResponse, error) {

	phone := helper.NewPhoneNumber(request.GetPhone())
	user, err := uc.getUserByPhone(phone.GetPhoneNumber())

	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.GetPin())); err != nil {
		return nil, errors.New("phone and password not match")
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

func (uc *usecase) getUserByPhone(phone string) (*model.Users, error) {
	criteria := make(map[string]interface{})
	criteria["phone"] = phone
	criteria["is_verified"] = true

	user, err := uc.userRepo.FindOneBy(criteria)
	if err != nil {
		return nil, err
	}

	return user, nil
}
