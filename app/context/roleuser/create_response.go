/*
 * Copyright (c) 2019
 * Created at 5/29/19 10:39 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

import "github.com/oktopriima/mark-i/domain/model"

type createResponse struct {
	Data *model.UserRole
}

func (resp createResponse) GetData() *model.UserRole {
	return resp.Data
}
