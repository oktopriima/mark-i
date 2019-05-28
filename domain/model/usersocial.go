/*
 * Copyright (c) 2019
 * Created at 5/24/19 1:00 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package model

import (
	"time"
)

type UserSocial struct {
	ID         int64     `json:"id"`
	UsersID    int64     `json:"users_id"`
	SocialName string    `json:"social_name"`
	SocialID   string    `json:"social_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
