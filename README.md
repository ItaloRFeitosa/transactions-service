# Transactions Service

## Requirements
* Go 1.21
* Docker and Docker Compose
* Makefile

## How To Run
Make sure that ports 8080 and 5432 are free. Otherwise, change the mapped ports in `/deployments/docker-compose.yml`
```
# run containers dattached
make prodlike

# to confirm that they are running
docker ps

# to see api logs

make prodlike_logs

# to remove containers

make prodlike_destroy
```
After running with `make prodlike`, check if api is running [here](http://localhost:8080/health)
## How To Use

After running the app, you can access the [swagger page](http://localhost:8080/swagger) to interact with the api.

Also, you can use the `/docs/api.http` to interact with the api, but install the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) vscode extension beforehand.

Inside `/docs` there are content about the capabilities of application, with diagrams and decision records. (Open in github itself, so the diagrams can be rendered)

## How To Test
```
make unit_test
```
Disclaimer: the amount of testing is not satisfactory yet, ideally the app package should be tested with the help of mocks or fakes.
Furthermore, e2e testing at the api layer should also be done to test contract and status codes.
## Local Development
```
# run api container with comstrek/air to live reload and delve debugger

make local
```
## Stack
* Golang 1.21
* Postgres 15
* Gin Gonic
* Sqlx
* Testify
* Go Swagger
* Go Migrate
* Comstrek Air