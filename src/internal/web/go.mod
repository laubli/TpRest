module github.com/laubli/TpRest/src/internal/web

go 1.15
require github.com/gorilla/mux v1.8.0

require entities v1.2.3
replace entities => ../entities

require persistence v1.0.0
replace persistence => ../persistence