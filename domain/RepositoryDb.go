package domain

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang/errs"
	"golang/logger"
)

type RepositoryDb struct {
	client *sqlx.DB
}

//type RepositoryDb struct {
//	client *sql.DB
//}

func (d RepositoryDb) FindAll(status string) ([]Student, *errs.AppError) {
	fmt.Println("entered Repository")

	var err error

	students := make([]Student, 0)

	if status == "" {
		fmt.Println("entered if")
		findAllSql := "select * from StudentInfo"
		err = d.client.Select(&students, findAllSql)
		return nil, errs.NewUnexpectedError("Unexpected database error")

	} else {
		//findAllSql := "select * from StudentInfo where Del_flag =$" + status
		findAllSql := "select * from StudentInfo"
		err = d.client.Select(&students, findAllSql, status)

	}

	if err != nil {
		fmt.Println("entered error")
		logger.Error("Error while Querying student table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	fmt.Println(&students)
	return students, nil

}

func (d RepositoryDb) FindById(id string) (*Student, *errs.AppError) {
	studentSql := "select * from StudentInfo where StudentID =" + id

	fmt.Println("entered findbyid", studentSql)
	var c Student
	err := d.client.Get(&c, studentSql)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("student not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error ")
		}
	}
	return &c, nil
}

func (d RepositoryDb) DelById(id string) *errs.AppError {
	fmt.Println(id, "70")
	//_, err := d.client.Query("Updatestudentinfo", id)
	//UPDATE StudentInfo
	//SET Del_flag = 1
	//WHERE StudentID = studid
	//res, err := d.client.Exec("Updatestudentinfo", id)
	query := "call Updatestudentinfo(" + id + ")"
	res, err := d.client.Query(query)

	fmt.Println(query)
	fmt.Println(res)
	fmt.Println(err)
	//fmt.Println(err.Error())
	fmt.Println(res.Err())
	fmt.Println(res.Columns())
	fmt.Println(res.Next())
	fmt.Println(res.ColumnTypes())
	fmt.Println(res.Scan())

	if res != nil {
		return errs.NewNotFoundError("id not found")
	}
	if err == nil && res != nil {
		fmt.Println("entered = nill")
		return errs.SuccessResponse("Deleted succeessfully")
	}

	if err != nil {
		fmt.Println("entered != nill")

		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customer " + err.Error())
			return errs.NewNotFoundError("student not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return errs.NewUnexpectedError("unexpected database error ")
		}
	}
	return nil
}

func NewRepositoryDb(dbClient *sqlx.DB) RepositoryDb {
	return RepositoryDb{dbClient}
}
