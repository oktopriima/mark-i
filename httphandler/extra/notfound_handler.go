/*
 * Copyright (c) 2019
 * Created at 6/2/19 11:07 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package extra

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
)

func (h *handler) NotFoundHandler(ctx *gin.Context) {
	url := ctx.Request.RequestURI
	httpresponse.NewErrorResponse(ctx, http.StatusNotFound, errors.New(url+" page not found"))
	return
}
