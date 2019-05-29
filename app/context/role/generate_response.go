/*
 * Copyright (c) 2019
 * Created at 5/29/19 3:36 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package role

import "github.com/oktopriima/mark-i/domain/model"

type generateResponse struct {
	Data []*model.Role
}

func (resp generateResponse) GetData() []*model.Role {
	return resp.Data
}
