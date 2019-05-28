/*
 * Copyright (c) 2019
 * Created at 5/26/19 1:49 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

type authRequest struct {
	SocialType  string `json:"social_type" binding:"required"`
	SocialToken string `json:"social_token" binding:"required"`
}

func (req authRequest) GetSocialType() string {
	return req.SocialType
}

func (req authRequest) GetSocialToken() string {
	return req.SocialToken
}
