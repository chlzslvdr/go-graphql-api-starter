# Go-Graphql API Starter

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

## Install Golang

Follow this link to install golang https://golang.org/doc/install

## Setting up Golang \$GOPATH directory

Follow this link to setup \$GOPATH env https://golang.org/doc/gopath_code.html

## Clone your repository in \$GOPATH/go directory

1. Navigate to gopath directory by executing command

```bash
$ cd $GOPATH/go/src
```

2. Clone the repository by executing command

```bash
$ git clone [your repository]
```

## Tools and Libraries

- viper

```bash
$ go get github.com/spf13/viper
```

- gorm

```bash
$ go get -u github.com/jinzhu/gorm
```

- goose

```bash
$ go get -u github.com/pressly/goose
```

- gqlgen

```bash
$ go get github.com/99designs/gqlgen
```

- uuid

```bash
$ go get github.com/satori/go.uuid
```

## Set your environment variable for database

```bash
export dbhost=127.0.0.1
export username=[your username]
export password=[your password]
```

## Foldering structure

- config - configuration about the environment variables, database, files, etc. should be place in the folder
- graph - graphql server for golang
- data - database connection e.g. postgresql, mongodb, redis
- migrations - for migrating changes to the database
- models - domain model and table structure are the content of this folder

## Create Database

Create a database in PostgreSQL with the name:

```
sample_db
```

## Database migrations

Go to `go-graphql-api-starter/migarations` directory and execute the following command in the terminal:

1. Create sql migration file

```bash
goose create [file name or migration description] sql
```

2. Check the status of the migrations

```bash
goose postgres "user=[your user] password=[your password] dbname=psql_lookups_db host=127.0.0.1 sslmode=disable" status
```

3. Update the migrations

```bash
goose postgres "user=[your user] password=[your password] dbname=psql_lookups_db host=127.0.0.1 sslmode=disable" up
```

4. Downgrade the migrations

```bash
goose postgres "user=[your user] password=[your password] dbname=psql_lookups_db host=127.0.0.1 sslmode=disable" down
```

## Running the API

### Run directly

To run the API directly without creating executable, run the command in terminal:

```bash
$ go run server.go
```

### Create executable file

1. To create an executable file, run the command in terminal:

```bash
$ go build
```

2. Execute the generate file by specifying the name e.g. `./sample-api`

### GraphQL playground

```
http://localhost:4000/
```
