package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestCreateLab(t *testing.T) {
	var (
		request  *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := MakeTestServer()

	labName := "Lab10"

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/labs/"+labName, nil)
	server.ServeHTTP(recorder, request)

	var lab Lab

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	err = json.Unmarshal(payload, &lab)
	if err != nil {
		t.Errorf("Error unmarshaling  response to Lab: %v", err)
	}

	if lab.Name != "Lab10" {
		t.Errorf("Expected Lab10, got %v", lab.Name)
	}

}

func NewLab(id int, name string) *Lab {
	if id < 0 {
		return nil
	}
	l := Lab{id, name, "foo", "bar", time.Now()}
	return &l
}

func MakeTestServer() *negroni.Negroni {
	return NewServer()
}
