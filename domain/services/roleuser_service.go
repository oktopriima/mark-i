/*
 * Copyright (c) 2019
 * Created at 5/29/19 5:23 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package services

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-i/domain/model"
	"github.com/oktopriima/mark-i/domain/repository"
)

type roleUserService struct {
	db *gorm.DB
}

func NewRoleUserService(db *gorm.DB) repository.RoleUserRepository {
	return &roleUserService{db}
}

func (srv *roleUserService) Create(user *model.RoleUser, tx *gorm.DB) (m *model.RoleUser, err error) {
	db := tx.Create(&user)

	if err = db.Error; err != nil {
		return
	}

	byteData, err := json.Marshal(db.Value)
	if err != nil {
		return
	}
	if err = json.Unmarshal(byteData, &m); err != nil {
		return
	}

	return
}

func (srv *roleUserService) Update(user *model.RoleUser, tx *gorm.DB) (err error) {
	err = tx.Save(&user).Error
	return
}

func (srv *roleUserService) Delete(user *model.RoleUser, tx *gorm.DB) (err error) {
	err = tx.Delete(&user).Error
	return
}

func (srv *roleUserService) Find(ID int64) (*model.RoleUser, error) {
	m := new(model.RoleUser)
	row := srv.db.Table("user_role").Select("*").Where("id = ?", ID).Row()
	err := row.Scan(&m.ID, &m.RoleID, &m.UserID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *roleUserService) FindOneBy(criteria map[string]interface{}) (*model.RoleUser, error) {
	m := new(model.RoleUser)
	row := srv.db.Table("user_role").Select("*").Where(criteria).Row()
	err := row.Scan(&m.ID, &m.RoleID, &m.UserID, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}
