package models

type (
	AuthLoginDto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	JwtResponse struct {
		AccessToken string `json:"accessToken"`
	}
)
