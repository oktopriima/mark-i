/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:26 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package model

import (
	"time"
)

type Users struct {
	ID             int64        `json:"id"`
	Name           string       `json:"name"`
	Username       string       `json:"username"`
	Email          string       `json:"email"`
	Phone          string       `json:"phone"`
	Password       string       `json:"password"`
	LastLogin      time.Time    `json:"last_login"`
	SecondaryEmail string       `json:"secondary_email"`
	IsVerified     bool         `json:"is_verified"`
	Avatar         string       `json:"avatar"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	Socials        []UserSocial `json:"socials"`
	RoleUsers      []UserRole   `json:"role_users"`
}
