/*
 * Copyright (c) 2019
 * Created at 5/27/19 1:26 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package user

import "github.com/oktopriima/mark-i/domain/repository"

type UserUsecase interface {
	Find(ID int64) (FindResponse, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}
