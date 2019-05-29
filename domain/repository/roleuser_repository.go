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
	Create(user *model.RoleUser, tx *gorm.DB) (*model.RoleUser, error)
	Update(user *model.RoleUser, tx *gorm.DB) error
	Delete(user *model.RoleUser, tx *gorm.DB) error
	Find(ID int64) (*model.RoleUser, error)
	FindOneBy(criteria map[string]interface{}) (*model.RoleUser, error)
}
