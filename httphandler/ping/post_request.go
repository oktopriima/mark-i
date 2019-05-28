/*
 * Copyright (c) 2019
 * Created at 5/21/19 2:03 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package ping

type postRequest struct {
	Param string `json:"param" binding:"required"`
}

func (req postRequest) GetParam() string {
	return req.Param
}
