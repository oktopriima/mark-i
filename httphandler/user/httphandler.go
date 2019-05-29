/*
 * Copyright (c) 2019
 * Created at 5/27/19 1:33 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-i/app/context/user"
)

type Handler interface {
	FindHandler(ctx *gin.Context)
}

type handler struct {
	uc user.Usecase
}

func NewHandler(uc user.Usecase) Handler {
	return &handler{uc}
}
