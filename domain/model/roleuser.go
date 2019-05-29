/*
 * Copyright (c) 2019
 * Created at 5/29/19 5:21 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package model

import "time"

type RoleUser struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	RoleID    int64     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
