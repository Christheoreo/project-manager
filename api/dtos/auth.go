package dtos

type (
	AuthLoginDto struct {
		Email    string `bson:"email,omitempty" json:"email"`
		Password string `bson:"password,omitempty" json:"password"`
	}

	JwtResponse struct {
		AccessToken string `json:"accessToken"`
	}
)
