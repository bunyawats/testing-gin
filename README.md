go get github.com/gin-gonic/gin\
go get github.com/stretchr/testify\
go get github.com/rs/xid\
go get github.com/prometheus/client_golang


go get -u

curl localhost:8080

go test

GIN_MODE=release go test -v

GIN_MODE=release go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

docker-compose build
docker-compose up -d
