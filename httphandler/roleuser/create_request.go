/*
 * Copyright (c) 2019
 * Created at 5/29/19 10:48 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

type createRequest struct {
	RoleID int64 `json:"role_id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
}

func (req createRequest) GetRoleID() int64 {
	return req.RoleID
}

func (req createRequest) GetUserID() int64 {
	return req.UserID
}
