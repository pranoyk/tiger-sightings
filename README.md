## TIGER-SIGHTINGS

This application helps manage tiger sightings details.
All create endpoint are protected by auth0 authentication

User first needs to create a user account from `/register` endpoint and then login with `/login` endpoint.
The `/login` will return a `user_id` which needs to be set as a `Authorization` header bearer token for subsequent create tiger and create tiger sightings call.

#### Steps to run the application
1. Add the below env variables
   1. DB_USER
   2. DB_PASSWORD
   3. DB_NAME
   4. DB_HOST
   5. DB_PORT
   6. DB_SSLMODE
   7. AUTH0_DOMAIN
   8. AUTH0_CLIENT_ID
   9. AUTH0_CLIENT_SECRET
   10. AUTH0_AUDIENCE
2. run `make docker-compose up`
3. run `make migrate-up`
4. run `go run main.go`

#### Run test
`go test ./... -count=1`

