package utils

import (
	"cmd/models"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

func Run(query string, variable []byte, query_or_mutation bool, table_type int) string {
	run_service := grq_req(query, variable, query_or_mutation, table_type)
	return run_service
}

func grq_req(query string, mutation []byte, is_mutation bool, table_type int) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	client := graphql.NewClient(os.Getenv("HASURA_URL"))

	if is_mutation == false {
		current_id := string(mutation)

		req := graphql.NewRequest(query)
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("x-hasura-admin-secret", os.Getenv("ADMIN_SECRET"))
		req.Var("id", current_id)

		ctx := context.Background()

		var responseData map[string]interface{}
		if err := client.Run(ctx, req, &responseData); err != nil {
			log.Fatal(err)
		}
		jsonStr, errs := json.Marshal(responseData)

		if errs != nil {
			log.Fatal(errs)
		}

		return string(jsonStr)
	} else {

		req := graphql.NewRequest(query)
		if table_type == 1 {
			var current_models models.Users
			json.Unmarshal([]byte(mutation), &current_models)
			req.Var("email", current_models.Email)
			req.Var("number", current_models.Number)
			req.Var("name", current_models.Name)
			req.Var("surname", current_models.Surname)

		} else if table_type == 2 {
			var current_models models.Companies
			json.Unmarshal([]byte(mutation), &current_models)
			req.Var("company_email", current_models.CompanyEmail)
			req.Var("company_name", current_models.CompanyName)
			req.Var("address_line1", current_models.AdressLine1)
			req.Var("address_line2", current_models.AdressLine2)
			req.Var("company_desc", current_models.CompanyDesc)
			req.Var("company_title", current_models.CompanyTitle)
			req.Var("company_number", current_models.CompanyNumber)

		} else if table_type == 3 {
			var current_models models.Partnership
			json.Unmarshal([]byte(mutation), &current_models)
			req.Var("user_id", current_models.UserId)
			req.Var("company_id", current_models.CompanyId)
			req.Var("id", current_models.Id)

		} else if table_type == 4 {
			var details_models models.PartnershipDetails
			json.Unmarshal([]byte(mutation), &details_models)
			req.Var("partnership_id", details_models.PartnershipId)
			req.Var("account_type", details_models.AccountType)

		} else if table_type == 5 {
			var users models.Users
			json.Unmarshal([]byte(mutation), &users)
			req.Var("user_id", users.Id)
			req.Var("deleted_at", users.DeletedAt)

		} else if table_type == 6 {
			var companies models.Companies
			json.Unmarshal([]byte(mutation), &companies)
			req.Var("company_id", companies.Id)
		} else if table_type == 7 {
			var partnership_details models.PartnershipDetails
			json.Unmarshal([]byte(mutation), &partnership_details)
			req.Var("partnership_id", partnership_details.PartnershipId)
			req.Var("deactive_at", partnership_details.DeactiveAt)

		} else if table_type == 8 {
			var users models.Users
			json.Unmarshal([]byte(mutation), &users)
			req.Var("user_id", users.Id)

		} else if table_type == 9 {
			var companies models.Companies
			json.Unmarshal([]byte(mutation), &companies)
			req.Var("company_id", companies.Id)

		} else if table_type == 10 {
			var users models.Users
			json.Unmarshal([]byte(mutation), &users)
			req.Var("updated_at", users.UpdatedAt)
			req.Var("name", users.Name)
			req.Var("surname", users.Surname)
			req.Var("user_id", users.Id)
			req.Var("number", users.Number)
			req.Var("email", users.Email)

		} else if table_type == 11 {
			var companies models.Companies
			json.Unmarshal([]byte(mutation), &companies)
			req.Var("updated_at", companies.UpdatedAt)
			req.Var("company_email", companies.CompanyEmail)
			req.Var("company_name", companies.CompanyName)
			req.Var("address_line1", companies.AdressLine1)
			req.Var("address_line2", companies.AdressLine2)
			req.Var("company_desc", companies.CompanyDesc)
			req.Var("company_title", companies.CompanyTitle)
			req.Var("company_number", companies.CompanyNumber)
		}

		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("x-hasura-admin-secret", os.Getenv("ADMIN_SECRET"))

		ctx := context.Background()

		var responseData map[string]interface{}
		if err := client.Run(ctx, req, &responseData); err != nil {
			err_return, errs := json.Marshal(err)
			if errs != nil {
				return errs.Error()
			}
			return string(err_return)

		}
		jsonStr, errs := json.Marshal(responseData)

		if errs != nil {
			return errs.Error()
		}
		return string(jsonStr)
	}
}
