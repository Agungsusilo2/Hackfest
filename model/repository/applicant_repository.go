package repository

import (
	"context"
	"database/sql"
	"github.com/Agungsusilo2/Hackfest/model/domain"
)

type ApplicantRepository interface {
	Save(ctx context.Context, tx *sql.Tx, applicant domain.Applicant) domain.Applicant
	Update(ctx context.Context, tx *sql.Tx, applicant domain.Applicant) domain.Applicant
	Delete(ctx context.Context, tx *sql.Tx, applicant domain.Applicant)
	FindById(ctx context.Context, tx *sql.Tx, applicantId int) (domain.Applicant, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Applicant
}
