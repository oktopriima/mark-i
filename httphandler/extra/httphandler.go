/*
 * Copyright (c) 2019
 * Created at 6/2/19 11:06 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package extra

import "github.com/gin-gonic/gin"

type Handler interface {
	NotFoundHandler(ctx *gin.Context)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}
