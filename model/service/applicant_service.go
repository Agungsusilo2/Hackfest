package service

import (
	"context"
	"github.com/Agungsusilo2/Hackfest/model/web"
)

type ApplicantService interface {
	Create(ctx context.Context, request web.ApplicantCreateReq) web.ApplicantResponse
	FindAll(ctx context.Context) []web.ApplicantResponse
	FindById(ctx context.Context, ApplicantId int) web.ApplicantResponse
	Delete(ctx context.Context, ApplicantId int) error
	Update(ctx context.Context, request web.ApplicantUpdateReq) (web.ApplicantResponse, error)
}
