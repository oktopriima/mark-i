/*
 * Copyright (c) 2019
 * Created at 5/29/19 10:47 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

import (
	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-i/app/context/roleuser"
)

type Handler interface {
	CreateHandler(ctx *gin.Context)
}

type handler struct {
	uc roleuser.Usecase
}

func NewHandler(uc roleuser.Usecase) Handler {
	return &handler{uc}
}
