/*
 * Copyright (c) 2019
 * Created at 5/29/19 2:29 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/domain/model"
)

type RoleRepository interface {
	Create(role *model.Role, tx *gorm.DB) (*model.Role, error)
	FindOneBy(criteria map[string]interface{}) (*model.Role, error)
	Find(ID int64) (*model.Role, error)
}
