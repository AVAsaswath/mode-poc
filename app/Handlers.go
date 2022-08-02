package app

import (
	"encoding/json"
	"net/http"
)

//type Handlers struct {
//	Service Service.Service
//}

type Success struct {
	Success bool
	ErrorID int
	Message string
}

//func (h *Handlers) VerifiLogin(w http.ResponseWriter, r *http.Request) {
//
//	response := Success{
//		Success: true,
//		ErrorID: 2,
//		Message: "Backend connected Successfully",
//	}
//
//	writeResponse(w, http.StatusOK, response)
//
//}

//func (h *Handlers) GetallStudent(w http.ResponseWriter, r *http.Request) {
//
//	fmt.Println("entered handler")
//	status := r.URL.Query().Get("status")
//
//	customers, err := h.Service.GetAllstudent(status)
//
//	fmt.Println(customers, "handler")
//	if err != nil {
//		writeResponse(w, err.Code, err.AsMessage())
//	} else {
//		writeResponse(w, http.StatusOK, customers)
//	}
//}
//
//func (h *Handlers) GetStudentinfo(w http.ResponseWriter, r *http.Request) {
//
//	vars := mux.Vars(r)
//	id := vars["student_id"]
//
//	customer, err := h.Service.Getstudentinfo(id)
//	if err != nil {
//		writeResponse(w, err.Code, err.AsMessage())
//	} else {
//		writeResponse(w, http.StatusOK, customer)
//	}
//}
//
//func (h *Handlers) DeleteStudentinfo(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["student_id"]
//
//	error := h.Service.DeleteStudentinfo(id)
//	fmt.Println(error)
//
//	writeResponse(w, http.StatusOK, error)
//
//}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
