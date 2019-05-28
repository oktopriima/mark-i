/*
 * Copyright (c) 2019
 * Created at 5/21/19 2:04 AM
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

func (h *handler) PostHandle(ctx *gin.Context) {
	request := postRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.post.Main(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
