# Project Title

This is a Go project that uses SQL to manage entries in a database. The project uses the sqlc tool to generate Go code from SQL queries, and it uses the github.com/jackc/puddle/v2 package for database connection pooling.  
Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.  

## Prerequisites

* Go 1.18 or later
* SQLC 1.10.0 or later
* PostgreSQL database

## Installing

Clone the repository to your local machine.
Navigate to the project directory.
Run go mod tidy to download the necessary Go modules.
Running the tests
This project uses the testing package from the Go standard library for testing. To run the tests, navigate to the project directory and run `go test ./.... ` 

## Built With

* Go - The programming language
* SQLC - Used to generate Go code from SQL
* PostgreSQL - The database 

## Contributing

Please read CONTRIBUTING.md for details on our code of conduct, and the process for submitting pull requests to us.  
Versioning
We use SemVer for versioning. For the versions available, see the tags on this repository. 

### Helper Commands
1. To create a migrate file in the migrations directory, run the following
     command:  `migrate create -ext sql -dir db/migration -seq users`
2. To generate Go code from SQL queries, run `sqlc generate` in the project directory.
3. UUID dependency: `go get github.com/google/uuid`
4. JWT dependency: `go get -u github.com/golang-jwt/jwt/v5`
5.  Paseto dependency: `go get -u github.com/o1egl/paseto`
6. 