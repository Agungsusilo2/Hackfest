package service

import (
	"context"
	"database/sql"
	"github.com/Agungsusilo2/Hackfest/model/domain"
	"github.com/Agungsusilo2/Hackfest/model/helper"
	"github.com/Agungsusilo2/Hackfest/model/repository"
	"github.com/Agungsusilo2/Hackfest/model/web"
	"github.com/go-playground/validator"
)

type ApplicantServiceImp struct {
	ApplicationRepository repository.ApplicantRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewApplicationService(categoryRepository repository.ApplicantRepository, DB *sql.DB, validate *validator.Validate) *ApplicantServiceImp {
	return &ApplicantServiceImp{
		ApplicationRepository: categoryRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (a *ApplicantServiceImp) Create(ctx context.Context, request web.ApplicantCreateReq) web.ApplicantResponse {
	err := a.Validate.Struct(request)
	helper.PanicErr(err)
	tx, err := a.DB.Begin()
	helper.PanicErr(err)
	defer helper.ErrorTx(tx)

	applicant := domain.Applicant{
		EventName:             request.EventName,
		EventVenues:           request.EventVenues,
		RequeirementMaterials: request.RequeirementMaterials,
		Date:                  request.Date,
	}

	applicant = a.ApplicationRepository.Save(ctx, tx, applicant)

	return helper.ToApplicantResponse(applicant)
}

func (a ApplicantServiceImp) Update(ctx context.Context, updateReq web.ApplicantUpdateReq) web.ApplicantResponse {
	err := a.Validate.Struct(updateReq)
	helper.PanicErr(err)
	tx, err := a.DB.Begin()
	helper.PanicErr(err)
	defer helper.ErrorTx(tx)

	applicant, err := a.ApplicationRepository.FindById(ctx, tx, updateReq.Id)

	applicant.EventName = updateReq.EventName
	applicant.Date = updateReq.Date
	applicant.EventVenues = updateReq.EventVenues
	applicant.RequeirementMaterials = updateReq.RequeirementMaterials

	applicant = a.ApplicationRepository.Update(ctx, tx, applicant)

	return helper.ToApplicantResponse(applicant)
}

func (a ApplicantServiceImp) Delete(ctx context.Context, ApplicantId int) {
	tx, err := a.DB.Begin()
	helper.PanicErr(err)
	defer helper.ErrorTx(tx)

	applicant, err := a.ApplicationRepository.FindById(ctx, tx, ApplicantId)
	helper.PanicErr(err)

	a.ApplicationRepository.Delete(ctx, tx, applicant)
}

func (a ApplicantServiceImp) FindAll(ctx context.Context) []web.ApplicantResponse {
	tx, err := a.DB.Begin()
	helper.PanicErr(err)
	defer helper.ErrorTx(tx)

	applicants := a.ApplicationRepository.FindByAll(ctx, tx)

	var applicantResponses []web.ApplicantResponse

	for _, applicant := range applicants {
		applicantResponses = append(applicantResponses, helper.ToApplicantResponse(applicant))
	}

	return applicantResponses
}

func (a ApplicantServiceImp) FindById(ctx context.Context, ApplicantId int) web.ApplicantResponse {
	tx, err := a.DB.Begin()
	helper.PanicErr(err)
	defer helper.ErrorTx(tx)

	applicant, err := a.ApplicationRepository.FindById(ctx, tx, ApplicantId)
	helper.PanicErr(err)

	return helper.ToApplicantResponse(applicant)
}
