package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/company/{id}", GetCompaines).Methods(("GET"))
	r.HandleFunc("/user/{id}", getUsers).Methods("GET")
	r.HandleFunc("/insert-user", insertUser).Methods("POST")
	r.HandleFunc("/insert-company", insertCompanies).Methods("POST")
	r.HandleFunc("/insert-partnership", insertPartnership).Methods("POST")
	r.HandleFunc("/insert-partnership-details", insertPartnershipDetails).Methods("POST")
	r.HandleFunc("/delete-user", deletedUsers).Methods("DELETE")
	r.HandleFunc("/delete-company", deletedCompanies).Methods("DELETE")
	r.HandleFunc("/delete-partnership", partnershipCancelled).Methods("DELETE")
	r.HandleFunc("/all-user", getUseridAllPartnership).Methods("POST")
	r.HandleFunc("/all-company", getCompanyIdAllPartner).Methods("POST")
	r.HandleFunc("/edit-user", editUser).Methods("PUT")
	r.HandleFunc("/edit-company", editCompany).Methods("PUT")

	server := &http.Server{Addr: ":4545", Handler: r}
	server.ListenAndServe()
}
