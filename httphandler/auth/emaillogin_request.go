/*
 * Copyright (c) 2019
 * Created at 5/28/19 12:58 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

type emailLoginRequest struct {
	Email string `json:"email" binding:"required"`
	Pin   string `json:"pin" binding:"required"`
}

func (req emailLoginRequest) GetEmail() string {
	return req.Email
}

func (req emailLoginRequest) GetPin() string {
	return req.Pin
}
