package web

type ApplicantUpdateReq struct {
	Id                    int
	EventName             string `validate:"required,max=255,min=1"`
	EventVenues           string `validate:"required,max=255,min=1"`
	RequeirementMaterials string `validate:"required,max=255,min=1"`
	Date                  string `validate:"required,max=255,min=1"`
}
