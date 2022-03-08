package models

type Apikey struct{
	ApiKey	string `json:"apiKey,omitempty" validate:"required"`
	Quota	int `json:"quota,omitempty" validate:"required"`
	Owner	string `json:"owner,omitempty" validate:"required"`
}