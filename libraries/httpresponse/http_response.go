/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:42 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package httpresponse

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewErrorResponse(c *gin.Context, code int, err error) {
	response := new(ResponseError)
	response.Status = code
	response.Message = err.Error()
	response.ErrorCode = code
	c.JSON(code, &response)
	return
}

func NewSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	return
}

type RequestReader interface {
	GetJsonForm(c *gin.Context, data interface{}) (err error)
}

type requestReader struct {
}

func NewRequestReader() RequestReader {
	return &requestReader{}
}

func (krr *requestReader) GetJsonForm(c *gin.Context, data interface{}) (err error) {
	err = json.NewDecoder(c.Request.Body).Decode(data)
	return
}
