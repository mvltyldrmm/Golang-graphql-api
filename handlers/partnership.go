package handlers

import (
	"cmd/models"
	"cmd/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func insertPartnership(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var partnership models.Partnership

	partnership.Id = uuid.New()
	partnership.CreatedAt = time.Now()

	json.NewDecoder(r.Body).Decode(&partnership)

	x, errs := json.Marshal(partnership)

	insert_partnership_mutation := `
	mutation insertPartnership($user_id: uuid = "", $id: uuid = "", $company_id: uuid = "") {
		insert_company_management_partnership(objects: {user_id: $user_id, id: $id, company_id: $company_id}) {
		  affected_rows
		  returning {
			id
		  }
		}
	  }
	   
	`

	insert_partnership := utils.Run(insert_partnership_mutation, x, true, 3)

	if errs != nil {
		w.Write([]byte(errs.Error()))
	}
	w.Write([]byte(insert_partnership))
}

func insertPartnershipDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var partnership_details models.PartnershipDetails

	json.NewDecoder(r.Body).Decode(&partnership_details)

	y, errs := json.Marshal(partnership_details)

	insert_partnership_details_mutation := `
	mutation insertPartnershipDetails($partnership_id: uuid = "", $account_type: numeric = "") {
		insert_company_management_partnership_details(objects: {partnership_id: $partnership_id, is_active: true, account_type: $account_type}) {
		  affected_rows
		}
	  }
	`
	insert_partnership_details := utils.Run(insert_partnership_details_mutation, y, true, 4)
	if errs != nil {
		w.Write([]byte(errs.Error()))
	}
	w.Write([]byte(insert_partnership_details))
}

func partnershipCancelled(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var partnership_details models.PartnershipDetails

	partnership_details.DeactiveAt = time.Now()

	json.NewDecoder(r.Body).Decode(&partnership_details)

	y, errs := json.Marshal(partnership_details)

	deleted_partnership := `
	mutation deletePartnership($partnership_id: uuid = "", $deactive_at: timestamp = "") {
		update_company_management_partnership_details(where: {partnership_id: {_eq: $partnership_id}}, _set: {deactive_at: $deactive_at, is_active: false, is_deactive: true}) {
		  affected_rows
		  returning {
			id
			partnership_id
		  }
		}
	  }		
	`
	cancelled_partnership := utils.Run(deleted_partnership, y, true, 7)
	if errs != nil {
		w.Write([]byte(errs.Error()))
	}
	w.Write([]byte(cancelled_partnership))
}

func getUseridAllPartnership(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users models.Users

	json.NewDecoder(r.Body).Decode(&users)

	y, errs := json.Marshal(users)

	all_user_partnership := `
	query getUseridPartnership($user_id: uuid = "") {
		company_management_users(where: {id: {_eq: $user_id}, is_active: {_eq: true}, is_deleted: {_eq: false}, partner: {partner_details: {is_active: {_eq: true}}}}) {
		  country_id
		  created_at
		  deleted_at
		  email
		  id
		  is_active
		  is_deleted
		  language_id
		  name
		  number
		  partner {
			detail_id
			created_at
			company_id
			id
			partner_details {
			  is_deactive
			  id
			  deactive_at
			  account_type
			  is_active
			  partnership_id
			}
		  }
		}
	  }
	   
	`
	get_all_partnership := utils.Run(all_user_partnership, y, true, 8)
	if errs != nil {
		w.Write([]byte(errs.Error()))
	}
	w.Write([]byte(get_all_partnership))
}

func getCompanyIdAllPartner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var companies models.Companies

	json.NewDecoder(r.Body).Decode(&companies)

	y, errs := json.Marshal(companies)

	all_company_partner := `
	query allCompanyPartner($company_id: uuid = "") {
		company_management_companies(where: {id: {_eq: $company_id}, is_active: {_eq: true}, partner: {partner_details: {is_active: {_eq: true}, is_deactive: {_eq: false}}}}) {
		  company_email
		  company_name
		  company_number
		  company_title
		  partner {
			company_id
			created_at
			detail_id
			id
			partner_details {
			  account_type
			  deactive_at
			  id
			  is_deactive
			  is_active
			  partnership_id
			}
			user_id
		  }
		}
	  }
	  
	`
	get_all_partnership := utils.Run(all_company_partner, y, true, 9)
	if errs != nil {
		w.Write([]byte(errs.Error()))
	}
	w.Write([]byte(get_all_partnership))
}
