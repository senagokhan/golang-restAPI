package Models

type Branches struct {
	BranchId   uint   `json:"branchId"`
	BranchName string `json:"branchName"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zipCode"`
}
