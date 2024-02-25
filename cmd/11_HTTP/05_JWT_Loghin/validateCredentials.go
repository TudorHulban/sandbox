package main

type UserData struct {
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`
	AllowedIPs []string `json:"IPs"`

	ID      int  `json:"id"`
	IsAdmin bool `json:"isadmin"`
}

func ValidateCredentials(user, password string) (*UserData, error) {
	// TODO: validate credentials against persistence.

	return &UserData{
			FirstName: "John",
			LastName:  "Smith",
		},
		nil
}
