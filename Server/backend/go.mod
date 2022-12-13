module backendserver

go 1.18

replace dbConnect => ../db

require (
	dbConnect v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require github.com/lib/pq v1.10.7 // indirect
