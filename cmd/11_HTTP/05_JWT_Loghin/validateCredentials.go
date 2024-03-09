package main

type UserData struct {
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`
	AllowedIPs []string `json:"IPs"`

	ID      int  `json:"id"`
	IsAdmin bool `json:"isadmin"`
}

// TODO: validate credentials against persistence.
func ValidateCredentials(user, password string) (*UserData, error) {
	return &UserData{
			FirstName: "John",
			LastName:  "Smith",
		},
		nil
}
