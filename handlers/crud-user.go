package handlers

import (
	"cmd/models"
	"cmd/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func insertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users models.Users

	users.Id = uuid.New()
	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()
	users.IsActıve = true

	json.NewDecoder(r.Body).Decode(&users)

	x, errs := json.Marshal(users)

	insert_users_mutation := `
	mutation insertUsers($email: String = "", $name: String = "", $number: String = "", $surname: String = "") {
		insert_company_management_users(objects: {email: $email, name: $name, number: $number, surname: $surname}) {
		  returning {
			id
		  }
		}
	  }
	  
	`
	insert_user := utils.Run(insert_users_mutation, x, true, 1)

	if errs != nil {
		w.Write([]byte(errs.Error()))
	}

	w.Write([]byte(insert_user))
}

func deletedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users models.Users

	users.DeletedAt = time.Now()

	json.NewDecoder(r.Body).Decode(&users)

	y, errs := json.Marshal(users)

	deleted_users := `
	mutation deletedUsers($user_id: uuid = "", $deleted_at: timestamp = "") {
		update_company_management_users(where: {id: {_eq: $user_id}}, _set: {is_deleted: true, is_active: false, deleted_at: $deleted_at}) {
		  affected_rows
		  returning {
			id
			is_deleted
		  }
		}
	  }	  
	`
	user_deleted := utils.Run(deleted_users, y, true, 5)

	if errs != nil {
		w.Write([]byte(errs.Error()))
	}
	w.Write([]byte(user_deleted))
}

func editUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users models.Users

	users.UpdatedAt = time.Now()

	json.NewDecoder(r.Body).Decode(&users)

	y, errs := json.Marshal(users)

	update_users := `
	mutation updateUsers($user_id: uuid = "", $updated_at: timestamp = "", $surname: String = "", $name: String = "", $number: String = "", $email: String = "") {
		update_company_management_users(where: {id: {_eq: $user_id}}, _set: {updated_at: $updated_at, surname: $surname, name: $name, number: $number, email: $email}) {
		  affected_rows
		  returning {
			name
			number
			language_id
			surname
		  }
		}
	  }	  
		
	`
	update_user := utils.Run(update_users, y, true, 10)

	if errs != nil {
		w.Write([]byte(errs.Error()))
	}
	w.Write([]byte(update_user))
}
