/*
 * Copyright (c) 2019
 * Created at 5/30/19 5:58 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
)

func (h *handler) DeleteHandler(ctx *gin.Context) {
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		httpresponse.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}

	resp, err := h.uc.
}
