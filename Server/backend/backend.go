package backendserver

import (
	"database/sql"
	"dbConnect"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BackendRoute struct {
	Router *mux.Router
	Port   string
	DB     *sql.DB
}

type transaction struct {
	TransactionId string `json: "transactionId"`
	Endorsements  string `json: "endorsements"`
	Args          string `json: "args"`
	ChaincodeName string `json: "chaincodeName"`
	FunctionName  string `json: "functionName"`
	RecordId      string `json: "recordId"`
	Ts            string `json: "ts"`
}

func (s *BackendRoute) InitRoute() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/transaction", s.getAllTransaction).Methods("GET")
	s.Router.HandleFunc("/transaction/{transactionId}", s.getTransactionById).Methods("GET")
	s.Router.HandleFunc("/newTransaction", s.createTransactionHandler).Methods("POST")

	s.Router.Handle("/", s.Router)

}

/*
Following function is for Router handler
*/
func (s *BackendRoute) createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Some error happend")
		responsewithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("reqBody", string(reqBody))

	var t transaction
	json.Unmarshal(reqBody, &t)
	fmt.Println(t.FunctionName)
	// err = t.createTransaction(s.DB)
	// if err != nil {
	// 	fmt.Println("New transaction creation error")
	// 	responsewithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// responseWithJson(w, http.StatusOK, t)

}
func (s *BackendRoute) getAllTransaction(w http.ResponseWriter, r *http.Request) {
	transaction, err := getAllTransaction(s.DB)
	if err != nil {
		fmt.Println("Some error happend")
		responsewithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, transaction)
}
func (s *BackendRoute) getTransactionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionId := vars["transactionId"]
	fmt.Println("transactionId", transactionId)
	transaction, err := getTransactionById(s.DB, transactionId)
	if err != nil {
		fmt.Println("Some error happend")
		responsewithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, transaction)
}

/*
Following function are DB call like select ,insert and delete
*/
func (t *transaction) createTransaction(DB *sql.DB) error {

	qlStatement := `
	INSERT INTO EvonikEvent (transactionId, endorsements, args, chaincodeName,functionName,recordId,ts)
	VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := DB.Exec(qlStatement, t.TransactionId, t.Endorsements, t.Args, t.ChaincodeName, t.FunctionName, t.RecordId, t.Ts)
	if err != nil {
		return err
	}
	return nil

}
func getAllTransaction(DB *sql.DB) ([]transaction, error) {

	rows, err := DB.Query("SELECT transactionId , endorsements , args FROM EvonikEvent")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []transaction

	for rows.Next() {
		// fmt.Println(rows)
		var t transaction
		err = rows.Scan(&t.TransactionId, &t.Endorsements, &t.Args)
		if err != nil {
			return nil, err
		}
		fmt.Println(&t.TransactionId, &t.Endorsements, &t.Args)

		transactions = append(transactions, t)
	}
	return transactions, nil

}

func getTransactionById(DB *sql.DB, transactionId string) (transaction, error) {
	var t transaction

	row := DB.QueryRow(`SELECT transactionId , endorsements , args,ChaincodeName from EvonikEvent where transactionId = '` + transactionId + `'`)
	if err := row.Scan(&t.TransactionId, &t.Endorsements, &t.Args, &t.ChaincodeName); err != nil { // scan will release the connection
		fmt.Println("row", err.Error())

		return t, err
	}
	return t, nil

}

func responsewithError(w http.ResponseWriter, code int, message string) {
	responseWithJson(w, code, map[string]string{"error": message})
}
func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
func (s *BackendRoute) InitDB() {
	a := dbConnect.App{}
	a.Init()
	s.DB = a.DB
}
func (s *BackendRoute) StartServer() {
	s.InitDB()
	fmt.Println("BackendRoute Started !")
	log.Fatal(http.ListenAndServe(":9003", s.Router))
}
