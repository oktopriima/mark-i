/*
 * Copyright (c) 2019
 * Created at 5/29/19 5:21 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package model

import "time"

type UserRole struct {
	ID        int64     `json:"id"`
	UsersID   int64     `json:"users_id"`
	RolesID   int64     `json:"roles_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Roles     Role      `json:"roles"`
}
