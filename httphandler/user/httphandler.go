/*
 * Copyright (c) 2019
 * Created at 5/27/19 1:33 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package user

import (
	"github.com/oktopriima/mark-i/app/context/user"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	FindHandler(ctx *gin.Context)
}

type handler struct {
	kll httpresponse.KulinaRequestReader
	uc  user.UserUsecase
}

func NewHandler(kll httpresponse.KulinaRequestReader, uc user.UserUsecase) Handler {
	return &handler{kll, uc}
}
