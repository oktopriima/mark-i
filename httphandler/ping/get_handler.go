/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:41 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package ping

import (
	"net/http"

	"github.com/oktopriima/mark-i/libraries/httpresponse"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetHandle(ctx *gin.Context) {
	request := getRequest{}
	request.Param = ctx.Query("param")

	response, err := h.get.Main(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	ctx.JSON(http.StatusOK, response.GetData())
}
