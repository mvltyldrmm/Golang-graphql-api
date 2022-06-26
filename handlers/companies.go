package handlers

import (
	"cmd/utils"
	"net/http"

	"github.com/gorilla/mux"
)

//TEST BASE
func GetCompaines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_key := vars["id"]
	data := []byte(id_key)
	get_companies_query := `
    query MyQuery($id: uuid = "") {
		company_management_companies(where: {is_active: {_eq: true}, id: {_eq: $id}}) {
		  company_desc
		  company_name
		  company_email
		  company_number
		  company_title
		  is_active
		  created_at
		}
	  }
	  
`
	x := utils.Run(get_companies_query, data, false, 0)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(x))
}
