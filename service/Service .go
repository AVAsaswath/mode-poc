package Service

//
//import (
//	"fmt"
//	"golang/domain"
//	"golang/dto"
//	"golang/errs"
//)
//
//type Service interface {
//	GetAllstudent(string) ([]dto.StudentResponse, *errs.AppError)
//	Getstudentinfo(string) (*dto.StudentResponse, *errs.AppError)
//	DeleteStudentinfo(string) *errs.AppError
//}
//
//type DefaultService struct {
//	repo domain.Repository
//}
//
//func (s DefaultService) GetAllstudent(status string) ([]dto.StudentResponse, *errs.AppError) {
//	fmt.Println("entered service")
//
//	Persons, err := s.repo.FindAll(status)
//	if err != nil {
//		return nil, err
//	}
//	response := make([]dto.StudentResponse, 0)
//	for _, c := range Persons {
//		response = append(response, c.ToSDto())
//	}
//	return response, err
//
//}
//
//func (s DefaultService) Getstudentinfo(id string) (*dto.StudentResponse, *errs.AppError) {
//
//	c, err := s.repo.FindById(id)
//	if err != nil {
//		return nil, err
//	}
//	response := c.ToSDto()
//	return &response, nil
//
//}
//
//func (s DefaultService) DeleteStudentinfo(id string) *errs.AppError {
//
//	response := s.repo.DelById(id)
//
//	return response
//}
//
//func NewService(repository domain.Repository) DefaultService {
//	return DefaultService{repository}
//}
