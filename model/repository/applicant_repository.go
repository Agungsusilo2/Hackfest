package repository

import (
	"context"
	"database/sql"
	"github.com/Agungsusilo2/Hackfest/model/domain"
)

type ApplicantRepository interface {
	Save(ctx context.Context, tx sql.Tx, applicant domain.Category) domain.Category
	Update(ctx context.Context, tx sql.Tx, applicant domain.Category) domain.Category
	Delete(ctx context.Context, tx sql.Tx, applicant domain.Category)
	FindById(ctx context.Context, tx sql.Tx, applicantId int) domain.Category
	FindByAll(ctx context.Context) []domain.Category
}
