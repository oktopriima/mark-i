/*
 * Copyright (c) 2019
 * Created at 5/26/19 1:44 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-i/app/context/auth"
)

type Handler interface {
	GoogleLoginHandle(ctx *gin.Context)
	FacebookLoginHandle(ctx *gin.Context)
	PhoneLoginHandler(ctx *gin.Context)
	EmailLoginHandler(ctx *gin.Context)
}

type handler struct {
	uc auth.Usecase
}

func NewHandler(uc auth.Usecase) Handler {
	return &handler{uc}
}
