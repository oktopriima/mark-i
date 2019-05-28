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
	ID          int       `json:"id"`
	UserID      int64     `json:"user_id"`
	SocialMedia string    `json:"social_media"`
	SocialID    string    `json:"social_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
