package handlers

import (
	"cmd/models"
	"cmd/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func insertCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var companies models.Companies

	companies.Id = uuid.New()
	companies.CreatedAt = time.Now()
	companies.UpdatedAt = time.Now()
	companies.IsActived = true

	json.NewDecoder(r.Body).Decode(&companies)

	x, errs := json.Marshal(companies)

	insert_company_mutation := `
	mutation insertCompany($company_email: String = "", $company_name: String = "", $company_number: String = "", $company_title: String = "", $company_desc: String = "", $address_line2: String = "", $address_line1: String = "") {
		insert_company_management_companies(objects: {company_email: $company_email, company_name: $company_name, company_number: $company_number, company_title: $company_title, company_desc: $company_desc, address_line2: $address_line2, address_line1: $address_line1}) {
		  affected_rows
		  returning {
			id
			address_line1
			address_line2
			company_desc
			company_email
			company_name
			company_number
			company_title
			is_active
		  }
		}
	  }
	  
	`
	go_insert := utils.Run(insert_company_mutation, x, true, 2)

	if errs != nil {
		w.Write([]byte(errs.Error()))
	}

	w.Write([]byte(go_insert))
}

func deletedCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var companies models.Companies

	companies.DeletedAt = time.Now()

	json.NewDecoder(r.Body).Decode(&companies)

	y, errs := json.Marshal(companies)

	deleted_company := `
	mutation deletedCompany($company_id: uuid = "") {
		update_company_management_companies(where: {id: {_eq: $company_id}}, _set: {is_active: false}) {
		  affected_rows
		  returning {
			id
			is_active
			company_name
		  }
		}
	  }	  
	`
	deactive_at_company := utils.Run(deleted_company, y, true, 6)
	if errs != nil {
		w.Write([]byte(errs.Error()))
	}
	w.Write([]byte(deactive_at_company))
}

func editCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var companies models.Companies

	companies.UpdatedAt = time.Now()

	json.NewDecoder(r.Body).Decode(&companies)

	y, errs := json.Marshal(companies)

	update_users := `
	mutation updateCompany($company_id: uuid = "", $company_email: String = "", $company_name: String = "", $company_number: String = "", $updated_at: timestamp = "", $company_desc: String = "", $company_title: String = "", $address_line1: String = "", $address_line2: String = "") {
		update_company_management_companies(where: {id: {_eq: $company_id}}, _set: {company_email: $company_email, company_name: $company_name, company_number: $company_number, updated_at: $updated_at, company_desc: $company_desc, company_title: $company_title, address_line1: $address_line1, address_line2: $address_line2}) {
		  affected_rows
		  returning {
			company_name
			company_email
			company_number
			id
			company_desc
			address_line2
			address_line1
		  }
		}
	  }	  
	`
	update_company := utils.Run(update_users, y, true, 11)

	if errs != nil {
		w.Write([]byte(errs.Error()))
	}

	w.Write([]byte(update_company))
}