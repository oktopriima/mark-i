/*
 * Copyright (c) 2019
 * Created at 5/29/19 5:23 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/domain/model"
)

type RoleUserRepository interface {
	Create(user *model.UserRole, tx *gorm.DB) (*model.UserRole, error)
	Update(user *model.UserRole, tx *gorm.DB) error
	Delete(user *model.UserRole, tx *gorm.DB) error
	Find(ID int64) (*model.UserRole, error)
	FindOneBy(criteria map[string]interface{}) (*model.UserRole, error)
}
