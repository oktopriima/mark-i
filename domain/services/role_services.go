/*
 * Copyright (c) 2019
 * Created at 5/29/19 2:30 PM
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

type roleServices struct {
	db *gorm.DB
}

func NewRoleServices(db *gorm.DB) repository.RoleRepository {
	return &roleServices{db}
}

func (srv *roleServices) Create(role *model.Role, tx *gorm.DB) (m *model.Role, err error) {
	db := tx.Create(&role)
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

func (srv *roleServices) Find(ID int64) (*model.Role, error) {
	m := new(model.Role)
	row := srv.db.Table("roles").Select("*").Where("id = ?", ID).Row()
	err := row.Scan(&m.ID, &m.Value, &m.Description, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (srv *roleServices) FindOneBy(criteria map[string]interface{}) (*model.Role, error) {
	m := new(model.Role)
	row := srv.db.Table("roles").Select("*").Where(criteria).Row()
	err := row.Scan(&m.ID, &m.Value, &m.Description, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}
