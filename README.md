# Test Biofarma

Backend Service of Test Biofarma

## Prerequisites

- Golang 1.20
- PostgreSQL

## Built With

- [Go](https://go.dev/)
- [Go-Swagger](https://github.com/go-swagger/go-swagger)
- [Gorm](https://gorm.io/)
- [Golang-Migrate](https://github.com/golang-migrate/migrate)
- [Urfave cli](https://github.com/urfave/cli)

## Getting Started

1. Install Go and Go-Swagger on you device
2. Clone repository
3. Setup env file

    a. Run this command
    ```bash
    cp config.env.example config.env
    ```
    b. Change the value of every env variable with your configuration

4. Migrate all migration file
    ```bash
    make migrate-up
    ```

5. Run the project

    a. Regenerate generated code from swagger specs with this command
    ```bash
    make all
    ```
    b. Running the app
    ```bash
    make run-binary
    ```
6. For API Documentation you can open [localhost:8080/docs](localhost:8080/docs)

## Database Migration
### Create new file migration with run this command
```bash
make migrate-create-file your_file_name
```
### Migrate up with this command
```bash
make migrate-up
```

## Adding API

1. Edit swagger file in folder *api*. If you need help with swagger docs use the [Editor](https://swagger.io/docs/open-source-tools/swagger-editor/) and read the [Documentation](https://swagger.io/docs/specification/about/).
2. To Generate Routes and Validation from Swagger API run `make build` or `make all`.
