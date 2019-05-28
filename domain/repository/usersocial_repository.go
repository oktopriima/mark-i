/*
 * Copyright (c) 2019
 * Created at 5/26/19 11:14 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/domain/model"
)

type UserSocialRepository interface {
	FindOneBy(criteria map[string]interface{}) (*model.UserSocial, error)
	Create(social *model.UserSocial, tx *gorm.DB) (*model.UserSocial, error)
}
