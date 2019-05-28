/*
 * Copyright (c) 2019
 * Created at 5/28/19 11:52 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package auth

import (
	"net/http"

	"github.com/oktopriima/mark-i/libraries/httpresponse"
	"github.com/gin-gonic/gin"
)

func (h *handler) FacebookLoginHandle(ctx *gin.Context) {
	request := authRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.FacebookLogin(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
