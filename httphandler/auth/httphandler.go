/*
 * Copyright (c) 2019
 * Created at 5/26/19 1:44 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"github.com/oktopriima/mark-i/app/context/auth"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GoogleLoginHandle(ctx *gin.Context)
	FacebookLoginHandle(ctx *gin.Context)
	PhoneLoginHandler(ctx *gin.Context)
	EmailLoginHandler(ctx *gin.Context)
}

type handler struct {
	kll httpresponse.KulinaRequestReader
	uc  auth.AuthUsecase
}

func NewHandler(reader httpresponse.KulinaRequestReader,
	uc auth.AuthUsecase) Handler {
	return &handler{reader, uc}
}
