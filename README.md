GO + POSTGRESQL DB + HASURA (GRAPHQL) API

GO MOD TIDY (PACKAGE INSTALL)

GO RUN cmd/main.go

GO VERSION = go1.18.3 linux/amd64

<hr>
REQUIRED => # HASURA URL - HASURA SECRET KEY
<hr>

## Routes

| METHOD | PATH                      | AUTH |
| ------ | ------------------------- | ---- |
| GET    | company/:id               | NO   |
| GET    | user/:id                  | NO   |
| POST   | insert-user               | NO   |
| POST   | insert-company            | NO   |
| POST   | insert-partnership        | NO   |
| POST   | insert-partnership-detail | NO   |
| DELETE | delete-user               | NO   |
| DELETE | delete-company            | NO   |
| DELETE | delete-partnership        | NO   |
| POST   | all-company               | NO   |
| POST   | all-user                  | NO   |
| PUT    | edit-user                 | NO   |
| PUT    | edit-company              | NO   |

## DB TABLE

<b>TABLE</b>
COMPANIES , USERS , PARTNERSHIP , PARTNERSHIP_DETAILS

<hr>

<b>RELATIONSHIP</b>
PARTNERSHIP.company_id = COMPANIES.id
PARTNERSHIP.user_id = USERS.id
PARTNERSHIP_DETAILS.partnership_id = PARTNERSHIP.id
