/*
 * Copyright (c) 2019
 * Created at 5/28/19 12:27 PM
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

func (h *handler) PhoneLoginHandler(ctx *gin.Context) {
	request := phoneLoginRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.PhoneLogin(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
