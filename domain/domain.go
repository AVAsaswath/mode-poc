package domain

import (
	"golang/dto"
	"golang/errs"
)

type Repository interface {
	FindAll(string) ([]Student, *errs.AppError)
	FindById(string) (*Student, *errs.AppError)
	DelById(string) *errs.AppError
}

type Student struct {
	StudentId   string `db:"StudentID"`
	FirstName   string `db:"FirstName"`
	LastName    string `db:"LastName"`
	Adtype      string `db:"AdType"`
	Dept_ID     string `db:"Dept_ID"`
	Dept_Name   string `db:"Dept_Name"`
	DateofBirth string `db:"DOB"`
	Age         string `db:"age"`
	Status      string `db:"Del_flag"`
}

func (s Student) statusAsText() string {

	statusAsText := "active"
	if s.Status == "1" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (s Student) ToSDto() dto.StudentResponse {

	return dto.StudentResponse{
		StudentId:   s.StudentId,
		FirstName:   s.FirstName,
		LastName:    s.LastName,
		Adtype:      s.Adtype,
		Dept_ID:     s.Dept_ID,
		DateofBirth: s.DateofBirth,
		Del_Flag:    s.statusAsText(),
	}
}
