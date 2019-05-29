/*
 * Copyright (c) 2019
 * Created at 5/29/19 2:44 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package role

import (
	"github.com/oktopriima/mark-i/app/config"
	"github.com/oktopriima/mark-i/domain/model"
)

type GenerateResponse interface {
	GetData() []*model.Role
}

func (uc *usecase) Generate() (GenerateResponse, error) {
	var myrole []string
	var roles []*model.Role

	myrole = append(myrole, config.ADMIN, config.MERCHANT, config.CONSUMER)

	for _, assignrole := range myrole {
		tx := uc.db.Begin()
		role, err := uc.findByValue(assignrole)
		if err != nil || role.Value != assignrole {
			/** create new role value */
			param := new(model.Role)
			param.Value = assignrole
			param.Description = ""
			role, err = uc.create(param, tx)
			if err != nil {
				tx.Rollback()
				continue
			}
		}
		tx.Commit()

		roles = append(roles, role)
	}

	return generateResponse{
		Data: roles,
	}, nil
}
