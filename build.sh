go get github.com/gorilla/mux
go get github.com/jmoiron/sqlx
go get github.com/go-sql-driver/mysql
go get github.com/urfave/negroni
go build -o bin/application *.go
rm application.go