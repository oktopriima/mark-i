/*
 * Copyright (c) 2019
 * Created at 5/29/19 5:19 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/domain/repository"
)

type Usecase interface {
	Create(request CreateRequest) (CreateResponse, error)
}

type usecase struct {
	roleUserRepo repository.RoleUserRepository
	db           *gorm.DB
}

func NewUsecase(roleUserRepo repository.RoleUserRepository, db *gorm.DB) Usecase {
	return &usecase{roleUserRepo, db}
}
