package helper

import (
	"github.com/Agungsusilo2/Hackfest/model/domain"
	"github.com/Agungsusilo2/Hackfest/model/web"
)

func ToApplicantResponse(applicant domain.Applicant) web.ApplicantResponse {
	return web.ApplicantResponse{
		Id:                   applicant.Id,
		EventName:            applicant.EventName,
		DateEvent:            applicant.Date,
		EventVenues:          applicant.EventVenues,
		RequirementMaterials: applicant.RequeirementMaterials,
	}
}
