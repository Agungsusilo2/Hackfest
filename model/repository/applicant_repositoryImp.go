package repository

import (
	"context"
	"database/sql"
	"github.com/Agungsusilo2/Hackfest/model/domain"
)

type ApplicantRepositoryImp struct {
}

func (a ApplicantRepositoryImp) Save(ctx context.Context, tx sql.Tx, applicant domain.Category) domain.Category {
	//TODO implement me
	panic("implement me")
}

func (a ApplicantRepositoryImp) Update(ctx context.Context, tx sql.Tx, applicant domain.Category) domain.Category {
	//TODO implement me
	panic("implement me")
}

func (a ApplicantRepositoryImp) Delete(ctx context.Context, tx sql.Tx, applicant domain.Category) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicantRepositoryImp) FindById(ctx context.Context, tx sql.Tx, applicantId int) domain.Category {
	//TODO implement me
	panic("implement me")
}

func (a ApplicantRepositoryImp) FindByAll(ctx context.Context) []domain.Category {
	//TODO implement me
	panic("implement me")
}
