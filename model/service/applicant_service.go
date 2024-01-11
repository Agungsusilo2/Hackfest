package service

import (
	"context"
	"github.com/Agungsusilo2/Hackfest/model/web"
)

type ApplicantService interface {
	Create(ctx context.Context, request web.ApplicantCreateReq) web.ApplicantResponse
	Update(ctx context.Context, request web.ApplicantUpdateReq) web.ApplicantResponse
	Delete(ctx context.Context, ApplicantId int)
	FindAll(ctx context.Context) []web.ApplicantResponse
	FindById(ctx context.Context, ApplicantId int) web.ApplicantResponse
}
