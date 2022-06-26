package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/all-companies/{id}", GetCompaines).Methods(("GET"))
	r.HandleFunc("/all-users/{id}", getUsers).Methods("GET")
	r.HandleFunc("/insert-user", insertUser).Methods("POST")
	r.HandleFunc("/insert-company", insertCompanies).Methods("POST")
	r.HandleFunc("/insert-partnership", insertPartnership).Methods("POST")
	r.HandleFunc("/insert-partnership-details", insertPartnershipDetails).Methods("POST")
	r.HandleFunc("/delete-user", deletedUsers).Methods("POST")
	r.HandleFunc("/delete-company", deletedCompanies).Methods("POST")
	r.HandleFunc("/delete-partnership", partnershipCancelled).Methods("POST")
	r.HandleFunc("/all-userid", getUseridAllPartnership).Methods("POST")
	r.HandleFunc("/all-companyid", getCompanyIdAllPartner).Methods("POST")
	r.HandleFunc("/edit-user", editUser).Methods("POST")
	r.HandleFunc("/edit-company", editCompany).Methods("POST")

	server := &http.Server{Addr: ":4545", Handler: r}
	server.ListenAndServe()
}
