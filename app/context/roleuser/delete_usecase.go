/*
 * Copyright (c) 2019
 * Created at 5/30/19 5:52 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package roleuser

type DeleteResponse interface {
}

type DeleteRequest interface {
	GetID() int64
}

func (uc *usecase) Delete(request DeleteRequest) (DeleteResponse, error) {
	m, err := uc.roleUserRepo.Find(request.GetID())
	if err != nil {
		return nil, err
	}

	tx := uc.db.Begin()
	if err = uc.roleUserRepo.Delete(m, tx); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return deleteResponse{
		Message: "Success",
	}, nil
}
