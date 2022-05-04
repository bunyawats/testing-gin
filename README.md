go get github.com/gin-gonic/gin\
go get github.com/stretchr/testify\
go get github.com/rs/xid

go get -u

curl localhost:8080

go test

GIN_MODE=release go test -v