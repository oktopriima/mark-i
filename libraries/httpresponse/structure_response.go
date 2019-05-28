/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:43 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package httpresponse

type ResponseError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
}

type ResponsePaged struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
}

type ResponseObject struct {
	Data interface{} `json:"data"`
}
