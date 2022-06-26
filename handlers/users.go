package handlers
//TEST BASE
import (
	"cmd/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id_key := vars["id"]
	data := []byte(id_key)
	get_users_query := `
    query getUsers($id: uuid = "") {
		company_management_users(where: {id: {_eq: $id}, is_active: {_eq: true}, is_deleted: {_eq: false}}) {
		  email
		  id
		  is_active
		  is_deleted
		  name
		  number
		  surname
		}
	  }	  
`

	x := utils.Run(get_users_query, data, false, 0)
	w.Write([]byte(x))
}
