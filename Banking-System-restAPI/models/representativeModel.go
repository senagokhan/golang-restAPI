package Models

type Representative struct {
	Representative uint   `json:"Representative"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	Email          string `json:"Email"`
	OfficePhone    string `json:"OfficePhone"`
}
