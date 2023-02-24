set dotenv-load := true

run:
    go run cmd/main.go app start

build:
    go build -o main cmd/main.go

test:
    go test .../. -coverprofile=coverage.out
    
show-cov:
    go tool cover -html=coverage.out

initdb:
    go run cmd/main.go db init

create-sql:
    go run cmd/main.go db create_sql migration

migrate:
    go run cmd/main.go db migrate

rollback:
    go run cmd/main.go db rollback

codegen: 
    oapi-codegen --config .../oapi_config.yaml .../openapi.yml

gen-swagger:
    swag init --parseDependency --parseInternal --dir ./docs/swaggo

docker-build:
    docker build --no-cache -t gin-basic-api:1.0 -f Dockerfile .

docker-run:
    docker run -p:8080:8080 --env-file ./etc/local.env --name gin-basic-api gin-basic-api:1.0

release:
    git sv next-version

project-structure:
    tree -d -L 3 .