package entity

type GoogleUser struct {
	Sub       string `json:"sub"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	LastName  string `json:"family_name"`
	FirstName string `json:"given_name"`
	Picture   string `json:"picture"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}
