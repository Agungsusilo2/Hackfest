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

func NewApplicationts(applicationRepository repository.ApplicantRepository, DB *sql.DB, validate *validator.Validate) ApplicantService {
	return &ApplicantServiceImp{
		ApplicationRepository: applicationRepository,
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

func (service *ApplicantServiceImp) Update(ctx context.Context, request web.ApplicantUpdateReq) (web.ApplicantResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.ApplicantResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.ApplicantResponse{}, err
	}
	defer helper.ErrorTx(tx)

	applicant, err := service.ApplicationRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return web.ApplicantResponse{}, err
	}

	applicant.EventName = request.EventName
	applicant.Date = request.Date
	applicant.EventVenues = request.EventVenues
	applicant.RequeirementMaterials = request.RequeirementMaterials

	return helper.ToApplicantResponse(applicant), nil
}

func (a *ApplicantServiceImp) FindAll(ctx context.Context) []web.ApplicantResponse {
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

func (a *ApplicantServiceImp) FindById(ctx context.Context, ApplicantId int) web.ApplicantResponse {
	tx, err := a.DB.Begin()
	helper.PanicErr(err)
	defer helper.ErrorTx(tx)

	applicant, err := a.ApplicationRepository.FindById(ctx, tx, ApplicantId)
	helper.PanicErr(err)

	return helper.ToApplicantResponse(applicant)
}

func (a *ApplicantServiceImp) Delete(ctx context.Context, applicantId int) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer helper.ErrorTx(tx)

	err = a.ApplicationRepository.Delete(ctx, tx, applicantId)
	if err != nil {
		return err
	}

	return nil
}
