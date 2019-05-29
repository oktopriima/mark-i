/*
 * Copyright (c) 2019
 * Created at 5/29/19 2:57 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package role

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-i/app/context/role"
)

type Handler interface {
	GenerateHandler(ctx *gin.Context)
}

type handler struct {
	uc role.Usecase
}

func NewHandler(uc role.Usecase) Handler {
	return &handler{uc}
}
