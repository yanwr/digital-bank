package dtos

type LoginRequestDTO struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}
