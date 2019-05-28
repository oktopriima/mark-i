/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:39 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package ping

type getRequest struct {
	Param string `json:"param"`
}

func (req getRequest) GetParam() string {
	return req.Param
}
