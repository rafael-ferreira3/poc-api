package dto

type LoginRequestDTO struct {
	Username string
	Password string
}

type LoginResponseDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"nome"`
	AccessToken string `json:"access_token"`
}
