/*
 * Copyright (c) 2019
 * Created at 5/29/19 2:41 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package role

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/domain/model"
	"github.com/oktopriima/mark-i/domain/repository"
)

type Usecase interface {
	Generate() (GenerateResponse, error)
}

type usecase struct {
	roleRepo repository.RoleRepository
	db       *gorm.DB
}

func NewUsecase(roleRepo repository.RoleRepository,
	db *gorm.DB) Usecase {
	return &usecase{roleRepo, db}
}

func (uc *usecase) findByValue(role string) (roles *model.Role, err error) {
	criteria := make(map[string]interface{})
	criteria["value"] = role

	roles, err = uc.roleRepo.FindOneBy(criteria)
	return
}

func (uc *usecase) create(roles *model.Role, tx *gorm.DB) (*model.Role, error) {
	role, err := uc.roleRepo.Create(roles, tx)
	if err != nil {
		return nil, err
	}
	return role, nil
}
