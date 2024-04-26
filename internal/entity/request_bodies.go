package entity

type RegisterBody struct {
	Phonenumber string `json:"phonenumber" example:"9009009090" validate:"required,numeric,len=10"`
}

type ConfirmRegisterBody struct {
	Uid  string `json:"uid" example:"9c884669-0dbf-497d-b94f-cfd196278d8f" validate:"required,uuid4"`
	Code string `json:"code" example:"1234" validate:"required,numeric,len=4"`
}

type LoginBody struct {
	Phonenumber string `json:"phonenumber" example:"9009009090" validate:"required,numeric,len=10"`
}

type ConfirmLoginBody struct {
	Uid  string `json:"uid" example:"9c884669-0dbf-497d-b94f-cfd196278d8f" validate:"required,uuid4"`
	Code string `json:"code" example:"1234" validate:"required,numeric,len=4"`
}
