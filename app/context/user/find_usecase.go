/*
 * Copyright (c) 2019
 * Created at 5/27/19 1:28 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package user

type FindResponse interface {
}

func (uc *usecase) Find(ID int64) (FindResponse, error) {
	users, err := uc.userRepo.Find(ID)
	if err != nil {
		return nil, err
	}

	return users, nil
}
