/*
 * Copyright (c) 2019
 * Created at 5/28/19 12:54 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"errors"

	"github.com/oktopriima/mark-i/app/config"
	"github.com/oktopriima/mark-i/libraries/middleware"
	"golang.org/x/crypto/bcrypt"
)

func (uc *usecase) EmailLogin(request EmailLoginRequest) (AuthResponse, error) {
	user, err := uc.getUserByEmail(request.GetEmail())
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.GetPin())); err != nil {
		return nil, errors.New("email and password not match")
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
