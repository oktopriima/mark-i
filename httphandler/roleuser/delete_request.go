/*
 * Copyright (c) 2019
 * Created at 5/30/19 5:57 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

type deleteRequest struct {
	ID int64 `json:"id"`
}

func (req deleteRequest) GetID() int64 {
	return req.ID
}
