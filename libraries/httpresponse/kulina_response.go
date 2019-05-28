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

type KulinaRequestReader interface {
	GetJsonForm(c *gin.Context, data interface{}) (err error)
}

type kulinaRequestReader struct {
}

func NewKulinaRequestReader() KulinaRequestReader {
	return &kulinaRequestReader{}
}

func (krr *kulinaRequestReader) GetJsonForm(c *gin.Context, data interface{}) (err error) {
	err = json.NewDecoder(c.Request.Body).Decode(data)
	return
}