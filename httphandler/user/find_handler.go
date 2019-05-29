/*
 * Copyright (c) 2019
 * Created at 5/27/19 1:34 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package user

import (
	"net/http"

	"github.com/oktopriima/mark-i/libraries/helper"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
	"github.com/gin-gonic/gin"
)

func (h *handler) FindHandler(ctx *gin.Context) {
	userID, err := helper.GetAuthenticatedUser(ctx.Request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.Find(userID)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp)
}
