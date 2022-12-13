# Golang-RestAPI-Server

## Add some extenstion of go

- Go v0.36.0

## Add third-party

```
// before start need to initizate the module
 go mod init <name>
// this one is for routes
go get -u github.com/gorilla/mux
// this one is for database
go get -u github.com/mattn/go-sqlite3
```

## Testing

any file with \_test will run when you type

```
go run test
```

We have simple Go server Example with connect to Postgress DB

to start server you have to go server folder

```
cd Server
go run .
```

in Server folder you will finde three different folder

- backend in which we define the routes and connect with DB to perform our action
- db only have feature to connect with DB
- backend-test we write one test case to just get to know testing
