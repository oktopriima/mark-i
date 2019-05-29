/*
 * Copyright (c) 2019
 * Created at 5/27/19 1:26 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package user

import "github.com/oktopriima/mark-i/domain/repository"

type Usecase interface {
	Find(ID int64) (FindResponse, error)
}

type usecase struct {
	userRepo repository.UserRepository
}

func NewUsecase(userRepo repository.UserRepository) Usecase {
	return &usecase{userRepo}
}
