package models

type Response struct {
	Meta Meta `json:"meta"`
	Data Data `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Match            Match            `json:"match"`
	Matches          Matches          `json:"matches"`
	User             User             `json:"user"`
	UserQueue        Users            `json:"users"`
	RegisterResponse RegisterResponse `json:"user"`
}

type RegisterResponse struct {
	User        User   `json:"user"`
	AccessToken string `json:"access_token"`
}
