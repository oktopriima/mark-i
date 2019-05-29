/*
 * Copyright (c) 2019
 * Created at 5/29/19 2:28 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package model

import "time"

type Role struct {
	ID          int64     `json:"id"`
	Value       string    `json:"value"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
