/*
 * Copyright (c) 2019
 * Created at 5/29/19 10:50 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
)

func (h *handler) CreateHandler(ctx *gin.Context) {
	var request createRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.Create(request)
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	ctx.JSON(http.StatusOK, resp.GetData())
}
