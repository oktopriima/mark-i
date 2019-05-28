/*
 * Copyright (c) 2019
 * Created at 5/28/19 12:10 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

type phoneLoginRequest struct {
	Phone string `json:"phone" binding:"required"`
	Pin   string `json:"pin" binding:"required"`
}

func (req phoneLoginRequest) GetPhone() string {
	return req.Phone
}

func (req phoneLoginRequest) GetPin() string {
	return req.Pin
}
