/*
 * Copyright (c) 2019
 * Created at 5/29/19 3:05 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package role

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
)

func (h *handler) GenerateHandler(ctx *gin.Context) {
	resp, err := h.uc.Generate()
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	httpresponse.NewSuccessResponse(ctx, resp.GetData())
}
