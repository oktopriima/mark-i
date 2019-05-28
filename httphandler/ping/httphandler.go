/*
 * Copyright (c) 2019
 * Created at 5/21/19 1:53 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package ping

import (
	"github.com/oktopriima/mark-i/app/context/ping"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetHandle(ctx *gin.Context)
	PostHandle(ctx *gin.Context)
}

type handler struct {
	kll  httpresponse.KulinaRequestReader
	get  ping.GetUsecase
	post ping.PostUsecase
}

func NewHandler(
	kll httpresponse.KulinaRequestReader,
	get ping.GetUsecase,
	post ping.PostUsecase) Handler {
	return &handler{kll, get, post}
}
