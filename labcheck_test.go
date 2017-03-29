package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func TestLabIndex(t *testing.T) {
	//do
}

func TestCreateLab(t *testing.T) {
	var (
		request  *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := MakeTestServer()

	labName := "Lab10"

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/lab/"+labName, nil)
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

func TestNoTextAction(t *testing.T) {

	var s Slack
	s.Text = ""

	labs := TryAction(s)

	if labs == nil {
		t.Errorf("Error Returning labs")
	}
}

func TestAddLabAction(t *testing.T) {

	var labname = "Lab10"
	var s Slack
	s.Text = labname

	count1 := len(labs)
	labs := TryAction(s)
	count2 := len(labs)

	for _, lb := range labs {
		log.Print("lab &v", lb.Name)
	}
	log.Print("counts:", count1, count2)
	if count1 >= count2 {
		t.Errorf("Error creating Lab")
	}
}

func NewLab(name string) *Lab {
	if name != "" {
		return nil
	}
	l := Lab{"lab10", "out", "kmattson", time.Now()}
	return &l
}

func MakeTestServer() *negroni.Negroni {
	return NewServer()
}
