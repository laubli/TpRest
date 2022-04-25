module TpRest/src/internal/web

go 1.15

require github.com/gorilla/mux v1.8.0

require persistence v1.0.0
replace persistence => ../persistence

require entities v1.0.0
replace entities => ../entities
