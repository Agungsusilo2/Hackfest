package controller

import (
	"encoding/json"
	"github.com/Agungsusilo2/Hackfest/model/helper"
	"github.com/Agungsusilo2/Hackfest/model/service"
	"github.com/Agungsusilo2/Hackfest/model/web"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type ApplicantControllerImp struct {
	ApplicantService service.ApplicantService
}

func NewCategoryController(applicationService service.ApplicantService) *ApplicantControllerImp {
	return &ApplicantControllerImp{
		ApplicantService: applicationService,
	}
}

func (a *ApplicantControllerImp) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	applicantCreateReq := web.ApplicantCreateReq{}
	err := decoder.Decode(&applicantCreateReq)
	helper.PanicErr(err)

	applicantResponse := a.ApplicantService.Create(request.Context(), applicantCreateReq)
	responseTemp := web.ResponseTemp{
		Code:   200,
		Status: "OK",
		Data:   applicantResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(responseTemp)
	helper.PanicErr(err)
}

func (a *ApplicantControllerImp) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	applicantUpdateReq := web.ApplicantUpdateReq{}
	err := decoder.Decode(&applicantUpdateReq)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	applicantIDStr := params.ByName("applicantId")
	applicantID, err := strconv.Atoi(applicantIDStr)
	if err != nil {
		http.Error(writer, "Invalid applicant ID", http.StatusBadRequest)
		return
	}

	applicantUpdateReq.Id = applicantID

	applicantResponse, err := a.ApplicantService.Update(request.Context(), applicantUpdateReq)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responseTemp := web.ResponseTemp{
		Code:   200,
		Status: "OK",
		Data:   applicantResponse,
	}

	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(responseTemp)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *ApplicantControllerImp) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryId = params.ByName("applicantId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicErr(err)

	err = a.ApplicantService.Delete(request.Context(), id)
	if err != nil {
		responseTemp := web.ResponseTemp{
			Code:   500,
			Status: "Internal Server Error",
			Data:   nil,
		}
		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(responseTemp)
		helper.PanicErr(err)
		return
	}

	responseTemp := web.ResponseTemp{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(responseTemp)
	helper.PanicErr(err)
}

func (a *ApplicantControllerImp) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var categoryId = params.ByName("applicantId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicErr(err)

	applicantResponse := a.ApplicantService.FindById(request.Context(), id)
	responseTemp := web.ResponseTemp{
		Code:   200,
		Status: "OK",
		Data:   applicantResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(responseTemp)
	helper.PanicErr(err)
}

func (a *ApplicantControllerImp) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	applicantResponse := a.ApplicantService.FindAll(request.Context())
	responseTemp := web.ResponseTemp{
		Code:   200,
		Status: "OK",
		Data:   applicantResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(responseTemp)
	helper.PanicErr(err)
}
