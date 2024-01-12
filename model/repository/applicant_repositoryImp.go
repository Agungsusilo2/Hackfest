package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Agungsusilo2/Hackfest/model/domain"
	"github.com/Agungsusilo2/Hackfest/model/helper"
)

type ApplicantRepositoryImp struct {
}

func NewCategoryRepository() ApplicantRepository {
	return &ApplicantRepositoryImp{}
}
func (a *ApplicantRepositoryImp) Save(ctx context.Context, tx *sql.Tx, applicant domain.Applicant) domain.Applicant {
	script := "INSERT INTO applicant(event_name, date_name, event_venues, requeirement_materials) VALUES (?, ?, ?, ?)"
	prep, err := tx.PrepareContext(ctx, script)
	helper.PanicErr(err)
	defer prep.Close()

	rows, err := prep.ExecContext(ctx, applicant.EventName, applicant.Date, applicant.EventVenues, applicant.RequeirementMaterials)
	helper.PanicErr(err)

	id, err := rows.LastInsertId()
	helper.PanicErr(err)
	applicant.Id = int32(id)
	return applicant
}

func (repository *ApplicantRepositoryImp) Update(ctx context.Context, tx *sql.Tx, category domain.Applicant) domain.Applicant {
	script := "UPDATE applicant SET event_name=?, date_name=?, event_venues=?, requeirement_materials=? WHERE id=?"
	_, err := tx.ExecContext(ctx, script, category.EventName, category.Date, category.EventVenues, category.EventVenues, category.Id)
	helper.PanicErr(err)

	return category
}

func (a *ApplicantRepositoryImp) FindById(ctx context.Context, tx *sql.Tx, applicantId int) (domain.Applicant, error) {
	script := "SELECT id,event_name,date_name,event_venues,requeirement_materials FROM applicant WHERE id = ? LIMIT 1"
	prep, err := tx.PrepareContext(ctx, script)
	helper.PanicErr(err)
	rows, err := prep.QueryContext(ctx, applicantId)
	helper.PanicErr(err)

	applicant := domain.Applicant{}
	if rows.Next() {
		err := rows.Scan(&applicant.Id, &applicant.EventName, &applicant.Date, &applicant.EventVenues, &applicant.RequeirementMaterials)
		helper.PanicErr(err)
		return applicant, nil
	} else {
		return applicant, errors.New("Applicant is not found")
	}
}

func (a ApplicantRepositoryImp) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Applicant {
	script := "SELECT id,event_name,date_name,event_venues,requeirement_materials FROM applicant"
	prep, err := tx.PrepareContext(ctx, script)
	helper.PanicErr(err)

	defer prep.Close()
	rows, err := prep.QueryContext(ctx)
	var applicants []domain.Applicant

	for rows.Next() {
		applicant := domain.Applicant{}
		err := rows.Scan(&applicant.Id, &applicant.EventName, &applicant.Date, &applicant.EventVenues, &applicant.RequeirementMaterials)
		helper.PanicErr(err)
		applicants = append(applicants, applicant)
	}
	return applicants
}

func (i *ApplicantRepositoryImp) Delete(ctx context.Context, tx *sql.Tx, applicantId int) error {
	script := "DELETE FROM applicant WHERE id = ? LIMIT 1"
	prep, err := tx.PrepareContext(ctx, script)
	if err != nil {
		return err
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, applicantId)
	if err != nil {
		return err
	}

	return nil
}
