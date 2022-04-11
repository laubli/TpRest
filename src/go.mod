module github.com/laubli/TpRest/src

go 1.15

require github.com/gorilla/mux v1.8.0

require (
	github.com/boltdb/bolt v1.3.1
	github.com/go-sql-driver/mysql v1.6.0
)

require internal/entities v1.0.0
replace internal/entities => ./internal/entities
