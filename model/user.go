package model

type UserInfo struct {
	User  string `json:"user" example:"Ye"`
	Color string `json:"value" example:"Blue"`
}

type EditResponse struct {
	Status string `json:"status" example:"ok"`
}
