package dto

type StudentResponse struct {
	StudentId   string `json:"studentid"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Adtype      string `json:"Admission_type"`
	Dept_ID     string `json:"Department_id"`
	DateofBirth string `json:"date_of_birth"`
	Del_Flag    string `json:"status"`
}
