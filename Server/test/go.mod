module backend_test

go 1.18

require (
	backendserver v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require (
	dbConnect v0.0.0-00010101000000-000000000000 // indirect
	github.com/lib/pq v1.10.7 // indirect
)

replace backendserver => ../backend

replace dbConnect => ../db
