package backend_test

import (
	"backendserver"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var backend backendserver.BackendRoute
var tableGetQuer = "SELECT transactionId , endorsements , args FROM EvonikEvent"

func TestMain(m *testing.M) {
	backend = backendserver.BackendRoute{}
	backend.InitRoute()
	backend.InitDB()
	ensureTableExist()
	code := m.Run()
	os.Exit(code)
}

func ensureTableExist() {
	if _, err := backend.DB.Exec(tableGetQuer); err != nil {
		log.Fatal(err.Error())
	}
}

func TestHandlerNonExitingTransacrtion(t *testing.T) {
	req, _ := http.NewRequest("GET", "/transaction/20", nil)
	response := executeRequest(req)
	fmt.Println("response", response)
	checkResponseCode(t, http.StatusInternalServerError, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "sql: no rows in result set" {
		t.Errorf("Expected the 'error' key of the response to be set to 'sql:no rows in result set'. Got '%s'", m["error"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	backend.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
